package service

import "errors"

type ArgException struct {
    error
}

func NewArgException(msg string) *ArgException {
    return &ArgException{errors.New(msg)}
}

func ThrowArgException(msg string) {
    panic(NewArgException(msg))
}
