package gresults

import (
	"errors"
	"strconv"
	"testing"
)

func TestGenericsErrorHandler(t *testing.T) {
	want := 1
	errStr := "this failed!"
	strErrorResult := NewResult(&want, &errStr)
	errorHandler := func(err string) {
		if err != errStr {
			t.Errorf("got error %q, wanted %q", err, errStr)
		}
	}
	_ = strErrorResult.OnError(errorHandler).Unwrap(0)
}

func TestPanicIfErrHandlerAbsent(t *testing.T) {
	err := errors.New("this failed!")
	defer func() {
		if r := recover(); r != nil {
			if !errors.Is(r.(error), err) {
				t.Errorf("expected panic with error %q, wanted %q", r, err)
			}
		} else {
			t.Errorf("got nil, expected unwrap to panic")
		}
	}()
	panicingResult := NewResult[int](nil, &err)
	_ = panicingResult.Unwrap(0)
}

func TestSuccessfulUnwrap(t *testing.T) {
	want := 1
	successResult := NewResult[int, error](&want, nil)
	errorHandler := func(err error) {
		t.Errorf("got err %q, wanted nil", err)
	}
	got := successResult.OnError(errorHandler).Unwrap(0)
	if got != want {
		t.Errorf("unwrap failed. got %q, wanted %q", got, want)
	}
}

func TestWrappedCallable(t *testing.T) {
	want := -42
	wrappedAtoiResult := Resultify[int, error](strconv.Atoi, "-42")
	errorHandler := func(err error) {
		t.Errorf("got err %q, wanted nil", err)
	}
	got := wrappedAtoiResult.OnError(errorHandler).Unwrap(0)
	if got != want {
		t.Errorf("unwrap failed. got %q, wanted %q", got, want)
	}
}

func TestUnwrapChaining(t *testing.T) {
	want := -42
	wrappedAtoiResult := Resultify[int, error](strconv.Atoi, "-42")
	errorHandler := func(err error) {
		t.Errorf("got err %q, wanted nil", err)
	}
	wrappedAtoiResult.OnError(errorHandler).UnwrapAndThen(0, func(got int) (i int, e error) {
		if got != want {
			t.Errorf("unwrap failed. got %q, wanted %q", got, want)
		}
		return
	})
}
