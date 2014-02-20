package main

// loading in the Martini package
import (
  "net/http"  // this will allow us to use http.StatusOK and http.StatusNotFound instead of 200 and 404
  "strings"   // I’m adding this so I can ensure that we are comparing lower case strings
  "encoding/json" 
  "github.com/codegangsta/martini"
)



type Attribute struct {
  Name string `json:"name"`
  DataType string `json:"type"`
  Description string `json:"description"`
  Required bool `json:"required"`
}

type ResourceAttributes struct {
  ResourceName string `json: "resourceName"`
  Attributes []Attribute `json: "attributes"`
}

type ErrorMsg struct {
  Msg string `json:"msg"`
}

type jsonConvertible interface {}

func JsonString( obj jsonConvertible ) (s string) {
  jsonObj, err := json.Marshal( obj )
  
  if err != nil {
    s = ""
  } else {
    s = string( jsonObj )
  }

  return
}

func main() {
  // := is a short variable declaration it types and instantiates the var 'on the fly'
  m := martini.Classic()

  m.Get("/attributes/:resource", func( params martini.Params, writer http.ResponseWriter ) (int, string) {
    resource :=  strings.ToLower( params["resource"] )
    writer.Header().Set("Content-Type", "application/json")

    if resource  == "tv" {
      resourceAttrs := ResourceAttributes{"tv", make([]Attribute, 1)}
      resourceAttrs.Attributes[0] = Attribute{"Location","string", "What facility is the TV located in.", true}
     
      return http.StatusOK, JsonString( resourceAttrs )
    } else {
      return http.StatusNotFound, JsonString( ErrorMsg{"Resource not found: " + resource} )
    }

  })

  m.Run()
}
