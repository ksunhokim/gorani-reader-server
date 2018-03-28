package main

import (
	"fmt"
	"strconv"
)

func main() {
	maxUint := ^uint(0)
	maxInt := int(maxUint >> 1)
	maxi := strconv.FormatInt(int64(maxInt), 10)
	fmt.Println(maxi)
}
