package main

import (
    "context"
    "fmt"
    "github.com/coreos/etcd/clientv3"
    //"github.com/coreos/etcd/clientv3"
    "time"
)

//获取指定“目录”下的kvs：
func main() {
    var (
        config  clientv3.Config
        client  *clientv3.Client
        err     error
        kv      clientv3.KV
        getResp *clientv3.GetResponse
    )

    //客户端配置
    config = clientv3.Config{
        Endpoints:   []string{"123.57.51.133:2379"}, //集群列表
        DialTimeout: 5 * time.Second,
    }
    //建立客户端
    if client, err = clientv3.New(config); err != nil {
        fmt.Println(err)
        return
    }
    //用于读写etcd的键值对
    kv = clientv3.NewKV(client)

    //读取/cron/jobs/为前缀的所有key
    if getResp, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
        fmt.Println(err)
    } else { //获取成功，我们遍历所有kvs
        fmt.Println(getResp.Kvs)
    }
}