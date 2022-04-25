package gresults

import (
	"reflect"
)

type Result[T any, E any] interface {
	OnError(errHandler func(err E)) Result[T, E]
	Unwrap(defaultValue T) T
}

type result[T any, E any] struct {
	data       *T
	err        *E
	errHandler func(err E)
}

func (res result[T, E]) OnError(errHandler func(err E)) Result[T, E] {
	res.errHandler = errHandler
	return res
}

func (res result[T, E]) Unwrap(defaultValue T) T {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case E:
				err := r.(E)
				if res.errHandler != nil {
					res.errHandler(err)
				} else {
					panic(err)
				}
			default:
				panic(r)
			}
		}
	}()

	if res.err != nil {
		panic(*res.err)
	}

	if res.data == nil {
		return defaultValue
	}
	return *res.data
}

func NewResult[T any, E any](data *T, err *E) Result[T, E] {
	return result[T, E]{
		data: data,
		err:  err,
	}
}

func Resultify[T any, E any](fn interface{}, args ...interface{}) Result[T, E] {
	functionDef := reflect.TypeOf(fn)
	var inputArgs []reflect.Value
	for i := 0; i < functionDef.NumIn(); i++ {
		inputArgs = append(inputArgs, reflect.ValueOf(args[i]))
	}

	f := reflect.ValueOf(fn)
	r := f.Call(inputArgs)

	var data *T
	var err *E
	if r[0].Interface() != nil {
		t := r[0].Interface().(T)
		data = &t
	} else {
		data = nil
	}

	if r[1].Interface() != nil {
		e := r[1].Interface().(E)
		err = &e
	} else {
		err = nil
	}

	return NewResult(data, err)
}
