package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	templateDir := flag.String("template_dir", "../templates", "out")
	commonDir := flag.String("common_dir", "../common", "out")
	outFile := flag.String("out_dir", "../templates.go", "out")

	files, err := ioutil.ReadDir(*templateDir)
	if err != nil {
		log.Fatal(err)
	}
	fileMap := make(map[string]string)

	for _, f := range files {
		data, err := ioutil.ReadFile(*templateDir + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		fileMap[f.Name()] = string(data)
	}

	commonFiles, err := ioutil.ReadDir(*commonDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range commonFiles {
		data, err := ioutil.ReadFile(*commonDir + "/" + f.Name())
		if err != nil {
			log.Fatal(err)
		}

		fileMap[f.Name()] = strings.Replace(string(data), "package common", "package main", 1)
	}

	mp := GenerateMap(fileMap)
	f, err := os.Create(*outFile)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	if err = ioutil.WriteFile(*outFile, []byte(mp), 0644); err != nil {
		log.Fatal(err)
	}
}

func GenerateMap(m map[string]string) string {
	mapDeclaration := "package main; var loadedTempl = map[string]string{"
	for k, v := range m {
		mapDeclaration = mapDeclaration + " " + "\"" + k + "\"" + " : " + strconv.Quote(v) + ","
	}
	mapDeclaration = strings.TrimSuffix(mapDeclaration, ",")
	mapDeclaration = mapDeclaration + "}"
	return mapDeclaration
}

func copyFile(fileName string, dest string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("fillTemplate:createfile %v", err)
	}
	f.Close()

	if err = ioutil.WriteFile(dest, data, 0644); err != nil {
		return fmt.Errorf("Error creating file %s", dest)
	}
	return nil
}
