package utils

import (
  "fmt"
  "os"

  "github.com/coreos/go-etcd/etcd"  
)

type ServiceDescription struct {
  Name string `json:"name"`
  URI string `json:"uri"`
  Description string `json:"description"`
}

type MongoInfo struct {
  Collection string `json:"collection"`
  ConnString string `json:"connString"`
  Database string `json:"database"`
}

func (mi MongoInfo) String() string {
  return fmt.Sprintf("[MongoInfo]\nConnString: %s\nCollection: %s\nDatabase: %s\n", mi.ConnString, mi.Collection, mi.Database)
}

func GetDbConfig( service string ) MongoInfo {
  etcdServer := os.Getenv("GD_ETCD_SERVER")
  if etcdServer == "" {
    etcdServer = "http://127.0.0.1:4001"
  }

  etcdClient := etcd.NewClient([]string{ etcdServer }) 
  results, error := etcdClient.Get("mongo/" + service , true, false)

  if error != nil {
    panic( error );
  }
  
  mongoInfo := MongoInfo{ results.Node.Nodes[0].Value, results.Node.Nodes[1].Value, results.Node.Nodes[2].Value }
   
  return mongoInfo
}

func RegisterService( serviceDesc ServiceDescription ) bool {
  return true
}
