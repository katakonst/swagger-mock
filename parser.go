package main

import (
	"fmt"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
)

type Parser struct {
	fileName string
}

func (p *Parser) parseSpec() (*analysis.Spec, error) {
	specDoc, err := loads.Spec(p.fileName)
	if err != nil {
		return nil, fmt.Errorf("parseSpec: %v", err)
	}
	return analysis.New(specDoc.Spec()), nil
}

func (p *Parser) parseServer(config Config) (Server, error) {
	spec, err := p.parseSpec()
	if err != nil {
		return Server{}, err
	}

	var routes []Route
	for _, op := range spec.OperationIDs() {
		method, path, _, _ := spec.OperationForName(op)
		routes = append(routes, Route{Method: method, OpId: op, Path: path})
	}
	return Server{Host: config.Host, Routes: routes, RuleFile: config.RuleFile}, nil
}
