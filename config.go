package main

import (
	"flag"
)

type Config struct {
	SpecFile string
	OutDir   string
	RuleFile string
	Host     string
	Embedded bool
}

func InitConfig() Config {
	specFile := flag.String("spec", "spec.yml", "spec path")
	outDir := flag.String("out", "./source", "out directory")
	ruleFile := flag.String("rule", "./rules.json", "rule file")
	host := flag.String("host", "localhost:9000", "server host address")
	embedded := flag.Bool("embedded", false, "embedded server")

	flag.Parse()

	return Config{
		SpecFile: *specFile,
		OutDir:   *outDir,
		RuleFile: *ruleFile,
		Host:     *host,
		Embedded: *embedded,
	}
}
