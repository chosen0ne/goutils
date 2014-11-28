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
)

func NewErr(format string, args ...interface{}) error {
    var buf bytes.Buffer
    _, fname, lno, ok := runtime.Caller(1)
    if !ok {
        fname, lno = "unkown", -1
    }

    fmt.Fprintf(&buf, format, args...)

    // errors.go:15 => some error
    return errors.New(fname + ":" + strconv.Itoa(lno) + "=>" + string(buf.Bytes()))
}

func WrapErr(err error) error {
    _, fname, lno, ok := runtime.Caller(1)
    if !ok {
        fname, lno = "unkown", -1
    }

    return errors.New(fname + ":" + strconv.Itoa(lno) + "<=" + err.Error())
}
