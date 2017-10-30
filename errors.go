/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2014/11/28 11:51:28
 */

package goutils

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// wrapedError is uesed to wrap an error, and show the stacktrace
// information.
type wrapedError struct {
	depth int // nested depth
	err   error
	fname string // file name the error occurs
	lno   int    // line number the error occurs
}

// Error provides more information about the error stacktrace.
func (we wrapedError) Error() (e string) {
	switch we.err.(type) {
	case wrapedError:
		e = we.err.Error() + "\n" + strings.Repeat("  ", we.depth) + "=> " + position(we.fname, we.lno)
	default:
		e = position(we.fname, we.lno) + " => " + we.err.Error()
	}
	return
}

// NewErr will create an error which contains the postion information
// of the error.
func NewErr(format string, args ...interface{}) error {
	var buf bytes.Buffer
	_, fname, lno, ok := runtime.Caller(1)
	if !ok {
		fname, lno = "unkown", -1
	}

	fmt.Fprintf(&buf, format, args...)

	// errors.go:15 => some error
	return errors.New(position(fname, lno) + " => " + string(buf.Bytes()))
}

func WrapErr(err error) error {
	_, fname, lno, ok := runtime.Caller(1)
	if !ok {
		fname, lno = "unkown", -1
	}

	switch err.(type) {
	case wrapedError:
		we := err.(wrapedError)
		return &wrapedError{we.depth + 1, err, fname, lno}
	case *wrapedError:
		we := err.(*wrapedError)
		return &wrapedError{we.depth + 1, err, fname, lno}
	default:
		return &wrapedError{0, err, fname, lno}
	}
}

func position(fname string, lno int) string {
	return filepath.Base(fname) + ":" + strconv.Itoa(lno)
}
