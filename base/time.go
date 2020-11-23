package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("yyyy-MM-dd HH:mm:ss"))
}
