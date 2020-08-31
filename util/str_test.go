package util

import (
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	shardingItems := []float32{1, 2, 3, 4, 5, 6}
	items, err := ToStringSlice(shardingItems)
	if err != nil {
		t.Error(err)
	}

	str := strings.Join(items, "@")
	if "1@2@3@4@5@6" != str {
		t.Error("join failed!")
	}
}
