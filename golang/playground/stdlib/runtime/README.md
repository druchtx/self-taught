# Go `runtime` Package: Q&A Guide

This guide answers common questions about Go's `runtime` package, categorized from essential basics to expert-level tools.

## Tier 1: Core Essentials (Debugging & Concurrency Basics)

### Q: What is `runtime.Gosched()` and when should I use it?

**A:** `runtime.Gosched()` yields the current Goroutine's execution time, allowing the scheduler to run other waiting Goroutines.

*   **Use Case:** In older Go versions, it was used in long-running, CPU-bound tasks to prevent them from monopolizing a CPU.
*   **Modern Relevance:** Since Go 1.14+, the preemptive scheduler makes manual calls largely unnecessary, but it's still a vital concept to understand for concurrency.

### Q: How does `runtime.GOMAXPROCS` work and do I need to set it?

**A:** `runtime.GOMAXPROCS(n)` sets the maximum number of OS threads that can execute Go code simultaneously.

*   **Modern Usage:** You almost never need to set it. Since Go 1.5, it defaults to the number of available CPU cores.
*   **Current Use Case:** Its main use today is to query the current value with `runtime.GOMAXPROCS(0)`.

### Q: How can I get a Goroutine's stack trace for debugging?

**A:** Use `runtime.Stack(buf []byte, all bool)`. It captures the stack trace of the current Goroutine (or all Goroutines if `all` is true) and stores it in the `buf`.

*   **Use Case:** This is a critical debugging tool. After a `recover()`, call `runtime.Stack` to log the full trace that caused the panic, providing much more context.

---

## Tier 2: Advanced & Practical (Profiling & Introspection)

### Q: How can I monitor my Go application's memory usage?

**A:** Use `runtime.ReadMemStats(m *MemStats)`. This function populates a `MemStats` struct with detailed memory allocation and GC statistics.

*   **Use Case:** Periodically call this to export metrics (like `Alloc`, `HeapObjects`, `NumGC`) to a monitoring system (e.g., Prometheus) to track service health or diagnose memory leaks.

### Q: How do I get the caller's file and line number for logging?

**A:** Use `runtime.Caller(skip int)` for a single caller or `runtime.Callers(skip int, pc []uintptr)` for the full stack.

*   **Use Case:** Essential for logging frameworks. `runtime.Caller(1)` (or `2`, depending on call depth) can automatically capture where a log message was generated. The `runtime.Callers` + `runtime.CallersFrames` pattern is the best practice for capturing a full stack trace for error reporting.

### Q: How can I write code that behaves differently on Windows vs. Linux?

**A:** Use the compile-time constants `runtime.GOOS` and `runtime.GOARCH`.

*   **Use Case:** They allow you to write platform-specific code.
    ```go
    if runtime.GOOS == "windows" {
        // Windows-specific logic
    }
    ```

---

## Tier 3: Expert & Special-Case

### Q: When should I use `runtime.SetFinalizer`?

**A:** Its main legitimate use is for releasing **non-Go resources** managed via CGO. When a Go object wrapping a C resource is about to be garbage collected, the finalizer can be run to call the necessary C `free()` function.

*   **Warning:** Use with extreme caution. It is not a replacement for `defer`. Finalizers are not guaranteed to run and can cause complex issues.

### Q: What are `runtime.LockOSThread` and `runtime.UnlockOSThread` for?

**A:** They bind a Goroutine to a specific OS thread.

*   **Use Case:** This is required only when interacting with external C libraries (e.g., some GUI or OpenGL libraries) that demand all calls originate from the same thread. It breaks Go's normal scheduling model and should be used sparingly.

### Q: Is it a good idea to call `runtime.GC()` manually?

**A:** Almost **never** in production code. Go's GC is highly optimized, and manual calls usually hurt performance.

*   **Use Case:** It is a tool for **testing and benchmarking only**, used to create a consistent memory state between test runs.

---

## Additional Essential Tools

### Q: What is the standard way to profile a Go application?

**A:** The `runtime/pprof` package is the standard tool for CPU and memory profiling. You can use it to start a web server for live profiling or to write profile files for offline analysis with `go tool pprof`.

### Q: How do I find data races in my concurrent Go code?

**A:** Use the built-in race detector by compiling and running your code with the `-race` flag (e.g., `go test -race` or `go run -race`). It is an indispensable tool for writing correct concurrent programs.

### Q: What is `runtime.KeepAlive` for?

**A:** It's a special function that tells the garbage collector that an object is still in use, even if it's not apparent from the code. This is typically only needed when dealing with `unsafe` pointers or CGO, to prevent the GC from prematurely freeing memory that an external C function is still using.
