package base

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ti := time.Now()
	fmt.Println(ti.Format("yyyy-MM-dd HH:mm:ss"))
}
