package main

import ( 
  "strings"
  "net/http"

  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/binding"
  "github.com/rippinrobr/learning-go-with-martini/utils"
  "github.com/codegangsta/martini-contrib/render"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson" 
)

// Handlers start here
func getAttributes( params martini.Params, writer http.ResponseWriter, db *mgo.Database) (int, string) {
  resource :=  strings.ToLower( params["resource"] )
  setJsonResponseHeader( writer )

  attrs := resourceAttributes{}
  err   := db.C( dbInfo.Collection ).Find(bson.M{"resource": resource }).One(&attrs)

  if err == nil {
    return http.StatusOK, utils.JsonString( attrs )
  } else {
    return http.StatusNotFound, utils.JsonString( errorMsg{"No attributes found for the resource: " + resource} )
  }
}

func addAttribute( attr attribute, err binding.Errors, params martini.Params, writer http.ResponseWriter, db *mgo.Database) (int, string)  {
  setJsonResponseHeader( writer )

  if err.Count() > 0 {
    return http.StatusConflict, utils.JsonString( errorMsg{ err.Overall["missing-requirement"] } )
  }
  
  resource :=  strings.ToLower( params["resource"] )
  query  := bson.M{"resource": resource }
  update := mgo.Change{  Upsert: true, Update: bson.M{ "$addToSet" : bson.M{ "attributes" : attr }} }

  if _, dbErr := db.C( dbInfo.Collection ).Find( query ).Apply( update, &attr); dbErr != nil {
    return http.StatusConflict, utils.JsonString( errorMsg{ dbErr.Error() } )
  }

  return http.StatusOK, "{}"
}

func displayResourcesPage( r render.Render, db *mgo.Database) {
  attrs := []resourceAttributes{}
  err   := db.C( dbInfo.Collection ).Find(bson.M{}).All(&attrs)

  if err == nil {  
    r.HTML(200, "resources", attrs)
  } else {
    r.HTML(404, "404", err) 
  }
}
