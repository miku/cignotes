package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()

		v1.mu.Lock()
		defer v1.mu.Lock()

		time.Sleep(2 * time.Second) // Q: Would it deadlock without the sleep as well?
		v2.mu.Lock()
		defer v2.mu.Lock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)

	go printSum(&a, &b)
	go printSum(&b, &a) // Q: Would it deadlock with printSum(&a, &b) as well?

	wg.Wait()
}

// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc00001c178)
//         /usr/local/go/src/runtime/sema.go:56 +0x42
// sync.(*WaitGroup).Wait(0xc00001c170)
//         /usr/local/go/src/sync/waitgroup.go:130 +0x64
// main.main()
//         /home/tir/code/miku/cignotes/x/deadlock/main.go:36 +0x122
//
// goroutine 6 [semacquire]:
// sync.runtime_SemacquireMutex(0xc00001c194, 0x1300, 0x1)
//         /usr/local/go/src/runtime/sema.go:71 +0x47
// sync.(*Mutex).lockSlow(0xc00001c190)
//         /usr/local/go/src/sync/mutex.go:138 +0xfc
// sync.(*Mutex).Lock(...)
//         /usr/local/go/src/sync/mutex.go:81
// main.main.func1(0xc00001c180, 0xc00001c190)
//         /home/tir/code/miku/cignotes/x/deadlock/main.go:24 +0x1f4
// created by main.main
//         /home/tir/code/miku/cignotes/x/deadlock/main.go:33 +0xea
//
// goroutine 7 [semacquire]:
// sync.runtime_SemacquireMutex(0xc00001c184, 0x1300, 0x1)
//         /usr/local/go/src/runtime/sema.go:71 +0x47
// sync.(*Mutex).lockSlow(0xc00001c180)
//         /usr/local/go/src/sync/mutex.go:138 +0xfc
// sync.(*Mutex).Lock(...)
//         /usr/local/go/src/sync/mutex.go:81
// main.main.func1(0xc00001c190, 0xc00001c180)
//         /home/tir/code/miku/cignotes/x/deadlock/main.go:24 +0x1f4
// created by main.main
//         /home/tir/code/miku/cignotes/x/deadlock/main.go:34 +0x114
// exit status 2
