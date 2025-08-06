package runtime

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 签名:
//
//	func Gosched()
//
// 用途:
//
//	让出当前的执行权，让其他goroutine执行，将goroutine从running状态切换到runable状态
//
// 用例：
//
//	go 1.14 版本后，异步抢占会监控goroutine,超过10ms的go会尝试抢占，
//	因此runtime.GoSched()这种主动出让执行权的使用场景很少。
//
//	一般情况下，Go 推荐使用 channel、sync 包等同步原语来做并发控制，而不是自旋+`runtime.Gosched()` 主要用于特殊场景或底层库开发。

var ready = false

func Test_Gosched(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Second)
		ready = true
	}()

	for !ready {
		runtime.Gosched() // 主动让出 CPU，避免忙等独占
	}
	fmt.Println("ready!")
}
