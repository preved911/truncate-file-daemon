package main

import (
	"os"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var log *logrus.Logger
var filenames []string

func main() {
	logLevel := pflag.StringP("log-level", "v", "info", "log level")
	interval := pflag.StringP("interval", "i", "@every 1h", "given files truncate interval in cron notation")
	pflag.Parse()
	filenames = pflag.Args()

	log = logrus.New()
	lvl, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		log.Fatal("failed to recognize given log level")
	}
	log.SetLevel(lvl)

	c := cron.New(
		cron.WithLogger(cron.VerbosePrintfLogger(log)))
	c.AddFunc(*interval, truncate)
	c.Run()

	for {
	}
}

func truncate() {
	for _, filename := range filenames {
		err := os.Truncate(filename, 0)
		if err != nil {
			log.Fatalf("failed to truncate filename %s: %s", filename, err)
			continue
		}

		log.Debugf("file %s was truncated", filename)
	}
}
