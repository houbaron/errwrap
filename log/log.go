package log

import (
	".."
	"log"
)

func Handle(errs ...error) {
	errwrap.BaseHandle(log.Println, errs...)
}

func ReturnResult(result interface{}, err error) interface{} {
	Handle(err)
	return result
}

func IgnoreResult(_ interface{}, err error) {
	Handle(err)
}
