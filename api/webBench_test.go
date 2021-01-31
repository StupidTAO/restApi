package api

import "testing"

func BenchmarkARankTop(b *testing.B) {
	b.StopTimer()
	//do something
	tokenName := "AOA"
	b.StartTimer()
	//number := 10
	for i := 0; i < b.N; i++ {
		_, err := ARankTop(tokenName)
		if err != nil {
			b.Error("BenchmarkARankTop occur error: ", err)
		}
	}
}