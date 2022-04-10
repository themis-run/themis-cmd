package cmd

import (
	"fmt"

	"go.themis.run/themisclient/pb"
)

func init() {
	Register("find", Find)
}

type FindResult struct {
	kvs []*pb.KV
}

func NewFindResult(kvs []*pb.KV) *FindResult {
	return &FindResult{}
}

func (f *FindResult) String() string {
	str := "-----------------------------\n"
	for _, kv := range f.kvs {
		str += fmt.Sprintf("key: %s\nvalue: %s\ncreate_time: %s\nttl: %d\n",
			kv.Key, kv.Value, transforTimestamp(kv.CreateTime), kv.Ttl)
		str += "-----------------------------\n"
	}

	return str
}

func Find(args ...string) Result {
	if len(args) == 0 {
		return NewErrorResult(errArgsEmpty)
	}

	if len(args) != 1 {
		return NewErrorResult(errArgsFormat)
	}

	c := GetThemisClient()

	if args[0] == "all" {
		kvs, err := c.ListAllKV()
		if err != nil {
			return NewErrorResult(err)
		}

		return NewFindResult(kvs)
	}

	kvs, err := c.SearchByPrefix(args[0])
	if err != nil {
		return NewErrorResult(err)
	}

	return NewFindResult(kvs)
}
