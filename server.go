package main

import "os/signal"
import "os"
import "fmt"

import "github.com/dg-ratiodata/local_webserver/generator"
import "github.com/dg-ratiodata/local_webserver/webserver"
import "github.com/dg-ratiodata/local_webserver/application"
import "gopkg.in/alecthomas/kingpin.v1"

var (
  port              = kingpin.Flag("port", "The port to be used.").Default(generator.RandomStr(1024, 65535)).Int()
  network_interface = kingpin.Flag("interface",  "The network interface to be used.").Default( "127.0.0.1").String()
  silent            = kingpin.Flag("silent", "Make server silent").Bool()
)

const version = "0.0.1"

func main() {
  go func(){
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, os.Interrupt)
    <-sigchan
    fmt.Print("Captured CTRL-C. Exit..\n")

    os.Exit(0)
  }()

  kingpin.Version(version)
  kingpin.Parse()

  webserver.Open_browser(fmt.Sprintf("http://%s:%d/index.html", *network_interface, *port), *silent)
  webserver.Start(application.Working_directory(), fmt.Sprintf("%s:%d", *network_interface, *port), *silent)
}
