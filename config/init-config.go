package config

import( 
	"fmt"
	"os"
	"github.com/rippinrobr/learning-go-with-martini/utils"
	"github.com/coreos/go-etcd/etcd"
)

func getEtcdClient() *etcd.Client {
  etcdServer := os.Getenv("GD_ETCD_SERVER")
  if etcdServer == "" {
    etcdServer = "http://127.0.0.1:4001"
  }

  return etcd.NewClient([]string{ etcdServer }) 
}

func GetDbConfig( service string ) utils.MongoInfo {
  etcdClient := getEtcdClient() 
  results, error := etcdClient.Get("mongo/" + service , true, false)

  if error != nil {
    panic( error );
  }
  
  mongoInfo := utils.MongoInfo{ results.Node.Nodes[0].Value, results.Node.Nodes[1].Value, results.Node.Nodes[2].Value }
   
  return mongoInfo
}

func RegisterService( serviceDesc utils.ServiceDescription ) bool {
  fmt.Println( serviceDesc )

  serverKV, descKV := serviceDesc.CreateEtcdKeyValues()
  fmt.Println( serverKV, descKV )
  etcdClient := getEtcdClient()
  _, err := etcdClient.Set(serverKV.Key, serverKV.Val, 300)

  if err != nil {
    _, err := etcdClient.Set(descKV.Key, descKV.Val, 300)
    return err == nil
  } else { 
    return false
  }
  return true
}
