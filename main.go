package main

import (
	"fmt"
	"net/http"
	"time"
)

var num_called int = 0 

func greet(w http.ResponseWriter, r *http.Request) {
  num_called+=1 
	fmt.Fprintf(w, "fg Hello World called %d ! %s",, num_called,  time.Now())
  os.exit(1)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
