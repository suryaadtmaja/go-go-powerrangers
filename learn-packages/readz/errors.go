package readz

import (
	"errors"
	"fmt"
)

func ReturnError(errorMessages string) {
	err := errors.New(errorMessages)
	fmt.Println(err)
}
