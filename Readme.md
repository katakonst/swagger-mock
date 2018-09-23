# Generate mock server from swagger spec

This tool generate a mock-server from a swagger spec

# Perequisites
 - install golang
 - Run `go get .` in swagger-mock directory

# Build
```sh
$ ./build.sh
```
Will copy files to template folder and build swagger-mock executable

# Usage

#### Generate Go Sources

```sh
$ ./swagger-mock -spec=test.yml -out=sources  -rule=rules.json -host=localhost:9000
```

This will generate go sources in source folder

#### Run Embedded server 
```sh
$ ./swagger-mock -spec=spec.yml  -rule=rules.json -host=localhost:9000 -embedded=true
```



