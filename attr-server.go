package main

import (
  "fmt"

  "github.com/rippinrobr/learning-go-with-martini/config"
  "github.com/rippinrobr/learning-go-with-martini/utils"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/codegangsta/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  m.Use( Mongo() )
  // add middle ware here for rendering the html
  m.Use( render.Renderer() )

  m.Get("/attributes/:resource",  getAttributes )
  m.Post("/attributes/:resource", binding.Json( attribute{} ), addAttribute  )

  service := utils.ServiceDescription{"attributes", "http://localhost:3000", "Service that manages the attributes available for each resource type"}
  fmt.Println( config.RegisterService( service ) )
  
  m.Run()
}
