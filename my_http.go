package main

import "fmt"
import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Response with This")
}

func main()  {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}