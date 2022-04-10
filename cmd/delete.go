package cmd

func init() {
	Register("del", Del)
	Register("delete", Del)
}

type DelResult struct {
}

func NewDelResult() *DelResult {
	return &DelResult{}
}

func (d *DelResult) String() string {
	return "success!"
}

func Del(args ...string) Result {
	if len(args) == 0 {
		return NewErrorResult(errArgsEmpty)
	}

	if len(args) != 2 {
		return NewErrorResult(errArgsFormat)
	}

	c := GetThemisClient()
	if err := c.Delete(args[0]); err != nil {
		return NewErrorResult(err)
	}

	return NewDelResult()
}
