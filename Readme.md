# Generate mock server from swagger spec

This tool generate a mock-server from a swagger spec

# Perequisites
 - Install Golang
 - Run `go get .` in swagger-mock directory

# Build
```sh
$ ./build.sh
```
Will copy files to template folder and build swagger-mock executable

# Usage

#### Generate Go Sources
This will generate go sources in source folder

```sh
$ ./swagger-mock -spec=test.yml -out=sources  -rule=rules.json -host=localhost:9000
```

#### Run Embedded server 
This will start the embedded server
```sh
$ ./swagger-mock -spec=spec.yml  -rule=rules.json -host=localhost:9000 -embedded=true
```

# Config
Example 
```json
{
  "rules": [
    {
      "opId": "createPets",
      "method" : "POST",
      "timeout": "10s",
      "statusCode":201,
      "args": [{
        "argType": "JSON",
        "body":{"id":"1"}
      }],
      "response":
        {
          "id":"3",
          "name": "4"
        }
      
    }
  ]
}
```
This will serve the response for a request for the path from createPets operation id if the request body is `{"id":"1"}`.
In the repo exists a rule.json and a spec.yml for testing

