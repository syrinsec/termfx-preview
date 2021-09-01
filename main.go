package main

import (
	"flag"
	"fmt"
  "io"
  "io/ioutil"
	"os"
  "strconv"
  "time"

  "github.com/syrinsec/termfx"
)

// Variables used for command line parameters
var (
	Path string
)
func init() {
	flag.StringVar(&Path, "f", "", "Termfx File Path")
	flag.Parse()
}

func main() {
  registry := termfx.New()

  registry.RegisterFunction("sleep", func(session io.Writer, args string) (int, error) {

    sleep, err := strconv.Atoi(args)
    if err != nil {
      return 0, err
    }

    time.Sleep(time.Millisecond * time.Duration(sleep))
    return 0, nil
  })
  registry.RegisterFunction("clear", func(session io.Writer, args string) (int, error) {
    fmt.Println("c[?25l[0;0H")
    return 0, nil
  })

  data, err := ioutil.ReadFile(Path)
  if err != nil {
    fmt.Println("Error finding termfx file", err, "\r\nSyntax: termfx-preview -f <path>")
    return
  }
  registry.Execute(string(data), os.Stdout)
}
