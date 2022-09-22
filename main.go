package main

import ( "fmt"
"net/http"
"os"
"strings"
)
func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	addr := ":" + port

	fmt.Printf("Starting web server on port %s\n", "addr")

	// err := http.ListenAndServe(addr, http.HandlerFunc(webServer))
	err := http.ListenAndServe(addr, &Server{})
	if err !=nil {
		panic(err)
	}
	}

// 	func webServer(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(200)
// 		w.Write([]byte("Hello World :/"))
// }

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" || r.Method == "HEAD" {
		noteID := strings.TrimPrefix(r.URL.Path, "/")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hi Cockhead, you requested the note with the ID %s", noteID)))
		return
}

if r.Method == "POST" && r.URL.Path == "/notes" {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi Cockhead, you posted to the /notes."))
	return
}

w.WriteHeader(http.StatusNotFound)
w.Write([]byte("Hi Cockhead, you requested a page that does not exist.")) }