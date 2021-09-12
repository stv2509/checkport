package checkport

import (
	flag "github.com/spf13/pflag"
	"log"
	"os"
)

var ShowHelp bool
var ConfigPath string

func init() {
	flag.StringVarP(&ConfigPath, "config", "c", "",
		"Config file path")
	flag.BoolVarP(&ShowHelp, "help", "h", false,
		"Show help message")
	flag.Parse()
	if ShowHelp {
		flag.Usage()
		os.Exit(0)
		return
	}

	if ConfigPath == "" {
		flag.Usage()
		os.Exit(1)
		return
	}
	checkBasePath(ConfigPath)
}

func checkBasePath(basepath string) {
	if _, err := os.Stat(basepath); os.IsNotExist(err) {
		log.Fatalf("Error: %s\nFile \"%s\" not found\nPlease check you configuration file", err, basepath)
	}
}
