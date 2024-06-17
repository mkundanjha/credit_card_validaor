package models

func CheckSumVerifier(inputNumber int64) bool {
	//loop in array
	//for odd places digit calculate the sum
	//for even places digit multyply digit by 2 then
	// if digit is greater than 9 digit=digit-9
	// find the sum of new even digit
	sumOdd, sumEven := 0, 0

	flag := 1
	for inputNumber > 0 {
		modDigit := inputNumber % 10

		//For even condition
		if flag%2 == 0 {
			modDigit = 2 * modDigit
			if modDigit > 9 {
				modDigit = modDigit - 9
			}
			sumEven = sumEven + int(modDigit)
		} else {
			sumOdd = sumOdd + int(modDigit)
		}
		//Reduce the input number
		inputNumber = inputNumber / 10
		flag = flag + 1
	}

	finalSum := sumEven + sumOdd
	// fmt.Println("The checksum is ", finalSum)

	if finalSum%10 == 0 {
		return true
	} else {
		return false
	}

}
