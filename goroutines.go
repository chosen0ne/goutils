/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-11-16 15:29:43
 */

package goutils

import (
	"time"
)

type Runnable func()

type ScheduledRunnable func(time.Time)

// GoroutineLoop is a safe loop function for goroutine, which can't be terminated by panic
func GoroutineLoop(runnable Runnable, log Logger) {
	defer recoverFn("GoroutineLoop", log)

	for {
		do(runnable, log)
	}
}

// GoroutineScheduled is a safe scheduled loop function for goroutine,
// which can't be terminated by panic
func GoroutineScheduled(interval time.Duration, runnable ScheduledRunnable, log Logger) {
	defer recoverFn("GoroutineScheduled", log)

	for t := range time.Tick(interval) {
		doScheduled(t, runnable, log)
	}
}

func do(runnable Runnable, log Logger) {
	defer recoverFn("do", log)

	runnable()
}

func doScheduled(t time.Time, runnable ScheduledRunnable, log Logger) {
	defer recoverFn("doScheduled", log)

	runnable(t)
}

func recoverFn(funcName string, log Logger) {
	if err := recover(); err != nil {
		switch err.(type) {
		case error:
			log.Exception(err.(error), "panic recovered in %s", funcName)
		default:
			log.Error("unknown error recovered in %s, err: %v", funcName, err)
		}
	}
}
