package main

import "testing"

// goos: linux
// goarch: amd64
// pkg: comparasion/cmd
// cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
// Benchmark_GinPointer-8   	       1	10000060135 ns/op	730231824 B/op	 6259022 allocs/op
// 2Benchmark_GinCallback-8   	       1	10000066233 ns/op	765509976 B/op	 6573168 allocs/op
// Benchmark_EchoPointer-8   	       1	10000155163 ns/op	776537592 B/op	 6616489 allocs/op
// Benchmark_EchoCallback-8   	       1	10000090828 ns/op	750889792 B/op	 6388141 allocs/op

func Benchmark_GinPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainGinPointer()
	}
}

func Benchmark_GinCallback(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainGinCallback()
	}
}

func Benchmark_EchoPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainEchoPointer()
	}
}

func Benchmark_EchoCallback(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainEchoCallback()
	}
}
