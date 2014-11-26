package application

import "os"
import "log"
import "path/filepath"

func Working_directory() string {
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

  if err != nil {
    log.Fatal(err)
  }

  return dir
}
