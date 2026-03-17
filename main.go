package main

import (
	"os"
	"fmt"
	"net/http"
	"time"
)

var num_called int = 0 

func greet(w http.ResponseWriter, r *http.Request) {
  num_called+=1 

  fmt.Printf( "{ctx: kafka-client, Adding consumer}");
  fmt.Printf( "{ctx: kafka-client, Consumer added >> {}}");
  fmt.Printf( "{ctx: app,\n####                                 ######                                                  \n #     # #       ####   ####  ######    #     # #####   ####  ##### #    # ###### #####   ####  }");
  
  // fmt.Printf( "fg Hello World called %d ! %s\n", num_called,  time.Now() )
  fmt.Fprintf(w, "fg Hello World called %d ! %s", num_called,  time.Now())
  os.Exit(1)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
