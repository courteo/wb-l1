package tasks

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	select {
	case  <-time.After(d * time.Second):				// запускаем селект из которого можно выйти только спустя d времени
		return
	}
}

func T25() {
	fmt.Println(time.Now())
	mySleep(5)
	fmt.Println(time.Now())
}