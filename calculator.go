package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	op1, operation, op2, formatFlag := enterData()
	if formatFlag {
		num := calculate(op1, op2, operation)
		if num <= 0 {
			throwException("результат работы с римскими цифрами меньше единицы;")
		}
		fmt.Println(arabicToRoman(num))
	} else {
		fmt.Println(calculate(op1, op2, operation))
	}
}

func throwException(text string) {
	fmt.Println(errors.New(text))
	os.Exit(1)
}

func checkRomanNumber(number string) bool {
	chars := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, c := range chars {
		if number == c {
			return true
		}
	}
	return false
}

func enterData() (int, string, int, bool) {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	f := strings.Split(data, " ")
	romanFlag := false

	if len(f) != 3 {
		throwException("введено не верное кол-во операндов (должно быть два)")
	}

	op1, err := strconv.Atoi(f[0])

	if err != nil {
		if !checkRomanNumber(f[0]) {
			throwException("первый аргумент не является числом ни арабским, ни римским!")
		} else {
			op1 = romanToArabic(f[0])
			romanFlag = true
		}
	}

	operation := f[1]
	op2, err := strconv.Atoi(f[2])

	if (!romanFlag && checkRomanNumber(f[2])) || (romanFlag && !checkRomanNumber(f[2])) {
		throwException("кодировка операндов не совпадают!")
	}

	if err != nil {
		if !checkRomanNumber(f[2]) {
			throwException("второй аргумент не является числом ни арабским, ни римским!")
		} else {
			if !romanFlag {
				throwException("")
			}
			op2 = romanToArabic(f[2])
		}
	}

	return op1, operation, op2, romanFlag
}

func romanToArabic(number string) int {
	conv := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	res, end := conv[string(number[len(number)-1])], len(number)-1
	var arabian int

	for i := end - 1; i >= 0; i-- {
		arabian = conv[string(number[i])]

		if arabian < conv[string(number[i+1])] {
			res -= arabian
		} else {
			res += arabian
		}
	}

	return res
}

func arabicToRoman(number int) string {
	var conversions = []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func calculate(op1 int, op2 int, operation string) int {
	res := 0
	switch operation {
	case "+":
		res = op1 + op2
	case "-":
		res = op1 - op2
	case "*":
		res = op1 * op2
	case "/":
		res = op1 / op2
	}
	return res
}
