package gotracy

/*
#cgo CXXFLAGS: -I${SRCDIR} -DTRACY_ENABLE -std=c++17
#cgo LDFLAGS: -L${SRCDIR} -l:libTracyClient.a -lstdc++ -lpthread
#include "TracyGoWrapper.h"
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

import (
	"runtime"
	"sync"
)

var tracyStringsMap map[string]*C.char = make(map[string]*C.char)
var allocStringMutex sync.Mutex

func allocString(text string) *C.char {

	// allocStringMutex.Lock()
	// defer allocStringMutex.Unlock()

	// val, ok := tracyStringsMap[text]
	// if ok {
	// 	return val
	// }

	// cgotext := C.CString(text)
	// tracyStringsMap[text] = cgotext
	// return cgotext


	cgotext := C.CString(text)
	return cgotext
}

func TracySetThreadName(name string) {
	runtime.LockOSThread() //so that this goroutine will only continue running on this thread

	C.GoTracySetThreadName(allocString(name))
}

func TracySetGoroutineName(name string){
	C.GoTracySetThreadName(allocString(name))
}

var tracyZoneBeginMutex sync.Mutex

func TracyZoneBegin(name string, color uint32) int {

	// tracyZoneBeginMutex.Lock()
	// defer tracyZoneBeginMutex.Unlock()

	// pc, filename, line, _ := runtime.Caller(1)
	// funcname := runtime.FuncForPC(pc).Name()

	// ret := C.GoTracyZoneBegin(allocString(name), allocString(funcname),
	// 	allocString(filename), C.uint(line), C.uint(color))

	// return int(ret)

	pc, filename, line, _ := runtime.Caller(1)
	funcname := runtime.FuncForPC(pc).Name()

	ret := C.GoTracyZoneBegin(allocString(name), allocString(funcname),
		allocString(filename), C.uint(line), C.uint(color))

	return int(ret)
}

func TracyZoneEnd(c int) {
	C.GoTracyZoneEnd(C.int(c))
}

func TracyZoneValue(c int, value int64) {
	C.GoTracyZoneValue(C.int(c), C.uint64_t(value))
}

func TracyZoneText(c int, text string) {
	C.GoTracyZoneText(C.int(c), allocString(text))
}

func TracyMessageL(msg string) {
	C.GoTracyMessageL(allocString(msg))
}

func TracyMessageLC(msg string, color uint32) {
	C.GoTracyMessageLC(allocString(msg), C.uint(color))
}

func TracyFrameMark() {
	C.GoTracyFrameMark()
}

func TracyFrameMarkName(name string) {
	C.GoTracyFrameMarkName(allocString(name))
}

func TracyFrameMarkStart(name string) {
	C.GoTracyFrameMarkStart(allocString(name))
}

func TracyFrameMarkEnd(name string) {
	C.GoTracyFrameMarkEnd(allocString(name))
}

func TracyPlotFloat(name string, val float32) {
	C.GoTracyPlotFloat(allocString(name), C.float(val))
}

func TracyPlotDouble(name string, val float64) {
	C.GoTracyPlotDoublet(allocString(name), C.double(val))
}

func TracyPlotInt(name string, val int) {
	C.GoTracyPlotInt(allocString(name), C.int(val))
}

func TracyMessageAppinfo(name string) {
	C.GoTracyMessageAppinfo(allocString(name))
}

