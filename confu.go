//Provides a way for doing the configuration just the way we use the command line.
package confu

import (
	"bufio"
	"strings"
)

//Tokenize prepares a string in a way than can get parsed using a `flag.FlagSet`,
//from standard Go library.
//We can have a configuration string with command line argument syntax, anywhere.
//So essentially there could be a config file  with command line argument syntax
//and we read it all as a single string. Then it can be used with `flag.FlagSet`
//to be loaded into the target `struct`.
//
//sample usage:
//	input := `--tag --comment="done" --port=8081  --path '/some/path'`
//	args := Tokenize(input)
//
//	var conf struct {
//		save    bool
//		path    string
//		port    int
//		tag     bool
//		comment string
//	}
//
//	set := &flag.FlagSet{}
//
//	set.BoolVar(&conf.save, "save", false, "")
//	set.StringVar(&conf.path, "path", "./", "")
//	set.IntVar(&conf.port, "port", 8080, "")
//	set.BoolVar(&conf.tag, "tag", false, "")
//	set.StringVar(&conf.comment, "comment", "", "")
//
//	set.Parse(args)
//	log.Printf("%+v", conf)
//
//and the output will be:
//	{save:false path:/some/path port:8081 tag:true comment:"done"}
//
//this function splits the input string based on spaces & new line char will get replaced by space
func Tokenize(stringArgs string) []string {
	args := sanitizeLines(stringArgs)

	var result []string
	var buffer string
	weAreInQ := empty

	sl := strings.Split(args, space)

	for _, v := range sl {
		eq := entering(v)
		if weAreInQ == empty && eq != empty {
			weAreInQ = eq
			buffer = buffer + space + v
			continue
		}

		if weAreInQ != empty && exiting(v, weAreInQ) {
			buffer = buffer + space + v
			//result = append(result, buffer)
			result = appendIt(result, buffer)
			buffer = empty
			weAreInQ = empty
			continue
		}

		if weAreInQ != empty {
			buffer = buffer + space + v
			continue
		}

		//result = append(result, v)
		result = appendIt(result, v)
	}

	return result
}

func sanitizeLines(s string) string {
	strReader := strings.NewReader(s)
	reader := bufio.NewReader(strReader)

	var lines []string
	for {
		line, _, err := reader.ReadLine()
		if err != nil || line == nil {
			break
		}
		lines = append(lines, string(line))
	}

	return strings.Join(lines, space)
}

func appendIt(slist []string, s string) []string {
	vstr := TrimQuote(s)

	if len(vstr) == 0 {
		return slist
	}

	return append(slist, vstr)
}

//sample usage:
//	conf.comment = TrimQuote(conf.comment)
//like when we have and arg of shape --comment="done"
func TrimQuote(s string) string {
	vstr := strings.TrimSpace(s)

	qlist := []string{dq, sq, bq}

	for _, vq := range qlist {
		if strings.HasPrefix(vstr, vq) && strings.HasSuffix(vstr, vq) {
			vstr = strings.TrimPrefix(vstr, vq)
			vstr = strings.TrimSuffix(vstr, vq)
		}
	}

	return vstr
}

func exitingHelper(v, q string) bool {
	cond := !strings.HasPrefix(v, q) && strings.HasSuffix(v, q)
	return cond
}

func exiting(v, q string) bool {
	if q == empty {
		return false
	}
	return exitingHelper(v, q)
}

func enteringHelper(v, q string) bool {
	cond := strings.HasPrefix(v, q) && !strings.HasSuffix(v, q)
	cond = cond || (strings.HasPrefix(v, `-`) && strings.Contains(v, q) && !strings.HasSuffix(v, q))
	return cond
}

func entering(v string) string {
	if enteringHelper(v, dq) {
		return dq
	}
	if enteringHelper(v, sq) {
		return sq
	}
	if enteringHelper(v, bq) {
		return bq
	}
	return empty
}

const (
	empty = ""
	space = " "
	dq    = `"`
	sq    = "'"
	bq    = "`"
)
