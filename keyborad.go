// Package keyborad: get user keybord input util
package keyborad

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// get float64 number
func GetFloatInPut() (input float64, err error) {
	fmt.Print("please enter you number: ")
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	space := strings.TrimSpace(readString)
	input, err = strconv.ParseFloat(space, 64)
	if err != nil {
		return 0, err
	}
	return input, nil
}

// get int number
func GetIntInPut() (input int, err error) {
	fmt.Print("please enter you number: ")
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	space := strings.TrimSpace(readString)
	input, err = strconv.Atoi(space)
	if err != nil {
		return 0, err
	}
	return input, nil
}
