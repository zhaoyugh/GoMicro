package main

import (
    "context"
    "fmt"
    "github.com/coreos/etcd/clientv3"
    //"github.com/coreos/etcd/clientv3"
    "time"
)

var (
    config  clientv3.Config
    client  *clientv3.Client
    err     error
    kv      clientv3.KV
    putResp *clientv3.PutResponse
)
//设置k---v
func main() {
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
    //操作k-v(用于读写etcd的键值对)
    kv = clientv3.NewKV(client)
    //context.TODO()代表什么都不做，占位就可以了
    //不在这里指定clientv3.WithPrevKV()，则不能得到putResp.PrevKv.Key,putResp.PrevKv.Value
    if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "world3", clientv3.WithPrevKV()); err != nil {
        fmt.Println(err)
    } else {
        //输出版本号
        fmt.Println("Revision:", putResp.Header.Revision)
        if putResp.PrevKv != nil {
            //输出上一个k-v
            fmt.Println(string(putResp.PrevKv.Key), ":", string(putResp.PrevKv.Value))
        }
    }
}