package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

// 生成 trace.out
// 使用 go tool trace trace.out 打开一个web界面
func main() {
	var buf bytes.Buffer

	defer func() {
		filename := "trace.out"
		if err := os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
			return
		}
	}()

	// 开始追踪
	if err := trace.Start(&buf); err != nil {
		return
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	// 定义task
	ctx, task := trace.NewTask(context.Background(), "new task")
	for i := range 1000 {
		wg.Go(func() {
			// 定义region
			r := trace.StartRegion(ctx, fmt.Sprintf("region %d", i))
			// 定义log
			// trace.Logf(ctx, "log", "log in region %d", i)
			defer r.End()
			_ = make([]byte, 1024*1024)
		})
	}
	wg.Wait()
	task.End()
}
