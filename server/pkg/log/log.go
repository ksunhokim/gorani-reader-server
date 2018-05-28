package log

import (
	"encoding/json"
	"fmt"
	"time"
)

var AppName = ""

type M map[string]interface{}

func Log(topic string, obj interface{}) {
	bytes, _ := json.Marshal(obj)
	str := string(bytes)
	fmt.Println(time.Now().UTC().UnixNano(), "/", AppName+"."+topic, "/", str)
}
