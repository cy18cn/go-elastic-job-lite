package main

import (
	"testing"
	"time"

	"github.com/go-zookeeper/zk"
)

func Test_zk(t *testing.T) {
	c, _, err := zk.Connect([]string{"10.35.22.61:2181"}, 10*time.Second)

	if err != nil {
		t.Error(err)
	}
	defer c.Close()

	//c.Create("/gotest", nil, 0, zk.WorldACL(zk.PermAll))

	// create and watch
	children, stat, ch, err := c.ChildrenW("/gotest")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v %+v\n", children, stat)
	e := <-ch
	data, _, err := c.Get(e.Path)
	t.Logf("%+v\n", string(data))
}

func callback() {

}

func Test_ZKCreate(t *testing.T) {
	c, _, err := zk.Connect([]string{"10.35.22.61:2181"}, 10*time.Second)

	if err != nil {
		t.Error(err)
	}

	defer c.Close()
	c.Create("/gotest/instance1", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
}
