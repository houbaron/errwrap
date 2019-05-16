package panic

import (
	"github.com/houbaron/errwrap"
	"log"
)

func Handle(errs ...error) {
	errwrap.BaseHandle(log.Panic, errs...)
}

func ReturnResult(result interface{}, err error) interface{} {
	Handle(err)
	return result
}

func IgnoreResult(_ interface{}, err error) {
	Handle(err)
}
