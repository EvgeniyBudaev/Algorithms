package main

import (
	"net/http"
	"net/http/pprof"
	"runtime"
	"runtime/debug"
)

func main() {
	// SetGCPercent устанавливает процент времени, в течение которого GC будет выполняться.
	debug.SetGCPercent(100)
	// SetMemoryLimit устанавливает максимальный размер памяти в мегабайтах, который может быть выделен для программы.
	debug.SetMemoryLimit(100)

	// SetMutexProfileFraction устанавливает частоту сбора профиля мьютексов.
	runtime.SetMutexProfileFraction(1)
	// SetBlockProfileRate устанавливает частоту сбора профиля блокировок.
	runtime.SetBlockProfileRate(1)
	// SetCPUProfileRate устанавливает частоту сбора профиля CPU.
	runtime.SetCPUProfileRate(1)
}

// registerProfilerEndpoints регистрирует эндпоинты для профилировки.
func registerProfilerEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("goroutine"))
}
