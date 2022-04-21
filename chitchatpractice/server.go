package main

import (
	"net/http"
)

func main() {

	//create new instance of the multiplexer
	mux := http.NewServeMux()

	//use the multplexer to serve a static file
	files := http.FileServer(http.Dir("/public"))
	//use the StripPrefix funtion to remove the given prefix from the request URLS's path
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	//	files := []string{"templates/layout.html",
	//		"templates/navbar.html",
	//		"templates/index.html"}
	//	templates := template.Must(template.ParseFiles(files...))
	//	threads, err := data.Threads()
	//	if err == nil {
	//		templates.ExecuteTemplate(w, "layout", threads)
	//}

}
