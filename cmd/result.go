package cmd

type Result interface {
	String() string
}

type ErrorResult struct {
	err error
}

func NewErrorResult(err error) *ErrorResult {
	return &ErrorResult{err: err}
}

func (e *ErrorResult) String() string {
	return e.err.Error()
}
