package config

import (
  "fmt"
  "github.com/coreos/go-etcd/etcd"  
)

func printResults( results etcd.Response ) {
  fmt.Println("Action: ", results.Action);
  fmt.Println("PrevNode: ", results.PrevNode);
  fmt.Println("EtcdIndex: ", results.EtcdIndex);
  fmt.Println("RaftIndex: ", results.RaftIndex);
  fmt.Println("RaftTerm: ", results.RaftTerm);
  conn := results.Node.Nodes[0]
  coll := results.Node.Nodes[1]
  fmt.Printf( "%s => %s\n", conn.Key, conn.Value)
  fmt.Printf( "%s => %s\n", coll.Key, coll.Value)
}

type MongoInfo struct {
  ConnString string `json:"connString"`
  Collection string `json:"collection"`
}

func (mi MongoInfo) String() string {
  return fmt.Sprintf("[MongoInfo]\nConnString: %s\nCollection: %s\n", mi.ConnString, mi.Collection)
}

func GetDbConfig( service string ) MongoInfo {
  etcdClient := etcd.NewClient([]string{"http://127.0.0.1:4001"}) 
  results, error := etcdClient.Get("mongo/resources", false, false)

  if error != nil {
    panic( error );
  }
  // if you want to see what the other properties look like 
  // in the etcd.Response struct uncomment the line below
  // printResults( *results )
  
  mongoInfo := MongoInfo{}
  for _, entry := range results.Node.Nodes {
    if entry.Key == "/mongo/resources/conn" {
      mongoInfo.ConnString = entry.Value
    } else {
      mongoInfo.Collection = entry.Value
    }
  }
   
  return mongoInfo
}
