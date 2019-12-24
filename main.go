package main

import (
  "fmt"
  "os"
  "os/user"
  "interpreter/repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }

  fmt.Printf("Hello %s, Feel free to type in commands\n", user.Username)
  repl.Start(os.Stdin, os.Stdout)
}
