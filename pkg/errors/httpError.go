package errs

import (
	"bytes"
	"fmt"

	"github.com/goccy/go-json"
)

type Code int
type Uuid string
type Message string
type Op string
type Params interface{}

type HTTPError struct {
	op        Op
	kind      kind
	echoToken Uuid
	message   Message
	err       error
	params    Params
	stack
}

func (e *HTTPError) SetParams(params any) {
	e.params = Params(params)
}

func (e *HTTPError) SetEchoToken(echoToken string) {
	e.echoToken = Uuid(echoToken)
}

func (e *HTTPError) EchoToken() string {
	return string(e.echoToken)
}

func (e *HTTPError) JSONParams() string {
	if e.params == nil {
		return ""
	}

	jsonParams, err := json.Marshal(e.params)
	if err != nil {
		return ""
	}

	return string(jsonParams)
}

func (e *HTTPError) Error() string {
	b := new(bytes.Buffer)
	if e.op != "" {
		pad(b, ": ")
		b.WriteString(string(e.op))
	}
	if e.kind != 0 {
		pad(b, ": ")
		b.WriteString(e.kind.String())
	}
	if e.message != "" {
		pad(b, ": ")
		b.WriteString(string(e.message))
	}
	if e.err != nil {
		// Indent on new line if we are cascading non-empty errors.
		if prevErr, ok := e.err.(*HTTPError); ok {
			if !prevErr.isZero() {
				pad(b, ": ")
				b.WriteString(e.err.Error())
			}
		} else {
			pad(b, ": ")
			b.WriteString(e.err.Error())
		}
	}
	if b.Len() == 0 {
		return "no error"
	}
	return b.String()
}

func (e *HTTPError) isZero() bool {
	return e.kind == 0 && e.message == "" && e.err == nil
}

func (he *HTTPError) Status() int {
	if he.kind != 0 {
		return he.kind.HttpStatus()
	}

	copy := *he
	for copy.err != nil {
		if err, ok := copy.err.(*HTTPError); ok {
			if err.kind != 0 {
				return err.kind.HttpStatus()
			} else {
				copy = *err
			}
		} else {
			break
		}
	}

	return he.kind.HttpStatus()
}

func (he *HTTPError) Message() string {
	b := new(bytes.Buffer)
	b.WriteString(string(he.message))

	copy := *he
	for copy.err != nil {
		if err, ok := copy.err.(*HTTPError); ok {
			if err.message != "" {
				pad(b, ": ")
				b.WriteString(string(err.message))
			}
			copy = *err
		} else {
			break
		}

	}
	return b.String()
}

func New(args ...interface{}) error {
	if len(args) == 0 {
		panic("call to errors.E with no arguments")
	}
	e := &HTTPError{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Uuid:
			e.echoToken = arg
		case Op:
			e.op = arg
		case string:
			e.message = Message(arg)
		case kind:
			e.kind = arg
		case *HTTPError:
			// Make a copy
			copy := *arg
			e.err = &copy
		case error:
			e.err = arg
		case Params:
			e.params = arg
		default:
			panic(fmt.Sprintf("unknown type %T, value %v in error call", arg, arg))
		}
	}

	e.populateStack()

	prev, ok := e.err.(*HTTPError)
	if !ok {
		return e
	}

	if prev.message == e.message {
		prev.message = ""
	}
	if prev.kind == e.kind {
		prev.kind = 0
	}
	return e
}

func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}
