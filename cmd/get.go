package cmd

import (
	"errors"
	"fmt"

	"go.themis.run/themisclient/pb"
)

var errArgsEmpty = errors.New("args can not be empty")
var errArgsFormat = errors.New("args format error")

func init() {
	Register("get", GetKV)
}

type GetResult struct {
	kv *pb.KV
}

func NewGetResult(kv *pb.KV) *GetResult {
	return &GetResult{
		kv: kv,
	}
}

func (r *GetResult) String() string {
	return fmt.Sprintf("key: %s\nvalue: %s\ncreate_time: %s\nttl: %d\n",
		r.kv.Key, r.kv.Value, transforTimestamp(r.kv.CreateTime), r.kv.Ttl)
}

func GetKV(args ...string) Result {
	if len(args) == 0 {
		return NewErrorResult(errArgsEmpty)
	}

	if len(args) != 1 {
		return NewErrorResult(errArgsFormat)
	}

	c := GetThemisClient()
	kv, err := c.Get(args[0])
	if err != nil {
		return NewErrorResult(err)
	}

	return NewGetResult(kv)
}
