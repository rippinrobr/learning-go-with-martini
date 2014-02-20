package main

import (
  "github.com/codegangsta/martini"
)

func main() {
  m := martini.Classic()
  m.Use( Mongo() )

  m.Get("/attributes/:resource", getAttributes )

  m.Run()
}
