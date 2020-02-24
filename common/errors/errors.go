package errors

import (
	"os"
	"reflect"
	"strings"

	"v2ray.com/core/common/log"
	"v2ray.com/core/common/serial"
)

type hasInnerError interface {
	// Inner returns the underlying error of this one.
	Inner() error
}

type hasSeverity interface {
	Severity() log.Severity
}

type Error struct {
	pathObj  interface{}
	prefix   []interface{}
	message  []interface{}
	inner    error
	severity log.Severity
}

func (err *Error) WithPathObj(obj interface{}) *Error {
	err.pathObj = obj
	return err
}

func (err *Error) pkgPath() string {
	if err.pathObj == nil {
		return ""
	}
	return reflect.TypeOf(err.pathObj).PkgPath()
}

// Error implements error.Error)().
func (err *Error) Error() string {
	builder := strings.Builder{}
	for _, prefix := range err.prefix {
		builder.WriteByte('[')
		builder.WriteString(serial.ToString(prefix))
		builder.WriteString("] ")
	}

	path := err.pkgPath()
	if len(path) > 0 {
		builder.WriteString(path)
		builder.WriteString(": ")
	}

	msg := serial.Concat(err.message...)
	builder.WriteString(msg)

	if err.inner != nil {
		builder.WriteString(" > ")
		builder.WriteString(err.inner.Error())
	}

	return builder.String()
}

func (err *Error) Base(e error) *Error {
	err.inner = e
	return err
}

func (err *Error) atSeverity(s log.Severity) *Error {
	err.severity = s
	return err
}

// AtWarning sets the severity to warning.
func (err *Error) AtWarning() *Error {
	return err.atSeverity(log.Severity_Warning)
}

func (err *Error) WriteToLog(opts ...ExportOption) {
	var holder ExportOptionHolder

	for _, opt := range opts {
		opt(&holder)
	}

	if holder.SessionID > 0 {
		err.prefix = append(err.prefix, holder.SessionID)
	}

	log.Record(&log.GeneralMessage{
		Severity: GetSeverity(err),
		Content:  err,
	})
}

type ExportOptionHolder struct {
	SessionID uint32
}

type ExportOption func(*ExportOptionHolder)

// New returns a new error object with message formed from given arguments.
func New(msg ...interface{}) *Error {
	return &Error{
		message:  msg,
		severity: log.Severity_Info,
	}
}

// Cause returns the root cause of this error.
func Cause(err error) error {
	if err == nil {
		return nil
	}
L:
	for {
		switch inner := err.(type) {
		case hasInnerError:
			if inner.Inner() == nil {
				break L
			}
			err = inner.Inner()
		case *os.PathError:
			if inner.Err == nil {
				break L
			}
			err = inner.Err
		case *os.SyscallError:
			if inner.Err == nil {
				break L
			}
			err = inner.Err
		default:
			break L
		}
	}
	return err
}

// GetSeverity returns the actual severity of the error, including inner errors.
func GetSeverity(err error) log.Severity {
	if s, ok := err.(hasSeverity); ok {
		return s.Severity()
	}
	return log.Severity_Info
}
