package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Template struct {
	FileName  string
	Entity    interface{}
	FinalPath string
}

func (t *Template) fillTemplate() error {
	data := loadedTempl[t.FileName]
	templ, err := template.New("todos").Parse(string(data))
	if err != nil {
		return fmt.Errorf("fillTemplate:parse %v", err)
	}

	os.Mkdir(filepath.Dir(t.FinalPath), 0777)
	f, err := os.Create(t.FinalPath)
	if err != nil {
		return fmt.Errorf("fillTemplate:createfile %v", err)
	}

	defer f.Close()
	err = templ.Execute(f, t.Entity)
	if err != nil {
		return fmt.Errorf("fillTemplate:execute %v", err)
	}
	return nil
}
