package common

import (
    "gopkg.in/go-playground/assert.v1"
    "testing"
)

func TestINT_Equal(t *testing.T) {
    s1 := STRING(1)
    assert.Equal(t, true, s1.Equal(STRING(1)))
}
