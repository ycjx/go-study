package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

//
type K struct {
	V     string
	I     int
	timer *time.Timer //
}

//
//type M struct {
//	kk *K
//}
//
//func myFunction(mm M) {
//	println(&mm)
//	println(mm.kk)
//}
//
//func de() {
//	defer fmt.Println("dddddddd")
//	defer fmt.Println("asdadada")
//	panic("hahha")
//}
//
//const (
//	mutexLocked = 1 << iota // mutex is locked
//	mutexWoken
//	mutexStarving
//	mutexWaiterShift = iota
//)

func main() {
	kk := &K{
		timer: time.AfterFunc(1*time.Second, func() {
			fmt.Println("time out")
		}),
	}
	time.Sleep(10 * time.Second)
	fmt.Println("呵呵")
	fmt.Println(kk.timer.Stop())

}

func closure() func() int {
	x := &K{}
	return func() int {
		x.I++
		return x.I
	}
}

func isStruct(i interface{}) bool {
	return reflect.ValueOf(i).Type().Kind() == reflect.Struct
}

var lastTotalFreed uint64

func printMemStats(s string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v Alloc = %v TotalAlloc = %v JustFreed = %v Sys = %v NumGc = %v\n", s,
		m.Alloc, m.TotalAlloc, (m.TotalAlloc-m.Alloc)-lastTotalFreed, m.Sys, m.NumGC)
	lastTotalFreed = m.TotalAlloc - m.Alloc
}

//
//func fe() {
//	arr := make([]K, 0)
//	arr = append(arr, K{
//		V: "a",
//	})
//	arr = append(arr, K{
//		V: "b",
//	})
//
//	arr2 := make([]K, 0)
//	for _, v := range arr {
//		arr2 = append(arr2, v)
//	}
//	for _, v := range arr2 {
//		fmt.Println(v.V)
//	}
//}
//
//func feMap() {
//	m := make(map[int]int, 0)
//	m[1] = 1
//	m[2] = 2
//	m[3] = 3
//	m[4] = 4
//	m[5] = 5
//	m[6] = 6
//	m[7] = 7
//	m[8] = 8
//	m[9] = 9
//	m[10] = 10
//	for _, v := range m {
//		fmt.Println(v)
//	}
//	fmt.Println("--------")
//	for _, v := range m {
//		fmt.Println(v)
//	}
//	fmt.Println("--------")
//	for _, v := range m {
//		fmt.Println(v)
//	}
//	fmt.Println("--------")
//	for _, v := range m {
//		fmt.Println(v)
//	}
//	fmt.Println("--------")
//}
//
//func fe2() {
//	arr := make([]*K, 0)
//	arr = append(arr, &K{
//		V: "a",
//	})
//	arr = append(arr, &K{
//		V: "b",
//	})
//	bb := new(K)
//	arr = append(arr, bb)
//
//	arr2 := make([]*K, 0)
//	for _, v := range arr {
//		fmt.Println(unsafe.Pointer(&v))
//		arr2 = append(arr2, v)
//	}
//	for _, v := range arr2 {
//		fmt.Println(v.V)
//	}
//
//}
//
//func testRef() {
//	v := reflect.ValueOf(3) // a reflect.Value
//	fmt.Println(v)
//}
//
//func testRef2() {
//	v := M{
//	}
//	value := reflect.ValueOf(&v)
//	el := value.Elem()
//	fmt.Println(value.CanSet())
//	fmt.Println(el.CanSet())
//}
//
//func contextTest() {
//	ctx := context.Background()
//	cc, _ := context.WithTimeout(ctx, time.Second*100)
//	cc.Done()
//}
