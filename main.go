package main

import (
	//"os"
	"fmt"
	"net/http"
	"time"
)

var num_called int = 0 

func greet(w http.ResponseWriter, r *http.Request) {
  num_called+=1 

 // fmt.Println( "2026-03-17T12:28:53.469Z info {ctx: kafka-client, Adding consumer}");
 // fmt.Println( "2026-03-17T22:28:53.469Z info {ctx: kafka-client, Consumer added >> {}}");
 // fmt.Println( "2026-03-17T12:28:53.469Z info {ctx: app,\n####                                 ######                                                  \n #     # #       ####   ####  ######    #     # #####   ####  ##### #    # ###### #####   ####  }");
  
  fmt.Println( "fgHelloWorld App on stage " )
  fmt.Fprintf(w, "fg Hello World called %d ! %s", num_called,  time.Now())
  //os.Exit(1)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
