package main

import (
  "fmt"
  "net/http"
  "strings"
  "encoding/json"
 "github.com/codegangsta/martini"
)

// Structs and related code
type attribute struct {
  Id string `json:"id"`
  Name string `json:"name"`
  DataType string `json:"type"`
  Description string `json:"description"`
  Required bool `json:"required"`
}

type resourceAttributes struct {
  ResourceName string `json: "resourceName"`
  Attributes []attribute `json: "attributes"`
}

type errorMsg struct {
  Msg string `json:"msg"`
}

type jsonConvertible interface {}

func jsonString( obj jsonConvertible ) (s string) {
  jsonObj, err := json.Marshal( obj )
  
  if err != nil {
    s = ""
  } else {
    s = string( jsonObj )
  }

  return
}

// Middleware
func Mongo() martini.Handler {
  return func (c martini.Context ) {
    fmt.Println("Someday I will be a MongoDB connection")
    c.Next()
  }
}

// Handlers start here
func createAttribute( params martini.Params, writer http.ResponseWriter ) (int, string) {
  resource :=  strings.ToLower( params["resource"] )
  writer.Header().Set("Content-Type", "application/json")

	return http.StatusOK, "Pseudo Create assigned to " + resource
}

func getAttributes( params martini.Params, writer http.ResponseWriter ) (int, string) {
    resource :=  strings.ToLower( params["resource"] )
    writer.Header().Set("Content-Type", "application/json")

    if resource  == "tv" {
      resourceAttrs := resourceAttributes{"tv", make([]attribute, 1)}
      resourceAttrs.Attributes[0] = attribute{"", "Location","string", "What facility is the TV located in.", true}
     
      return http.StatusOK, jsonString( resourceAttrs )
    } else {
      return http.StatusNotFound, jsonString( errorMsg{"Resource not found: " + resource} )
    }
}
