package main

import (
    "code.qschou.com/go_micro/etcd"
    "code.qschou.com/go_micro/lease"
    "context"
    "fmt"
    "github.com/coreos/etcd/clientv3"
    "time"
)

func main() {
    var (
        config  clientv3.Config
        client  *clientv3.Client
        err     error
        kv      clientv3.KV
        getResp *clientv3.GetResponse
    )
    //ctx,c := context.WithCancel(context.Background())
    //go printf(ctx)

    //time.Sleep(5*time.Second)
    //c()

    config = clientv3.Config{
        Endpoints:   []string{"123.57.51.133:2379"},
        DialTimeout: 5 * time.Second,
    }

    // 建立一个客户端
    if client, err = clientv3.New(config); err != nil {
        fmt.Println(err)
        return
    }


    for i:=0;i<100;i++ {
        go etcd.NewLease(client,"zzz","uuu")
        time.Sleep(11*time.Second)
    }



    // 用于读写etcd的键值对
    kv = clientv3.NewKV(client)

    // 写入
    kv.Put(context.TODO(), "name1", "lesroad")
    kv.Put(context.TODO(), "name2", "haha")

    // 读取name为前缀的所有key
    if getResp, err = kv.Get(context.TODO(), "name", clientv3.WithPrefix()); err != nil {
        fmt.Println(err)
        return
    } else {
        // 获取成功
        fmt.Println(getResp.Kvs)
    }

    // 删除name为前缀的所有key
    if _, err = kv.Delete(context.TODO(), "name", clientv3.WithPrevKV()); err != nil {
        fmt.Println(err)
        return
    }
}

func printf(ctx context.Context)  {
    for {
        fmt.Println("a\n")
        time.Sleep(time.Second)
    }

    select {
    case <-ctx.Done():
        fmt.Println("bbbbbbbbbb")
    }

}