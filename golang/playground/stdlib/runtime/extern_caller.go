package runtime

import (
	"fmt"
	"runtime"
)

// func Caller(skip int) (pc uintptr, file string, line int, ok bool)
// 用于运行时获取调用者信息
// skip=0 为当前函数, skip=1 为调用者, skip=2 为调用者的调用者,以此类推
// Log 用于打印调用者信息
func Log(s string) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		fmt.Printf("[%s :%d] %s\n", file, line, s)
	}
}
