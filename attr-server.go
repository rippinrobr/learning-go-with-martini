package main

import (
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
)

func main() {
  m := martini.Classic()
  m.Use( Mongo() )

  m.Get("/attributes/:resource",  getAttributes )
  m.Post("/attributes/:resource", binding.Json( attribute{} ), addAttribute  )

//  RegisterService
  m.Run()
}
