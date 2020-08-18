package errwrap

type (
	HandlerType func(err error)

	Closer interface {
		Close() error
	}

	Wrapper struct {
		Handler HandlerType
	}
)

func (wrapper *Wrapper) handle(errs ...error) {
	if wrapper.Handler == nil {
		return
	}

	for _, err := range errs {
		if err != nil {
			wrapper.Handler(err)
		}
	}
}

func (wrapper *Wrapper) ReturnResult(result interface{}, err error) interface{} {
	wrapper.handle(err)
	return result
}

func (wrapper *Wrapper) IgnoreResult(_ interface{}, err error) {
	wrapper.handle(err)
}

func (wrapper *Wrapper) Close(closer Closer) {
	if closer != nil {
		wrapper.handle(closer.Close())
	}
}

var (
	PanicWrapper = &Wrapper{Handler: func(err error) {
		panic(err)
	}}

	PrintLnWrapper = &Wrapper{Handler: func(err error) {
		println(err)
	}}

	DoNothingWrapper = &Wrapper{Handler: nil}
)
