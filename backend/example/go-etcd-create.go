package main

//https://godoc.org/go.etcd.io/etcd/clientv3

import (
        "time"
	"context"
        "fmt"

	"go.etcd.io/etcd/clientv3"
)

func main() {
    cli, err := clientv3.New(clientv3.Config{
    	Endpoints:   []string{"172.20.57.58:2379"},
    	DialTimeout: 5 * time.Second,
    })
    if err != nil {
    	// handle error!
    }
    defer cli.Close()

    //ctx, cancel := context.WithTimeout(context.Background(), timeout)
    //resp, err := cli.Put(ctx, "sample_key", "sample_value")
    resp, err := cli.Put(context.TODO(), "/skydns/local/skydns/www1", "{\"host\":\"1.1.1.3\",\"ttl\":60}")
    //resp, err := cli.Put(context.TODO(), "/skydns/local/skydns/www1", "1.1.1.4")
    //cancel()
    if err != nil {
        // handle error!
    }

    fmt.Println(resp)
}
