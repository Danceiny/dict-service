package common

import (
    "strconv"
)

type BID interface {
    String() string
    Empty() BID
    Equal(v BID) bool
}

type NodeLevel interface {
    String()
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

func (s STRING) Equal(v BID) bool {
    return s == v.(STRING)
}

func (id INT) String() string {
    return strconv.FormatInt(int64(id), 10)
}

func (id INT) Empty() BID {
    return INT(0)
}

func (id INT) Equal(v BID) bool {
    return id == v.(INT)
}
