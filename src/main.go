package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Msg string
	Err error
}

func (e *CustomError) Error() string {
	return e.Msg
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func main() {
	level1Err := errors.New("[Error in L1]: Boom")
	level2Err := fmt.Errorf("[Error in L2]: Wrap L1Err %w", level1Err)
	level3Err := fmt.Errorf("[Error in L3]: Wrap L2Err %w", level2Err)
	//level3Err := &CustomError{"[Error in L3]: Wrap L2Err", level2Err}
	level4Err := fmt.Errorf("[Error in L4]: Wrap L3Err %w", level3Err)
	fmt.Println(level4Err)
}

// Console output, when uncomment line28 and comment line29:
// [Error in L4]: Wrap L3Err [Error in L3]: Wrap L2Err [Error in L2]: Wrap L1Err [Error in L1]: Boom

// Console output, when uncomment line29 and comment line28:
// [Error in L4]: Wrap L3Err [Error in L3]: Wrap L2Err
