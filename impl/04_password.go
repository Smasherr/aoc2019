package impl

import (
	"fmt"
)

func Main4() {
	fmt.Println(CountValidPasswords(284639, 748759))
}

func CountValidPasswords(start int, end int) (int, int) {
	count1, count2 := 0, 0
	for i := start; i <= end; i++ {
		passArray := PasswordIntToIntArr(i)
		if IsValidPassword1(passArray) {
			count1++
			if IsValidPassword2(passArray) {
				count2++
			}
		}
	}
	return count1, count2
}

func IsValidPassword1(password []int) bool {
	fullfillsAdjRule := false
	fullfillsIncRule := false
	for i := 1; i < 6; i++ {
		if password[i-1] == password[i] {
			fullfillsAdjRule = true
		}
		if password[i-1] <= password[i] {
			fullfillsIncRule = true
		} else {
			fullfillsIncRule = false
			break
		}
	}
	return fullfillsAdjRule && fullfillsIncRule
}

func IsValidPassword2(password []int) bool {
	fullfillsAdjRule := false
	countAdj := 0
	current := -1
	for i := 1; i < 6; i++ {
		if password[i-1] == password[i] {
			if current != password[i] {
				countAdj = 1
			}
			current = password[i]
			countAdj++
			if countAdj == 2 && (i == 5 || password[i+1] != current) {
				fullfillsAdjRule = true
			}
		}
	}
	return fullfillsAdjRule
}
