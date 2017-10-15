package confu

import (
	"flag"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeInput(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func TestCommon(t *testing.T) {
	assert := assert.New(t)

	input := makeInput(`--bool -s %v --int %v`, "txt", 10)

	// 1 - tokenize input config string
	args := Tokenize(input)

	var conf struct {
		boolArg   bool
		stringArg string
		intArg    int
	}

	// 2 - load parsed config into out conf struct
	set := &flag.FlagSet{}

	set.BoolVar(&conf.boolArg, "bool", false, "")
	set.StringVar(&conf.stringArg, "s", "./", "")
	set.IntVar(&conf.intArg, "int", 1000, "")

	set.Parse(args)

	// 3 - use the conf (here; just testing)
	assert.True(conf.boolArg)
	assert.Equal("txt", conf.stringArg)
	assert.Equal(10, conf.intArg)
}

func TestParam(t *testing.T) {
	assert := assert.New(t)

	stringArg := "VAL-1"
	intArg := 10
	stringParam := "VAL-2"
	intParam := 11

	input := makeInput(`--bool -strarg %v 		--intarg %v
		-strparam=%v
		-intparam=%v`,
		stringArg,
		intArg,
		stringParam,
		intParam)

	// 1 - tokenize input config string
	args := Tokenize(input)

	var conf struct {
		boolArg     bool
		stringArg   string
		intArg      int
		stringParam string
		intParam    int
	}

	// 2 - load parsed config into out conf struct
	set := &flag.FlagSet{}

	set.BoolVar(&conf.boolArg, "bool", false, "")
	set.StringVar(&conf.stringArg, "strarg", "-", "")
	set.IntVar(&conf.intArg, "intarg", -1, "")
	set.StringVar(&conf.stringParam, "strparam", "-", "")
	set.IntVar(&conf.intParam, "intparam", -1, "")

	set.Parse(args)

	// 3 - use (here; just testing)
	assert.True(conf.boolArg)
	assert.Equal(stringArg, conf.stringArg)
	assert.Equal(intArg, conf.intArg)
	assert.Equal(stringParam, conf.stringParam)
	assert.Equal(intParam, conf.intParam)
}

func TestQuotes(t *testing.T) {
	assert := assert.New(t)

	string1 := `"VAL FOR 1"`
	string2 := `'VAL FOR 2'`
	string3 := "`VAL FOR 3`"

	input := makeInput("--str1 %v --str2 %v --str3 %v --str4=%v --str5=%v --str6=%v",
		string1,
		string2,
		string3,
		string1,
		string2,
		string3)

	// 1 - tokenize input config string
	args := Tokenize(input)

	var conf struct {
		string1 string
		string2 string
		string3 string
		string4 string
		string5 string
		string6 string
	}

	// 2 - load parsed config into out conf struct
	set := &flag.FlagSet{}

	set.StringVar(&conf.string1, "str1", "-", "")
	set.StringVar(&conf.string2, "str2", "-", "")
	set.StringVar(&conf.string3, "str3", "-", "")
	set.StringVar(&conf.string4, "str4", "-", "")
	set.StringVar(&conf.string5, "str5", "-", "")
	set.StringVar(&conf.string6, "str6", "-", "")

	set.Parse(args)

	// 3 - use (here; just testing)
	assert.Equal(TrimQuote(string1), TrimQuote(conf.string1))
	assert.Equal(TrimQuote(string2), TrimQuote(conf.string2))
	assert.Equal(TrimQuote(string3), TrimQuote(conf.string3))
	assert.Equal(TrimQuote(string1), TrimQuote(conf.string4))
	assert.Equal(TrimQuote(string2), TrimQuote(conf.string5))
	assert.Equal(TrimQuote(string3), TrimQuote(conf.string6))
}

func TestQuotesWithUnicode(t *testing.T) {
	assert := assert.New(t)

	string1 := `"تست ۱"`
	string2 := `'تست ۲'`
	string3 := "`تست ۳`"

	input := makeInput("--str1 %v --str2 %v --str3 %v --str4=%v --str5=%v --str6=%v",
		string1,
		string2,
		string3,
		string1,
		string2,
		string3)

	// 1 - tokenize input config string
	args := Tokenize(input)

	var conf struct {
		string1 string
		string2 string
		string3 string
		string4 string
		string5 string
		string6 string
	}

	// 2 - load parsed config into out conf struct
	set := &flag.FlagSet{}

	set.StringVar(&conf.string1, "str1", "-", "")
	set.StringVar(&conf.string2, "str2", "-", "")
	set.StringVar(&conf.string3, "str3", "-", "")
	set.StringVar(&conf.string4, "str4", "-", "")
	set.StringVar(&conf.string5, "str5", "-", "")
	set.StringVar(&conf.string6, "str6", "-", "")

	set.Parse(args)

	// 3 - use (here; just testing)
	assert.Equal(TrimQuote(string1), TrimQuote(conf.string1))
	assert.Equal(TrimQuote(string2), TrimQuote(conf.string2))
	assert.Equal(TrimQuote(string3), TrimQuote(conf.string3))
	assert.Equal(TrimQuote(string1), TrimQuote(conf.string4))
	assert.Equal(TrimQuote(string2), TrimQuote(conf.string5))
	assert.Equal(TrimQuote(string3), TrimQuote(conf.string6))
}

func ExampleTokenize() {
	// input could come from a file or any other source
	input := makeInput(`--bool -s %v --int %v`, "txt", 10)

	// 1 - tokenize input config string
	args := Tokenize(input)

	var conf struct {
		boolArg   bool
		stringArg string
		intArg    int
	}

	// 2 - load parsed config into out conf struct
	set := &flag.FlagSet{}

	set.BoolVar(&conf.boolArg, "bool", false, "")
	set.StringVar(&conf.stringArg, "s", "./", "")
	set.IntVar(&conf.intArg, "int", 1000, "")

	set.Parse(args)

	// 3 - use the conf
	// ...
}
