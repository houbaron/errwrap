package errwrap

func BaseHandle(handle func(...interface{}), errs ...error) {
	for _, err := range errs {
		if err != nil {
			handle(err)
		}
	}
}
