package assertions

type (
	// Result represents an assertion result
	Result struct {
		err error
	}
)

var (
	success = Result{}
)

// Success returns a success result
func Success() Result { return success }

// Failed returns a Result with an error
func Failed(err error) Result { return Result{err: err} }

// HasFailed returns true if the assertion is not OK and there is an error set
func (r *Result) HasFailed() bool { return r.err != nil }

// HasSucceeded returns true if the assertion is OK
func (r *Result) HasSucceeded() bool { return r.err == nil }

// Err returns the err
func (r *Result) Err() error { return r.err }

// SetErr sets the error field with the given error
func (r *Result) SetErr(err error) Result {
	r.err = err
	return *r
}

// SetResult sets the error field with the given result
func (r *Result) SetResult(other Result) Result {
	r.err = other.err
	return *r
}
