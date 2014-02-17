package main

// loading in the Martini package
import (
  "github.com/codegangsta/martini"
  "net/http"
  "encoding/json"
   "strings"
)

type ErrorMsg struct {
  Msg string `json:"msg"`
}

func main() {
  // if you are new to Go the := is a short variable declaration
  m := martini.Classic()

  m.Get("/attributes/:resource", func( params martini.Params ) (int, string) {
    resource :=  strings.ToLower( params["resource"] )

    if resource  == "tv" {
      return http.StatusOK, "/attributes/" + resource
    } else {
      return http.StatusNotFound, "JSON Object here"
    }
  })

  m.Run()
}
