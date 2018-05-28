package log

import (
	"encoding/json"
	"fmt"
	"time"
)

type M map[string]interface{}

func Log(topic Topic, obj interface{}) {
	bytes, _ := json.Marshal(obj)
	str := string(bytes)
	fmt.Println(time.Now().UTC().UnixNano(), "/", topic, "/", str)
}
