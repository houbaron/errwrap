package errwrap

type handlerType func(v interface{})

var handler handlerType = panic

func SetHandler(f handlerType) {
	handler = f
}

func baseHandle(handle handlerType, errs ...error) {
	for _, err := range errs {
		if err != nil {
			handle(err)
		}
	}
}

func Handle(errs ...error) {
	baseHandle(handler, errs...)
}

func ReturnResult(result interface{}, err error) interface{} {
	Handle(err)
	return result
}

func IgnoreResult(_ interface{}, err error) {
	Handle(err)
}
