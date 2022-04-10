package cmd

import (
	"strconv"
	"time"
)

func init() {
	Register("set", SetKV)
}

type SetResult struct {
}

func NewSetResult() *SetResult {
	return &SetResult{}
}

func (s *SetResult) String() string {
	return "success!"
}

func SetKV(args ...string) Result {
	if len(args) == 0 {
		return NewErrorResult(errArgsEmpty)
	}

	if len(args) != 2 || len(args) != 3 {
		return NewErrorResult(errArgsFormat)
	}

	c := GetThemisClient()
	var ttl int64 = 0
	var err error

	if len(args) == 3 {
		ttl, err = strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			return NewErrorResult(err)
		}
	}

	if err := c.SetWithExpireTime(args[0], args[1], time.Duration(ttl)); err != nil {
		return NewErrorResult(err)
	}

	return NewSetResult()
}
