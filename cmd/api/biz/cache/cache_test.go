package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()
	videos := RC.GetVideosFromCache(time.Now().Format("20060102.150405"))
	for _, item := range videos {
		fmt.Println(item)
	}
}
