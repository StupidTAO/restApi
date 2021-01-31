package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type demoCloser struct {
	Name string
	activeAt time.Time

}

func (p *demoCloser)Close() error {
	return p.Close()
}

func (p *demoCloser)GetActiveTime() time.Time {
	return p.activeAt
}

func main() {
	Max_Num := os.Getenv("MAX_NUM")
	MaxWorker := runtime.NumCPU()
	fmt.Println(Max_Num, MaxWorker)
	//pool, err := pool.NewGenericPool(0, 10, time.Minute * 10, func() (poolable pool.Poolable, e error) {
	//
	//	time.Sleep(time.Second)
	//	name := strconv.FormatInt(time.Now().Unix(), 10)
	//
	//	fmt.Printf("%s created \n", name)
	//	return &demoCloser{Name:name, activeAt:time.Now()}, nil
	//})
	//
	////acquire closer
	//if err != nil {
	//	fmt.Println("error: ", err)
	//	return
	//}
	//
	//for i := 0; i < 10; i++ {
	//	closer, err := pool.Acquire()
	//	if err != nil {
	//		fmt.Println("pool acquire error: ", err)
	//		continue
	//	}
	//
	//	//do something
	//
	//	closer.GetActiveTime()
	//}
	//
	//pool.Shutdown()
}

