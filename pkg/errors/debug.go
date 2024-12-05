package errs

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

var separator = ":\n\t"

type stack struct {
	callers []uintptr
}

func (e *HTTPError) populateStack() {
	e.callers = callers()

	e2, ok := e.err.(*HTTPError)
	if !ok {
		return
	}

	i := 0
	ok = false
	for ; i < len(e.callers) && i < len(e2.callers); i++ {
		if e.callers[len(e.callers)-1-i] != e2.callers[len(e2.callers)-1-i] {
			break
		}
		ok = true
	}
	if ok {
		head := e2.callers[:len(e2.callers)-i]
		tail := e.callers
		e.callers = make([]uintptr, len(head)+len(tail))
		copy(e.callers, head)
		copy(e.callers[len(head):], tail)
		e2.callers = nil
	}
}

func callers() []uintptr {
	var stk [64]uintptr
	const skip = 4
	n := runtime.Callers(skip, stk[:])
	return stk[:n]
}

func (e *HTTPError) StackTrace() string {
	printCallers := callers()
	b := new(bytes.Buffer)

	var prev string
	var diff bool
	for i := 0; i < len(e.callers); i++ {
		thisFrame := frame(e.callers, i)
		name := thisFrame.Func.Name()

		if !diff && i < len(printCallers) {
			if name == frame(printCallers, i).Func.Name() {
				continue
			}
			diff = true
		}

		if name == prev {
			continue
		}

		trim := 0
		for {
			j := strings.IndexAny(name[trim:], "./")
			if j < 0 {
				break
			}
			if !strings.HasPrefix(prev, name[:j+trim]) {
				break
			}
			trim += j + 1 // skip over the separator
		}

		pad(b, separator)
		fmt.Fprintf(b, "%v:%d: ", thisFrame.File, thisFrame.Line)
		if trim > 0 {
			b.WriteString("...")
		}
		b.WriteString(name[trim:])

		prev = name
	}

	return b.String()
}

func frame(callers []uintptr, n int) *runtime.Frame {
	frames := runtime.CallersFrames(callers)
	var f runtime.Frame
	for i := len(callers) - 1; i >= n; i-- {
		var ok bool
		f, ok = frames.Next()
		if !ok {
			break
		}
	}
	return &f
}
