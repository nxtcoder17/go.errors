package errors

import (
	"errors"
	"fmt"
	"strings"
)

type Error struct {
	err error
	kv  []any
}

var _ error = (*Error)(nil)

func (e *Error) KV(kvPairs ...any) *Error {
	if e == nil {
		return e
	}

	if e.kv == nil {
		e.kv = make([]any, 0, len(kvPairs))
	}

	for i := 1; i < len(kvPairs); i += 2 {
		if ks, ok := kvPairs[i-1].(string); ok {
			e.kv = append(e.kv, ks, kvPairs[i])
		}
	}
	return e
}

// Error implements error.
func (e *Error) Error() string {
	var msg strings.Builder
	msg.WriteString(e.err.Error())

	if len(e.kv) > 0 {
		msg.WriteString(" <")
	}
	for i := 1; i < len(e.kv); i += 2 {
		msg.WriteString(fmt.Sprintf("%v=%v", e.kv[i-1], e.kv[i]))
		if i < len(e.kv)-2 {
			msg.WriteString(", ")
		}
	}

	if len(e.kv) > 0 {
		msg.WriteString(">")
	}
	return msg.String()
}

func (e *Error) GetErrMessage() string {
	return e.err.Error()
}

func (e *Error) GetKV() []any {
	return e.kv
}

func (e *Error) AsKeyValues() []any {
	kv := make([]any, 0, len(e.kv)+2)
	kv = append(kv, "err", e.err.Error())
	return append(kv, e.kv...)
}

func (e *Error) Wrap(err error) *Error {
	if v, ok := err.(*Error); ok {
		e.kv = append(e.kv, v.kv...)
		e.err = fmt.Errorf("%w |> %s", e.err, v.err.Error())
		return e
	}

	e.err = fmt.Errorf("%w |> %w", e.err, err)
	return e
}

// New creates a new error with the given message
func New(msg string) *Error {
	return &Error{err: errors.New(msg)}
}

// Below functions provide compatibility with std go errors
var (
	Is   = errors.Is
	As   = errors.As
	Join = errors.Join
)
