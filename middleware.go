package main

import (
  "fmt"
  "net/http"

  "github.com/rippinrobr/learning-go-with-martini/utils"
  "github.com/codegangsta/martini"
  "labix.org/v2/mgo"
)

func setJsonResponseHeader( writer http.ResponseWriter ) {
  writer.Header().Set("Content-Type", "application/json")
}

// Middleware
var dbInfo utils.MongoInfo

func Mongo() martini.Handler {
  dbInfo = utils.GetDbConfig( "resources" );
  fmt.Println( dbInfo )

  session, err := mgo.Dial( dbInfo.ConnString  )
  if err != nil {
    fmt.Println("[Mongo()] PANIC");
    panic( err )
  }

  return func (c martini.Context ) {
    reqSession := session.Clone()
    c.Map( reqSession.DB( dbInfo.Database ) )
    defer reqSession.Close()
    c.Next()
  }
}
