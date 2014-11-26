package webserver

import "fmt"
import "log"
import "net/http"
import "os"

import "github.com/skratchdot/open-golang/open"
import "github.com/gorilla/handlers"

func Open_browser(url string, silent bool) {
  if ( silent == false ) {
    fmt.Printf("Open browser with %s\n", url)
  }

  err := open.Start(url)

  if err != nil {
    log.Fatal(err)
  }
}

func Start(directory string, bind string, silent bool) {
  if ( silent == false ) {
    fmt.Printf("Server listens on %s\n", bind)
    fmt.Printf("\nRequests:\n")
  }

  var http_handler http.Handler

  if (silent == true) {
    http_handler = http.FileServer(http.Dir(directory))
  } else {
    http_handler = handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(directory)))
  }

  http.Handle("/", http_handler)
  panic(http.ListenAndServe(bind, nil))
}
