package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	PORT       = "80"
	ENV_PREFIX = "FEM"
)

// Hello handler greets the user.
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Display Server\n"))
}

// DisplayEnvVars handler displays environment variables with "FEM" prefix
func DisplayEnvVars(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "--- Environment Variables with '"+ENV_PREFIX+"' prefix ---\n\n\n")
	for key, value := range getEnvVarsWithPrefix(ENV_PREFIX) {
		fmt.Fprintf(w, "%s = %s\n", key, value)
	}
}

// getEnvVarsWithPrefix is implementation of DisplayEnvVars handler
func getEnvVarsWithPrefix(prefix string) map[string]string {
	envVars := make(map[string]string)

	for _, keypair := range os.Environ() {
		key, value, found := strings.Cut(keypair, "=")
		if found && strings.HasPrefix(key, prefix) {
			envVars[key] = value
		}
	}

	return envVars
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/env", DisplayEnvVars)

	// Print a log a message to say that the server is starting.
	log.Println("Starting server on port", PORT)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", PORT), mux))
}
