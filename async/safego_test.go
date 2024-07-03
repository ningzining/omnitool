package async

import (
	"log"
	"testing"
	"time"
)

func TestSaveGo(t *testing.T) {
	var c int
	SaveGo(func() {
		a := 1
		b := 2
		c = a + b
	})
	time.Sleep(time.Second)
	log.Println(c)
}
