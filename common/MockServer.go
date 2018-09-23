package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type MockServer struct {
	Rules []Rule
}

type Rule struct {
	OpId       string
	Timeout    string
	Method     string
	Args       []Argument
	StatusCode int
	Response   interface{}
}

type Argument struct {
	ArgType string
	ArgName string
	Body    interface{}
}

func ParseRules(filename string) *MockServer {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var mock MockServer
	err = json.Unmarshal(data, &mock)
	if err != nil {
		log.Fatal(err)
	}
	return &mock
}
