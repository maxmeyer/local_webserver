package webserver

import "fmt"
import "log"
import "net/http"
import "os"

import "github.com/skratchdot/open-golang/open"
import "github.com/gorilla/handlers"

func Open_browser(url string) {
  fmt.Printf("Open browser with %s\n", url)

  err := open.Start(url)

  if err != nil {
    log.Fatal(err)
  }
}

func Start(directory string, bind string) {
  fmt.Printf("Server listens on %s\n", bind)

  http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(directory))))
  panic(http.ListenAndServe(bind, nil))
}
