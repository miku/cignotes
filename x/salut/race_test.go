package salut

import (
	"testing"
)

func TestSalutation(t *testing.T) {
	Race()
}

// === RUN   TestSalutation
// ==================
// WARNING: DATA RACE
// Write at 0x00c0000a4490 by goroutine 9:
//   github.com/miku/cignotes/x/salut.TestSalutation.func1()
//       /home/tir/code/miku/cignotes/x/salut/race_test.go:15 +0x6c
//
// Previous write at 0x00c0000a4490 by goroutine 8:
//   github.com/miku/cignotes/x/salut.TestSalutation()
//       /home/tir/code/miku/cignotes/x/salut/race_test.go:17 +0x106
//   testing.tRunner()
//       /usr/local/go/src/testing/testing.go:909 +0x199
//
// Goroutine 9 (running) created at:
//   github.com/miku/cignotes/x/salut.TestSalutation()
//       /home/tir/code/miku/cignotes/x/salut/race_test.go:13 +0xf8
//   testing.tRunner()
//       /usr/local/go/src/testing/testing.go:909 +0x199
//
// Goroutine 8 (running) created at:
//   testing.(*T).Run()
//       /usr/local/go/src/testing/testing.go:960 +0x651
//   testing.runTests.func1()
//       /usr/local/go/src/testing/testing.go:1202 +0xa6
//   testing.tRunner()
//       /usr/local/go/src/testing/testing.go:909 +0x199
//   testing.runTests()
//       /usr/local/go/src/testing/testing.go:1200 +0x521
//   testing.(*M).Run()
//       /usr/local/go/src/testing/testing.go:1117 +0x2ff
//   main.main()
//       _testmain.go:44 +0x223
// ==================
// welcome
// --- FAIL: TestSalutation (0.00s)
//     testing.go:853: race detected during execution of test
// FAIL
// FAIL	github.com/miku/cignotes/x/salut	0.008s
// FAIL
