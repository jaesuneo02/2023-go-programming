package mymath

// func myAbs(number int) int {

func myAbs(number int) int { // 외부 파일 에 함수를 노출하려면 함수의 이름 첫 글잔느 반드시 대문자로 받는다.
	if number < 0 {
		return number * -1
	}
	return number
}

func MyPower(base int, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}