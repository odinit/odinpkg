package main

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/odinit/odinpkg/conf"
	"github.com/odinit/odinpkg/consul"
	"github.com/odinit/odinpkg/file"
)

var p string = "/Volumes/OdinDisk/Work/Project/Bigdata/Code/Source/源码-20220921/microbigdata"

func main() {
	//deletea()
	//finda()
	//getexts()
	vipercheck()
	//consulCheck()
}

func getexts() {
	a, err := file.ExtSum(p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(a)
}

func deletea() {
	err := file.DeleteByExt(p, ".log")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func finda() {
	b, err := file.FindByExt(p, ".conf")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(b)
}

func vipercheck() {
	type CONF struct {
		Host string
		Port int

		Host1 string
		Port1 int
		TCP   string
	}

	var c = new(CONF)
	err := conf.Init("a.yaml", c)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(c)
	fmt.Println(c.TCP)
}

func consulCheck() {
	err := consul.GlobalClientWithDefaultConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = consul.Client.KV().Delete("name", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = consul.Client.KV().Put(&consulApi.KVPair{Key: "self/name", Value: []byte("zxx")}, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//_, err = consul.Client.KV().Put(&consulApi.KVPair{Key: "name", Value: []byte("aaa")}, nil)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	//kv, _, err := consul.Client.KV().Get("name", nil)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println("kv-", kv.Key,":", string(kv.Value))

	ks, _, err := consul.Client.KV().Keys("self", "", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ks:", ks)

	kvs, _, err := consul.Client.KV().List("self", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i, kv_ := range kvs {
		fmt.Printf("%d-%s:%s\n", i, kv_.Key, string(kv_.Value))
	}

}
