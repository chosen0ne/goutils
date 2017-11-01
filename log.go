/**
 *
 * @author  chosen0ne(louzhenlin86@126.com)
 * @date    2017-10-30 15:54:13
 */

package goutils

// Logger interface which is depended by package 'goutils'
type Logger interface {
	Debug(fmt string, vals ...interface{})
	Info(fmt string, vals ...interface{})
	Warn(fmt string, vals ...interface{})
	Error(fmt string, vals ...interface{})
	Exception(e error, fmt string, vals ...interface{})
	Fatal(fmt string, vals ...interface{})
}
