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
	msg   string
}

// Error provides more information about the error stacktrace.
func (we wrapedError) Error() string {
	b := bytes.Buffer{}

	// msg for current error: we
	b.WriteString(position(we.fname, we.lno))
	if strings.Trim(we.msg, " ") != "" {
		b.WriteString("[")
		b.WriteString(we.msg)
		b.WriteString("]")
	}

	switch we.err.(type) {
	case wrapedError:
		b.WriteString("\r\n")
	case *wrapedError:
		b.WriteString("\r\n")
	case error:
		b.WriteString(" => ")
	}

	// msg for nested error
	b.WriteString(we.err.Error())

	return string(b.Bytes())
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

// WrapErr will wrap an error with the stack trace information.
func WrapErr(err error) error {
	_, fname, lno, ok := runtime.Caller(1)
	if !ok {
		fname, lno = "unkown", -1
	}

	return newWrapErr(err, fname, lno, "")
}

// WrapErr like WrapErr which support an extra message.
func WrapErrorf(err error, format string, args ...interface{}) error {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, format, args...)
	msg := string(buf.Bytes())

	_, fname, lno, ok := runtime.Caller(1)
	if !ok {
		fname, lno = "unkown", -1
	}

	return newWrapErr(err, fname, lno, msg)
}

func newWrapErr(err error, fname string, lno int, msg string) error {
	switch err.(type) {
	case wrapedError:
		we := err.(wrapedError)
		return &wrapedError{we.depth + 1, err, fname, lno, msg}
	case *wrapedError:
		we := err.(*wrapedError)
		return &wrapedError{we.depth + 1, err, fname, lno, msg}
	default:
		return &wrapedError{0, err, fname, lno, msg}
	}
}

func position(fname string, lno int) string {
	return filepath.Base(fname) + ":" + strconv.Itoa(lno)
}
