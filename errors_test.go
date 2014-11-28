/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2014/11/28 12:17:38
 */

package goutils

import (
    "testing"
)

func a() error {
    return NewErr("a()")
}

func b() error {
    if err := a(); err != nil {
        return WrapErr(err)
    }

    return nil
}

func c() error {
    return WrapErr(b())
}

func d() error {
    return WrapErr(c())
}

func TestWarpErr(t *testing.T) {
    err := d()
    t.Log(err.(*wrapedError).depth)
    t.Log(err)
}
