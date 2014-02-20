package main

import ( 
  "net/http"
  "strings"
  "encoding/json"
  "github.com/codegangsta/martini"
  "labix.org/v2/mgo"  // brings the mgo library into my code
  "labix.org/v2/mgo/bson" 
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
  session, err := mgo.Dial( "localhost/goattrs" )
  if err != nil {
    panic( err )
  }

  return func (c martini.Context ) {
    reqSession := session.Clone()
    c.Map( reqSession.DB( "goattrs" ) )
    defer reqSession.Close()
    c.Next()
  }
}

// Handlers start here
func createAttribute( params martini.Params, writer http.ResponseWriter ) (int, string) {
  resource :=  strings.ToLower( params["resource"] )
  writer.Header().Set("Content-Type", "application/json")

  return http.StatusOK, "Pseudo Create assigned to " + resource
}

func getAttributes( params martini.Params, writer http.ResponseWriter, db *mgo.Database) (int, string) {
  resource :=  strings.ToLower( params["resource"] )
  writer.Header().Set("Content-Type", "application/json")

  var attrs []resourceAttributes
  db.C("resource_attributes").Find(bson.M{"resource": resource }).All(&attrs);

  if attrs != nil {
    return http.StatusOK, jsonString( attrs ) //resourceAttrs )
  } else {
    return http.StatusNotFound, jsonString( errorMsg{"No attributes found for the resource: " + resource} )
  }
}
