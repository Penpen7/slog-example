package log

import (
	"fmt"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/errbase"
)

type stackTraceElement struct {
	Function string `json:"function,omitempty"`
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
}

func getStackTraceFromError(err error) ([]stackTraceElement, error) {
	st, ok := err.(errbase.StackTraceProvider)
	if !ok {
		return nil, errors.New("error does not implement StackTraceProvider")
	}

	stackTrace := st.StackTrace()
	stackTraceElements := make([]stackTraceElement, 0, len(stackTrace))
	for _, frame := range stackTrace {
		fn := fmt.Sprintf("%n", frame)
		file := fmt.Sprintf("%s", frame)
		line, err := strconv.Atoi(fmt.Sprintf("%d", frame))
		if err != nil {
			return nil, err
		}
		stackTraceElements = append(stackTraceElements, stackTraceElement{
			Function: fn,
			File:     file,
			Line:     line,
		})
	}

	return stackTraceElements, nil
}
