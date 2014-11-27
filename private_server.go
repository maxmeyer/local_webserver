package main

import "os/signal"
import "os"
import "fmt"

import "github.com/dg-ratiodata/local_webserver/generator"
import "github.com/dg-ratiodata/local_webserver/webserver"
import "github.com/dg-ratiodata/local_webserver/application"
import "github.com/dg-ratiodata/local_webserver/cli"


func main() {
  go func(){
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, os.Interrupt)
    <-sigchan
    fmt.Print("Captured CTRL-C. Exit..\n")

    os.Exit(0)
  }()

  port, network_interface, silent, directory := cli.Parse(generator.RandomStr(1024, 65535), "127.0.0.1", application.Working_directory())

  url := fmt.Sprintf("http://%s:%d/index.html", network_interface, port)
  listen := fmt.Sprintf("%s:%d", network_interface, port)

  webserver.OpenBrowser(url, silent)
  webserver.StartServer(directory, listen, silent)
}
