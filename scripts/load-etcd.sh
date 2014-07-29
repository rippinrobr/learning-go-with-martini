#!/bin/bash

res1=`curl -L http://127.0.0.1:4001/v2/keys/mongo/resources/conn -XPUT -d value="localhost/goattrs"`
res2=`curl -L http://127.0.0.1:4001/v2/keys/mongo/resources/collection -XPUT -d value="resource_attributes"`
res3=`curl -L http://127.0.0.1:4001/v2/keys/mongo/resources/database -XPUT -d value="goattrs"`
res4=`curl -L http://127.0.0.1:4001/v2/keys/services -XPUT`
echo "Connection Results"
echo $res1
echo ""
echo "Collection results"
echo $res2
echo ""
echo "Database results"
echo $res3
echo "Service results"
echo $res4
