/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2014/11/28 11:51:28
 */

package goutils

import (
    "runtime"
    "errors"
    "fmt"
    "bytes"
    "strconv"
    "strings"
    "path/filepath"
)

type wrapedError struct {
    depth   int
    err     error
    fname   string
    lno     int
}

func (we wrapedError) Error() string {
    return we.err.Error() + "\n" + strings.Repeat("  ", we.depth) + "=> " + position(we.fname, we.lno)
}

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
