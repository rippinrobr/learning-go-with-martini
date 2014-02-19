package main

// loading in the Martini package
import (
  "net/http"  // this will allow us to use http.StatusOK and http.StatusNotFound instead of 200 and 404
  "strings"   // I’m adding this so I can ensure that we are comparing lower case strings
  "encoding/json" 
  "github.com/codegangsta/martini"
)

// defining the struct that I will use to create my error message that
// will be 'jsonified' and sent back to the caller.  The `json:"msg"`
// tells the json encoder what name to use for this property in the json
// object
type ErrorMsg struct {
  Msg string `json:"msg"`
}

func (e ErrorMsg) String() (s string) {
  jsonObj, err := json.Marshal(e)
  
  if err != nil {
    s = ""
    return
  }

  s = string( jsonObj )
  return 
}

func main() {
  // := is a short variable declaration it types and instantiates the var 'on the fly'
  m := martini.Classic()

  m.Get("/attributes/:resource", func( params martini.Params ) (int, string) {
    resource :=  strings.ToLower( params["resource"] )

    if resource  == "tv" {
      return http.StatusOK, "a TV attributes object will be returned here"
    } else {
      return http.StatusNotFound, ErrorMsg{"Resource not found: " + resource}.String()
    }
  })

  m.Run()
}
