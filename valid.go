package valid

//Result represents a testing result
type Result struct {
	Message string
	Status  bool
}

//FailResult creates a failed result with message
func FailResult(message string) Result {
	return Result{
		Message: message,
		Status:  false,
	}
}

//OkResult creates a successful result
func OkResult() Result {
	return Result{
		Status: true,
	}
}

//PrintIfFail executes a print function, like t.Error, when status is false
func (r Result) PrintIfFail(fn func(...interface{})) {
	if !r.Status {
		fn(r.Message)
	}
}
