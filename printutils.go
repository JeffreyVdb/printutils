package printutils

import (
	"io"

	"github.com/JeffreyVdb/upperprinter/pkg/printerz"
)

func FPrintDouble(w io.Writer, m string) (n int, err error) {
	n, err = printerz.FPrintUpper(w, m)
	if err != nil {
		return
	}

	n2, err := printerz.FPrintUpper(w, m)
	if err != nil {
		return n + n2, err
	}

	return n + n2, nil
}
