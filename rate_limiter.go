package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Rate struct {
	m              sync.Mutex
	lastRefillTime int64
	averageRate    int64 //多少毫秒生成一个token
	Consumed       int64 //已被消耗的token
	BucketSize     int64 //最多消耗的token数量
}

var ConsumeCount int64

//refillToken 往桶里放入token
func (r *Rate) refillToken() {
	////上次获取令牌的时间
	lastT := r.lastRefillTime
	nowT := time.Now().UnixNano() / 1e6
	timeDelta := nowT - lastT
	newToken := timeDelta / r.averageRate
	if newToken > 0 {
		var newRefillTime int64
		if lastT == 0 {
			newRefillTime = nowT
		} else {
			newRefillTime = lastT + newToken*r.averageRate
		}
		if atomic.CompareAndSwapInt64(&r.lastRefillTime, lastT, newRefillTime) {
			if r.Consumed == 0 {
				return
			}
			for {
				var newConsumed int64
				if r.Consumed <= newToken {
					newConsumed = 0
				} else {
					newConsumed = r.Consumed - newToken
				}
				if atomic.CompareAndSwapInt64(&r.Consumed, r.Consumed, newConsumed) {
					return
				}
			}

		}
	}

}

func (r *Rate) Consume() bool {
	if r.Consumed >= r.BucketSize {
		return false
	}
	v := atomic.CompareAndSwapInt64(&r.Consumed, r.Consumed, r.Consumed+1)
	fmt.Println(fmt.Sprintf("CompareAndSwapInt64 : %v", r.Consumed))
	return v
}

func (r *Rate) Acquire() bool {
	r.refillToken()
	return r.Consume()
}

func main() {
	rate := Rate{
		averageRate:    500,
		BucketSize:     20,
		lastRefillTime: time.Now().UnixNano() / 1e6,
	}
	println()
	for i := 0; i < 100; i++ {

		go func(k int, r *Rate) {
			ran := rand.Int63n(200) + 300
			for {
				time.Sleep(time.Duration(ran) * time.Millisecond)
				if rate.Acquire() {
					ConsumeCount++
					fmt.Println(fmt.Sprintf("总共获取锁的次数：%v", ConsumeCount))
					fmt.Println(fmt.Sprintf("成功获取到锁：%v", k))
				}
			}

		}(i, &rate)
	}

	time.Sleep(1 * time.Hour)

}
