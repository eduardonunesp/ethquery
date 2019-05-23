package cmd

import (
	"fmt"
	"strconv"
)

func hexify(numericalString string) string {
	if len(numericalString) >= 3 && numericalString[:2] == "0x" {
		return numericalString
	}

	u64, err := strconv.ParseUint(numericalString, 10, 64)
	if err != nil {
		panic(err)
	}

	if len(numericalString) < 3 {
		return fmt.Sprintf("0x%x", u64)
	}

	if numericalString[:2] != "0x" {
		return fmt.Sprintf("0x%x", u64)
	}

	return numericalString
}
