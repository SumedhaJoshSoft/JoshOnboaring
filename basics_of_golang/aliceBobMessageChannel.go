package main

import (
	"strings"
)

func main() {
	str := "helloBob$helloALice#howareyou#fine,how are you?^"

	var aliceCount, bobCount = strings.Count(str, "$"), strings.Count(str, "#")

}
