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
}
type NodeId int64

type String string

func (s String) String() string {
    return string(s)
}
func (s String) Empty() BID {
    return String("")
}
func (id NodeId) String() string {
    return strconv.FormatInt(int64(id), 10)
}
func (id NodeId) Empty() BID {
    return NodeId(0)
}
func NodeLevelAgtB(a, b NodeLevel) bool {
    v := CompareNodeLevel(a, b)
    if v > 0 {
        return true
    } else {
        return false
    }
}
func NodeLevelAltB(a, b NodeLevel) bool {
    v := CompareNodeLevel(a, b)
    if v < 0 {
        return true
    } else {
        return false
    }
}

// todo
func CompareNodeLevel(a, b NodeLevel) int {
    return 0
}
