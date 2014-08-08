package main

import (
  "net/http"

  "github.com/codegangsta/martini-contrib/binding"
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
    errors.Overall["missing-requirement"] = " name is a required field";
  }

  if attr.DataType == "" {
    attr.DataType = "string"
  }
}

type resourceAttributes struct {
	Resource string `json: "resource"` 
  Attributes []attribute `json: "attributes"`
}

type resource struct {
	ResourceName string `form:"resourceName"`
}

type errorMsg struct {
  Msg string `json:"msg"`
}
