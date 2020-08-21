package main

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"time"
)

func main() {
	c, _, err := zk.Connect([]string{"10.35.22.61:2181"}, 10*time.Second)

	if err != nil {
		fmt.Println(err)
	}

	defer c.Close()

	c.Create("/gotest/instance1", []byte("active"), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
}
