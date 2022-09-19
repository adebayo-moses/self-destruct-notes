package main

import "fmt"
"net/http"
"os"

func main() {
	port := os.Getenv("PORT")
	IF len(port) == 0 {
		port = "3000"
	}
	addr := ";" + port

	fmt.Printf("Starting web server on port %s\n", addr")

	err := https.ListenAndServe(addre, https.HandlerFunc(webServer))
	if err !=nil {
		panic(err)
	}
	}

	function webServer(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World :/"))
}