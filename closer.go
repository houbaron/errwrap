package errwrap

type Closer interface {
	Close() error
}

func Close(closer Closer) {
	if closer == nil {
		return
	}

	if err := closer.Close(); err != nil {
		Handle(err)
	}
}
