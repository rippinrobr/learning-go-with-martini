package main

import ( 
  "fmt"
  "strings"
  "encoding/json"
  "net/http"

  "github.com/rippinrobr/learning-go-with-martini/config"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson" 
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
  fmt.Println( config.GetDbConfig( "resources" ) );
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

  attrs := resourceAttributes{}
  err   := db.C("resource_attributes").Find(bson.M{"resource": resource }).One(&attrs)

  if err == nil {
    return http.StatusOK, jsonString( attrs )
  } else {
    return http.StatusNotFound, jsonString( errorMsg{"No attributes found for the resource: " + resource} )
  }
}

func addAttribute( attr attribute, err binding.Errors, params martini.Params, writer http.ResponseWriter, db *mgo.Database) (int, string)  {
  writer.Header().Set("Content-Type", "application/json")

  if err.Count() > 0 {
    return http.StatusConflict, jsonString( errorMsg{ err.Overall["missing-requirement"] } )
  }
  
  resource :=  strings.ToLower( params["resource"] )
  query  := bson.M{"resource": resource }
  update := mgo.Change{  Upsert: true, Update: bson.M{ "$addToSet" : bson.M{ "attributes" : attr }} }

  if _, dbErr := db.C("resource_attributes").Find( query ).Apply( update, &attr); dbErr != nil {
    return http.StatusConflict, jsonString( errorMsg{ dbErr.Error() } )
  }

  return http.StatusOK, "{}"
}
