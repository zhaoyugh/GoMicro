package etcd

import (
    "context"
    "fmt"
    "github.com/coreos/etcd/clientv3"
    "time"
)

func NewLease(client *clientv3.Client,key string,val string){
    lease := clientv3.NewLease(client)
    var (
        leaseId clientv3.LeaseID
        leaseGrantResp *clientv3.LeaseGrantResponse
        err error
        keepRespChan <- chan *clientv3.LeaseKeepAliveResponse
        //putResp *clientv3.PutResponse
        //getResp *clientv3.GetResponse
        )

    if leaseGrantResp,err = lease.Grant(context.TODO(),25);err != nil{
        fmt.Println(err)
        return
    } else {
        leaseId = leaseGrantResp.ID
    }

    cancelCtx,cancelFunc := context.WithCancel(context.TODO())
    if keepRespChan,err = lease.KeepAlive(cancelCtx,leaseId);err != nil{
        return
    }

    go func() {
        for {
            var keepResp *clientv3.LeaseKeepAliveResponse
            select {
            case keepResp = <-keepRespChan:
                if keepResp == nil{
                    return
                }

                fmt.Println(keepResp)
            }
        }
    }()

    kv := clientv3.NewKV(client)
    txn := kv.Txn(context.TODO())

    txn.If(clientv3.Compare(clientv3.CreateRevision(key),"=",0)).
        Then(clientv3.OpPut(key,val,clientv3.WithLease(leaseId))).
        Else(clientv3.OpGet(key))

     txnRes,err := txn.Commit()
     if err != nil {
         fmt.Println(err)
     }

     if txnRes.Succeeded{
         fmt.Println("锁成功创建",txnRes.Header.Revision)
         time.Sleep(10*time.Second)
         cancelFunc()
         if leaseRevokeRes,err := lease.Revoke(context.TODO(),leaseId);err != nil{
             return
         }else {
             fmt.Println(leaseRevokeRes)
         }
     } else {
         fmt.Println("锁已经被占用",txnRes.Header.Revision)
     }

    //if putResp,err = kv.Put(context.TODO(),"zzz",val,clientv3.WithLease(leaseId));err != nil{
    //    fmt.Println(err)
    //    return
    //}
    //
    //fmt.Println("写入成功:", putResp.Header.Revision)
    //
    //for {
    //    if getResp,err = kv.Get(context.TODO(),"zzz");err != nil{
    //        fmt.Println(err)
    //        return
    //    }
    //
    //    if getResp.Count == 0{
    //        fmt.Println("已经过期")
    //        break
    //    }
    //    fmt.Println("还没过期:",getResp.Kvs)
    //    time.Sleep(time.Second)
    //}
}