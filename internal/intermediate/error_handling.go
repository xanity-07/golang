package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x int) (int, error) {
	if x < 0 {
		return 0, errors.New("math error: square root of negative number")
	}
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty data set")
	}
	return nil
}

func IntroErros() {
	res, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println(err)
		return
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
