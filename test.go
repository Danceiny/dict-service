package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func main() {
	s := "{\"bid\": null}"
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(s), &obj)

	log.Infof("err: %v bid: %v", err, obj["bid"])

	type S struct {
		bid string
	}
	var obj2 S
	err2 := json.Unmarshal([]byte(s), &obj2)

	log.Infof("err: %v bid: %v", err2, obj2.bid == "")

}
