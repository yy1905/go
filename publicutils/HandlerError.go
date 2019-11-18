package publicutils

import (
	"fmt"
	"os"
)

func HandlerError(fun string, err error) {
	if err != nil {
		fmt.Println(fun, err)
		os.Exit(1)
	}
}