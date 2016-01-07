package confu

import (
	"flag"
	"fmt"
	"testing"
)

func TestSmokeTokenize(t *testing.T) {
	__comment := `some 'comment' ::s`
	__path := "/some/path"
	__port := 8081

	input := fmt.Sprintf(`   --comment="%v" --port=%v
		-path   '%v'
		--save  `, __comment, __port, __path)

	//1 - tokenize input config string
	args := Tokenize(input)

	var conf struct {
		save    bool
		path    string
		port    int
		tag     bool
		comment string
	}

	//2 - load parsed config into out conf struct
	set := &flag.FlagSet{}

	set.BoolVar(&conf.save, "save", false, "")
	set.StringVar(&conf.path, "path", "./", "")
	set.IntVar(&conf.port, "port", 8080, "")
	set.BoolVar(&conf.tag, "tag", false, "")
	set.StringVar(&conf.comment, "comment", "", "")

	set.Parse(args)

	if conf.tag || !conf.save || conf.port != __port || conf.path != __path {
		t.Error(1)
		t.Fail()
	}

	if conf.comment == __comment {
		t.Error(2)
		t.Fail()
	}

	conf.comment = TrimQuote(conf.comment)

	if conf.comment != __comment {
		t.Error(3)
		t.Fail()
	}

	t.Logf("%+v", conf)
}

func BenchmarkDumb(b *testing.B) {
	input := `--tag --comment="done" --port=8081  --path '/some/path'`

	for n := 0; n < b.N; n++ {
		Tokenize(input)
	}
}

//func TestTemplate(t *testing.T) {
//}

//func BenchmarkTemplate(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//	}
//}
