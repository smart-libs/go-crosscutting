package transaction

func MakeError(panicArg any, optionalCurrentError error) (err error) {
	err = optionalCurrentError
	if panicArg != nil {
		var isErr bool
		if err, isErr = panicArg.(error); !isErr {
			err = WrapAsPanicError(err)
		}
	}
	return
}
