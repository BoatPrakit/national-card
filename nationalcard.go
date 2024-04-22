package nationalcard

import (
	"errors"
	"strconv"
	"strings"
)

func ValidateThaiID(id string) error {
	if len(id) < 13 {
		return errors.New("id digits incorrect")
	}
	var sum int

	digit := strings.Split(id, "")
	for i, j := 0, 13; i < 12; i++ {
		num, _ := strconv.Atoi(digit[i])
		sum += num * j
		j--
	}
	mod := sum % 11
	sub := 11 - mod
	comparer := sub % 10

	if strconv.Itoa(comparer) != string(id[12]) {
		return errors.New("id incorrect")
	}

	return nil
}
