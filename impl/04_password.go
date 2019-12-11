package impl

import (
	"fmt"
)

// Main4 solves Day4
func Main4() {
	fmt.Println(CountValidPasswords(284639, 748759))
}

// CountValidPasswords returns the amount of passwords fulfilling requirements for parts 1 and 2
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

// IsValidPassword1 determines if password fulfills requirements for the part 1
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

// IsValidPassword2 determines if password fulfills additional requirements for the part 2
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
