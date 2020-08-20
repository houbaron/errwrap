package errwrap

import "errors"

type (
	HandlerType func(err error)

	Closer interface {
		Close() error
	}

	Wrapper struct {
		handler HandlerType
	}
)

func (wrapper *Wrapper) HandleErrors(errs ...error) {
	if wrapper.handler == nil {
		return
	}

	for _, err := range errs {
		if err != nil {
			wrapper.handler(err)
		}
	}
}

func (wrapper *Wrapper) ReturnResult(result interface{}, err error) interface{} {
	wrapper.HandleErrors(err)
	return result
}

func (wrapper *Wrapper) IgnoreResult(_ interface{}, err error) {
	wrapper.HandleErrors(err)
}

func (wrapper *Wrapper) IsResultNil(val interface{}, msg string) interface{} {
	if val == nil {
		wrapper.handler(errors.New(msg))
	}

	return val
}

func (wrapper *Wrapper) Close(closer Closer) {
	if closer != nil {
		wrapper.HandleErrors(closer.Close())
	}
}

func New(handler HandlerType) *Wrapper {
	return &Wrapper{handler: handler}
}

var (
	PanicWrapper = &Wrapper{handler: func(err error) {
		panic(err)
	}}

	PrintLnWrapper = &Wrapper{handler: func(err error) {
		println(err.Error())
	}}

	DoNothingWrapper = &Wrapper{handler: nil}
)
