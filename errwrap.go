package errwrap

type handlerType func(v interface{})

func Handle(handler handlerType, errs ...error) {
	for _, err := range errs {
		if err != nil {
			handler(err)
		}
	}
}

func ReturnResult(handler handlerType, result interface{}, err error) interface{} {
	Handle(handler, err)
	return result
}

func IgnoreResult(handler handlerType, _ interface{}, err error) {
	Handle(handler, err)
}
