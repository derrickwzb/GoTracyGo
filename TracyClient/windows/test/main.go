package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"strconv"
	gotracy "test/gotracy"
	"time"
)


func otherThread() {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()

    gotracy.TracySetThreadName("otherThread")

    for {
        ido := gotracy.TracyZoneBegin("TEST", 0xF0F0FA)
        time.Sleep(1500 * time.Millisecond)
        gotracy.TracyZoneValue(ido, 1000)
        time.Sleep(500 * time.Millisecond)
        gotracy.TracyZoneValue(ido, 500)
        gotracy.TracyZoneEnd(ido)
        gotracy.TracyMessageL("Id from TEST: " + strconv.Itoa(ido))
        gotracy.TracyMessageLC("MESSAGE FROM TEST ZONE", 0xFF0F0F)
        gotracy.TracyFrameMarkName("oThread")
        time.Sleep(3 * time.Second)
    }
}

func main() {
    runtime.LockOSThread()
    defer runtime.UnlockOSThread()

    gotracy.TracySetThreadName("mainThread")

    fmt.Println("TEST")
    i := 0

    go otherThread()

    for {
        id := gotracy.TracyZoneBegin("BLABLA", 0xFF00FF)
        time.Sleep(100 * time.Nanosecond)

        id2 := gotracy.TracyZoneBegin("SUB BLABLA", 0xFF00FF)
        gotracy.TracyZoneValue(id2, 100)
        gotracy.TracyZoneText(id2, "Important information...")

        gotracy.TracyFrameMarkStart("sin(x)")
        time.Sleep(100 * time.Nanosecond)
        gotracy.TracyPlotDouble("sin", math.Sin(float64(i)/10))
        gotracy.TracyFrameMarkEnd("sin(x)")

        gotracy.TracyMessageLC("Important info: "+strconv.Itoa(id), 0xFF3344)
        time.Sleep(100 * time.Nanosecond)

        gotracy.TracyZoneValue(id2, 100)
        gotracy.TracyZoneEnd(id2)
        gotracy.TracyFrameMark()

        i++
        gotracy.TracyZoneEnd(id)
        gotracy.TracyMessageL("Id from BlaBla: " + strconv.Itoa(id))

        if i > 1000 {
            break
        }

        time.Sleep(time.Second)

        gostr := fmt.Sprintf("Loop %d", i)
        log.Println(gostr)
        gotracy.TracyMessageL("TEST" + gostr)
        gotracy.TracyFrameMark()
    }
}