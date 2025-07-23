package main

import "runtime/debug"

func main() {
	// Тюнинг памяти и GC
	// SetGCPercent устанавливает процент времени, в течение которого GC будет выполняться.
	debug.SetGCPercent(100)
	// SetMemoryLimit устанавливает максимальный размер памяти в мегабайтах, который может быть выделен для программы.
	debug.SetMemoryLimit(100)
}
