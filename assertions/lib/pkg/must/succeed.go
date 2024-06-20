package must

func Succeed(err error) {
	if err != nil {
		panic(err)
	}
}

// SucceedWith1 must succeed with 1 argument after the error argument.
func SucceedWith1[T any](t T, err error) T {
	Succeed(err)
	return t
}
