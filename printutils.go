package printutils

import (
	"fmt"
	"io"
)

func FPrintDouble(w io.Writer, m string) (n int, err error) {
	n, err = fmt.Fprintln(w, m)
	if err != nil {
		return
	}

	n2, err := fmt.Fprintln(w, m)
	if err != nil {
		return n + n2, err
	}

	return n + n2, nil
}
