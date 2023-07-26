package arguments

import (
	"errors"
	"flag"
	"fmt"
)

const (
	modeFlag   = "m"
	Duplicates = "d"
	Anagrammas = "a"
	pathFlag   = "p"

	pathUsage = "Specify path to log file. Example: -p /path/to/log"
	modeUsage = "Select mode to run app. Available modes: \"d\", \"a\". d - duplciates, a - anagrammas. Example: -m d"
)

type Arguments struct {
	Path string
	Mode string
}

func New() *Arguments {
	var (
		path *string = new(string)
		mode *string = new(string)
	)
	flag.Usage = func() {
		fmt.Println("Usage of log_parser:")
		flag.PrintDefaults()
	}
	flag.StringVar(path, pathFlag, "", pathUsage)
	flag.StringVar(mode, modeFlag, "", modeUsage)
	flag.Parse()

	return &Arguments{
		Path: *path,
		Mode: *mode,
	}
}

func (args *Arguments) Validate() error {
	if args.Path == "" {
		return errors.New("path to file is empty")
	}
	if args.Mode != Anagrammas && args.Mode != Duplicates {
		return errors.New("invalid mode value")
	}

	return nil
}
