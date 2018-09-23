package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/katakonst/swagger-mock/common"
)

type Server struct {
	Host     string
	RuleFile string
	Routes   []Route
}

type Route struct {
	OpId     string
	Method   string
	Path     string
	RuleFile string
}

type ServerRoute struct {
	route     Route
	routeFunc interface{}
}

var commonFiles = []string{"MockServer.go", "processRoute.go", "log.go"}

func GenerateServer(config Config) error {

	p := Parser{fileName: config.SpecFile}
	server, err := p.parseServer(config)
	if err != nil {
		return err
	}
	server.RuleFile = filepath.Base(config.RuleFile)

	t := Template{FileName: "main.tmpl", Entity: server, FinalPath: config.OutDir + "/main.go"}
	if err = t.fillTemplate(); err != nil {
		return err
	}

	for _, route := range server.Routes {
		route.RuleFile = server.RuleFile
		tmpl := Template{FileName: "route.tmpl", Entity: route, FinalPath: config.OutDir + "/" + route.OpId + ".go"}
		if err = tmpl.fillTemplate(); err != nil {
			return err
		}
	}
	for _, file := range commonFiles {
		if err = writeFile(loadedTempl[file], config.OutDir+"/"+file); err != nil {
			return err
		}
	}
	return copyFile(config.RuleFile, config.OutDir+"/"+filepath.Base(config.RuleFile))
}

func writeFile(data string, dest string) error {

	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("writeFile:createfile %v", err)
	}
	f.Close()

	if err = ioutil.WriteFile(dest, []byte(data), 0644); err != nil {
		return fmt.Errorf("writeFile %s: %v", dest, err)
	}
	return nil
}

func copyFile(fileName string, dest string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("copyFile:createfile %v", err)
	}
	f.Close()

	if err = ioutil.WriteFile(dest, data, 0644); err != nil {
		return fmt.Errorf("writeFile %s: %v", dest, err)
	}
	return nil
}

func GenerateEmbededServer(config Config) error {

	p := Parser{fileName: config.SpecFile}
	server, err := p.parseServer(config)
	if err != nil {
		return err
	}
	var serverRoutes []ServerRoute

	for _, route := range server.Routes {
		serverRoute := ServerRoute{}
		serverRoute.route = route
		serverRoute.routeFunc = func(w http.ResponseWriter, r *http.Request) {
			mockServer := common.ParseRules(config.RuleFile)
			logger := common.NewLogger("info")
			if err := common.ProcessRule(w, r, mockServer.Rules, serverRoute.route.OpId); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				logger.Errorf("Error while processing opid %s, %v", route.OpId, err)
				return
			}
		}
		serverRoutes = append(serverRoutes, serverRoute)
	}
	ListenServer(serverRoutes, config)
	return nil
}

func ListenServer(serverRoutes []ServerRoute, conf Config) {
	router := mux.NewRouter()
	for _, route := range serverRoutes {
		router.HandleFunc(route.route.Path,
			route.routeFunc.(func(http.ResponseWriter, *http.Request))).
			Methods(route.route.Method)
	}
	logger := common.NewLogger("info")
	logger.Infof("Server started on %s", conf.Host)
	log.Fatal(http.ListenAndServe(conf.Host, router))
}
