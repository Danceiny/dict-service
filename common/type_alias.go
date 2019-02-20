package common

import (
	"strconv"
)

type BID interface {
	String() string
	Empty() BID
}

type NodeLevel interface {
	String()
	Val() int
}

type NodeId struct {
	v interface{}
}

type INT int64

type STRING string

func (s STRING) String() string {
	return string(s)
}
func (s STRING) Empty() BID {
	return STRING("")
}
func (id INT) String() string {
	return strconv.FormatInt(int64(id), 10)
}
func (id INT) Empty() BID {
	return INT(0)
}
func NodeLevelAgtB(a, b *int) bool {
	v := CompareNodeLevel(a, b)
	if v > 0 {
		return true
	} else {
		return false
	}
}
func NodeLevelAltB(a, b *int) bool {
	v := CompareNodeLevel(a, b)
	if v < 0 {
		return true
	} else {
		return false
	}
}

// todo
func CompareNodeLevel(a, b *int) int {
	return 0
}
