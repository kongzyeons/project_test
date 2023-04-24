package main

import (
	"fmt"
	"strconv"
)

func rule_1(number string) bool {
	if len(number) < 6 {
		return false
	}
	return true
}
func rule_2(number string) bool {
	for i := range number {
		if i+2 == len(number)-1 {
			break
		}
		if (number[i] == number[i+1]) && (number[i] == number[i+2]) {
			return false
		}
	}
	return true
}
func rule_3(number string) bool {
	for i := range number {
		if i+2 == len(number)-1 {
			break
		}
		if (number[i] == number[i+1]-1) && (number[i] == number[i+2]-2) {
			return false
		} else if (number[i] == number[i+1]+1) && (number[i] == number[i+2]+2) {
			return false
		}
	}
	return true
}
func rule_4(number string) bool {
	check := 0
	for i, v1 := range number {
		if i == len(number)-1 {
			break
		}
		for _, v2 := range number[i+1:] {
			if v1 == v2 {
				check += 1
				if check > 2 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	var number string
	fmt.Print("Input a number: ")
	fmt.Scanf("%s", &number)

	_, err := strconv.Atoi(number)
	if err != nil {
		err = fmt.Errorf("Input must be integer")
		fmt.Println(false, err)
	} else if !rule_1(number) {
		err = fmt.Errorf("Rule 1 : Input must be longer more than 6")
		fmt.Println(false, err)
	} else if !rule_2(number) {
		err = fmt.Errorf("Rule 2 : Input must not be duplicate more than 2 consecutive number.")
		fmt.Println(false, err)
	} else if !rule_3(number) {
		err = fmt.Errorf("Rule 3 : Input of sort must not be more than 2 contiguous  number.")
		fmt.Println(false, err)
	} else if !rule_4(number) {
		err = fmt.Errorf("Rule 4 : Input must not have more than 2 sets of repeating number.")
		fmt.Println(false, err)
	} else {
		fmt.Println(true, "success")
	}

}
