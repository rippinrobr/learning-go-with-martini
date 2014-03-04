package config

import (
  "github.com/coreos/go-etcd/etcd"  
)

func GetDbConfig( service string ) string {
  etcdClient := etcd.NewClient([]string{"http://127.0.0.1:4001"})
  results, error := etcdClient.Get("mongo/resources", false, false)

  return service
}
