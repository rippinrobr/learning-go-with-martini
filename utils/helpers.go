package utils

import (
  "fmt"
  "encoding/json"
)

type Keyval struct {
  Key string
  Val string
}

type MongoInfo struct {
  Collection string `json:"collection"`
  ConnString string `json:"connString"`
  Database string `json:"database"`
}

func (mi MongoInfo) String() string {
  return fmt.Sprintf("[MongoInfo]\nConnString: %s\nCollection: %s\nDatabase: %s\n", mi.ConnString, mi.Collection, mi.Database)
}

type ServiceDescription struct {
  Name string `json:"name"`
  Server string `json:"server"`
  Description string `json:"description"`
}

func (sd ServiceDescription) CreateEtcdKeyValues() (Keyval, Keyval) {
  return Keyval{fmt.Sprintf("services/%s/Server", sd.Name ), sd.Server}, Keyval{fmt.Sprintf("services/%s/Description", sd.Name), sd.Description}
}

func (sd ServiceDescription) String() string {
  return fmt.Sprintf("[ServiceDescription]\nName: %s\nServer: %s\nDescription: %s\n", sd.Name, sd.Server, sd.Description)
}

func JsonString( obj JsonConvertible ) (s string) {
  jsonObj, err := json.Marshal( obj )
  
  if err != nil {
    s = ""
  } else {
    s = string( jsonObj )
  }

  return
}
