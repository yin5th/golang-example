package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	if longestCommonPrefix([]string{"flower", "flow", "flight"}) != "fl" {
		t.FailNow()
	}

	if longestCommonPrefix([]string{"", "b"}) != "" {
		t.Fail()
	}

	if longestCommonPrefix([]string{"a"}) != "a" {
		t.Fail()
	}
}

//go test -bench .
func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = longestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}

//go test -bench . -benchatime 5s
func BenchmarkSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sleep()
	}
}

//go test -bench . -benchmem -gcflags "-N -l" #禁用内联和优化，以便观察结果
//输出结果包含单词执行堆内存分配总量和次数
func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = heap()
	}
}
