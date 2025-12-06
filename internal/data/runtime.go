package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeType = errors.New("Invalid runtime type for JOSN")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {

	runtimeFmt := fmt.Sprintf("%d mins", r)

	// put the value in double qoute in order to make it a valid JSON value
	quotedRuntime := strconv.Quote(runtimeFmt)

	return []byte(quotedRuntime), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {

	unqoutedRuntime, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return err
	}

	parts := strings.Split(unqoutedRuntime, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeType
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeType
	}

	*r = Runtime(i)

	return nil
}
