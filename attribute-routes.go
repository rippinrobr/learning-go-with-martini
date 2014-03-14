package main

import ( 
  "strings"
  "net/http"

  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson" 
)

// Handlers start here
func createAttribute( params martini.Params, writer http.ResponseWriter ) (int, string) {
  resource :=  strings.ToLower( params["resource"] )
  setJsonResponseHeader( writer ) //.Header().Set("Content-Type", "application/json")

  return http.StatusOK, "Pseudo Create assigned to " + resource
}

func getAttributes( params martini.Params, writer http.ResponseWriter, db *mgo.Database) (int, string) {
  resource :=  strings.ToLower( params["resource"] )
  setJsonResponseHeader( writer ) //.Header().Set("Content-Type", "application/json")

  attrs := resourceAttributes{}
  err   := db.C( dbInfo.Collection ).Find(bson.M{"resource": resource }).One(&attrs)

  if err == nil {
    return http.StatusOK, jsonString( attrs )
  } else {
    return http.StatusNotFound, jsonString( errorMsg{"No attributes found for the resource: " + resource} )
  }
}

func addAttribute( attr attribute, err binding.Errors, params martini.Params, writer http.ResponseWriter, db *mgo.Database) (int, string)  {
  setJsonResponseHeader( writer ) //.Header().Set("Content-Type", "application/json")

  if err.Count() > 0 {
    return http.StatusConflict, jsonString( errorMsg{ err.Overall["missing-requirement"] } )
  }
  
  resource :=  strings.ToLower( params["resource"] )
  query  := bson.M{"resource": resource }
  update := mgo.Change{  Upsert: true, Update: bson.M{ "$addToSet" : bson.M{ "attributes" : attr }} }

  if _, dbErr := db.C( dbInfo.Collection ).Find( query ).Apply( update, &attr); dbErr != nil {
    return http.StatusConflict, jsonString( errorMsg{ dbErr.Error() } )
  }

  return http.StatusOK, "{}"
}
