package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var envPath = "/env/"
var listPath = "/list/"

func lookupEnvKey(key string) string {
	return os.Getenv(key)
}

func listEnv() string {
	return strings.Join(os.Environ(), ";")
}

func envHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len(envPath):]
	envValue := lookupEnvKey(key)
	fmt.Fprintf(w, envValue)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, listEnv())
}

func main() {
	http.HandleFunc(envPath, envHandler)
	http.HandleFunc(listPath, listHandler)
	http.ListenAndServe(":8080", nil)
}
