# confu
Provides a way for doing the configuration just the way we use the command line.

Function Tokenize prepares a string in a way than can get parsed using a `flag.FlagSet`, from standard Go library. We can have a configuration string with command line argument syntax, anywhere. So essentially there could be a config file with command line argument syntax and we read it all as a single string. Then it can be used with `flag.FlagSet` to be loaded into the target `struct`.

```go
//input could come from a file (conf.confu maybe?)
input := `--tag --comment="done" --port=8081  --path '/some/path'`
args := Tokenize(input)

var conf struct {
  save    bool
  path    string
  port    int
  tag     bool
  comment string
}

set := &flag.FlagSet{}

set.BoolVar(&conf.save, "save", false, "")
set.StringVar(&conf.path, "path", "./", "")
set.IntVar(&conf.port, "port", 8080, "")
set.BoolVar(&conf.tag, "tag", false, "")
set.StringVar(&conf.comment, "comment", "", "")

set.Parse(args)
log.Printf("%+v", conf)
```

And the output will be (our loaded `conf` struct): `{save:false path:/some/path port:8081 tag:true comment:"done"}`.
