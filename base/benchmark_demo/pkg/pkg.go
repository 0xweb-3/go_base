package pkg

import "time"

func F1() {
	time.Sleep(time.Millisecond * 10)
}

func F2() {
	time.Sleep(time.Millisecond * 20)
}
