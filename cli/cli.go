package cli

import "gopkg.in/alecthomas/kingpin.v1"

const version = "0.0.1"

func Parse(port_default string, network_interface_default string, directory_default string) (int, string, bool, string) {
  var (
    port              = kingpin.Flag("port", "The port to be used.").Default(port_default).Int()
    network_interface = kingpin.Flag("interface",  "The network interface to be used.").Default(network_interface_default).String()
    silent            = kingpin.Flag("silent", "Make server silent").Bool()
    directory         = kingpin.Flag("directory", "Directory to serve files from").Default(directory_default).String()
  )

  kingpin.Version(version)
  kingpin.Parse()

  return *port, *network_interface, *silent, *directory
}
