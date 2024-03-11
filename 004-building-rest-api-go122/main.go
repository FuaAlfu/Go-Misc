package main

import(
	"fmt"
	"net/http"
)

func server(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello world..")
	})

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "return all comments")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id")
		fmt.Fprintf(w, "return a single comment for comment with id: %s",id)
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "post a new comment")
	})

	if err := http.ListenAndServe("Localhost: 8080",mux); err != nil{
		fmt.Println(err.Error())
	}
}

func main() {
	fmt.Println("Testing the new features in Go")
	server()
}