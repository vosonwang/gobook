/*和etcd沟通的包*/
package etcd

import (
	"context"
	"time"
	"github.com/coreos/etcd/client"
	"config"
	"fmt"
)

var kapi client.KeysAPI

func init() {
	cfg := client.Config{
		Endpoints: []string{config.Protocol + config.Host + config.Etcd},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	var err error
	c, err := client.New(cfg)
	if err != nil {
		fmt.Println(err)
	}

	kapi = client.NewKeysAPI(c)
}

func Set(key string, value string) {
	resp, err := kapi.Set(context.Background(), key, value, &client.SetOptions{TTL: time.Hour, Refresh: true})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

func Get(key string) error {

	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println(resp)
		return nil
	}

}
