package runtime

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"testing"
)

// 签名:
//
//	func Caller(skip int) (pc uintptr, file string, line int, ok bool)
//
// 用途:
//
//	用于运行时获取调用者信息，skip=0 为当前函数, skip=1 为调用者, skip=2 为调用者的调用者,以此类推
//
// 用例：
//
//	Log 记录当前调用函数的信息
func Log(s string) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		fmt.Printf("[%s :%d] %s\n", file, line, s)
	}
}

func TestCallerLog(t *testing.T) {

	// 回避标准输出
	old := os.Stdout
	// 创建管道，写入w的数据，可以从r中读取
	r, w, _ := os.Pipe()
	os.Stdout = w

	Log("Hello, World!")

	_ = w.Close()
	os.Stdout = old // restoring the real stdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	got := buf.String()

	// The expected output is like "[/path/to/file.go:123] Hello, World!\n"
	// We can't know the exact file path and line number, so we use a regex.
	// This regex checks for:
	// - a prefix `[`
	// - a file path (anything but `:`)
	// - a colon and a line number
	// - a closing bracket `]`
	// - a space
	// - the message "Hello, World!"
	// - a newline
	re := regexp.MustCompile(`^\[.+:\d+\] Hello, World!\n$`)
	t.Logf("got:%q", got)
	if !re.MatchString(got) {
		t.Errorf("Log output format is incorrect. got: %q", got)
	}
}
