package common

import (
	"strconv"
)

type NodeLevel = uint8
type NodeId int64
type String string

func (s String) String() string {
	return string(s)
}
func (id NodeId) String() string {
	return strconv.FormatInt(int64(id), 10)
}
