package main

import "log"

func main() {
	conf := InitConfig()

	if conf.Embedded == true {
		if err := GenerateEmbededServer(conf); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := GenerateServer(conf); err != nil {
			log.Fatal(err)
		}
	}
}
