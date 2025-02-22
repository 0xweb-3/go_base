package pkg

import (
	"testing"
	"time"
)

//  go test -bench . -benchtime=2s  -benchmem 性能分析执行的时间

func BenchmarkF1(b *testing.B) {
	b.N = 1000 // 分析执行次数
	time.Sleep(time.Second)

	b.StartTimer() // 性能分析开始的时间位置
	for i := 0; i < b.N; i++ {
		F1()
	}
}

func BenchmarkF2(b *testing.B) {
	b.StopTimer() // 停止记时
	time.Sleep(time.Second)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		F2()
	}
}
