package main

import (
  "encoding/json"
  "net/http"

  "github.com/codegangsta/martini-contrib/binding"
  "github.com/rippinrobr/learning-go-with-martini/utils"
)

// Structs and related code
type attribute struct {
  Name string `json:"name"`
  DataType string `json:"type"`
  Description string `json:"description"`
  Required bool `json:"required"`
}

func (attr *attribute) Validate( errors *binding.Errors, req *http.Request ) {
  if attr.Name == "" {
    errors.Overall["missing-requirement"] = "name is a required field";
  }

  if attr.DataType == "" {
    attr.DataType = "string"
  }
}

type resourceAttributes struct {
  Resource string `json: "resource"` 
  Attributes []attribute `json: "attributes"`
}

type errorMsg struct {
  Msg string `json:"msg"`
}

func jsonString( obj utils.JsonConvertible ) (s string) {
  jsonObj, err := json.Marshal( obj )
  
  if err != nil {
    s = ""
  } else {
    s = string( jsonObj )
  }

  return
}
