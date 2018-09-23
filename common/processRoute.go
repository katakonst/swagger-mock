package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
)

func ProcessRule(w http.ResponseWriter, r *http.Request, rules []Rule, opId string) error {
	body := readBody(r)
	for _, rule := range rules {
		if rule.OpId == opId {
			found := false
			for _, arg := range rule.Args {
				switch argType := arg.ArgType; argType {
				case "query":
					if exists := CheckQueryParam(r, arg); exists == true {
						found = true
						break
					}

				case "form":
					if exists := CheckFormParam(r, arg); exists == true {
						found = true
						break
					}

				case "JSON":
					if exists := CheckJSONParam(r, arg, body); exists == true {
						found = true
						break
					}

				case "path":
					if exists := CheckPathParam(r, arg); exists == true {
						found = true
						break
					}

				default:
					return fmt.Errorf("Invalid arg type")
				}
			}
			if found == true {

				data, err := json.Marshal(rule.Response)
				if err != nil {
					return fmt.Errorf("Processing body: %v", err)
				}

				if d, err := time.ParseDuration(rule.Timeout); err == nil {
					time.Sleep(d)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(rule.StatusCode)
				w.Write(data)
				return nil
			}
		}
	}
	return fmt.Errorf("no rule found")
}

func CheckQueryParam(r *http.Request, arg Argument) bool {
	param := r.URL.Query().Get(arg.ArgName)
	return param == arg.Body
}

func CheckPathParam(r *http.Request, arg Argument) bool {
	vars := mux.Vars(r)
	return vars[arg.ArgName] == arg.Body
}

func CheckFormParam(r *http.Request, arg Argument) bool {
	r.ParseForm()
	param := r.FormValue(arg.ArgName)
	return param == arg.Body
}

func CheckJSONParam(r *http.Request, arg Argument, body []byte) bool {

	logger := NewLogger("info")
	var bodyReq interface{}
	if err := json.Unmarshal(body, &bodyReq); err != nil {
		logger.Errorf("Error while unmarshaling body %v", err)
		return false
	}
	return reflect.DeepEqual(bodyReq, arg.Body)
}

func readBody(r *http.Request) []byte {
	logger := NewLogger("info")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error while reading body %v", err)
		return nil
	}
	return body
}
