package errwrap

type Closer interface {
	Close() error
}

func Close(handler handlerType, closer Closer) {
	if closer == nil {
		return
	}

	if err := closer.Close(); err != nil {
		Handle(handler, err)
	}
}
