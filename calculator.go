package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var calculatorErrors []string

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ParseString(enter string) int {
	num, err := strconv.Atoi(enter)

	if err != nil {
		_, err1 := strconv.ParseFloat(enter, 32)
		if err1 == nil {
			calculatorErrors = append(calculatorErrors, "Можно вводить только целые числа.\n")
			return num
		} else {
			calculatorErrors = append(calculatorErrors, "Вы ввели неверное значение "+enter+"\n")
			return num
		}
	} else {
		return num
	}
}

func getResult(number1 int, number2 int, operator string) int {
	switch operator {

	case "+":
		return number1 + number2

	case "-":
		return number1 - number2

	case "*":
		return number1 * number2

	case "/":
		return number1 / number2

	}

	return 0
}

func getEnter() (string, string, string) {
	var enter1 string
	var enter2 string
	var operator string

	fmt.Print("Начните ввод : ")
	fmt.Scanf("%s %s %s", &enter1, &operator, &enter2)

	return enter1, enter2, operator
}

func checkOperator(operators []string, operator string) {
	isContains := Contains(operators, operator)
	if !isContains {
		calculatorErrors = append(calculatorErrors, "Вы ввели неверный оператор, оператор должен быть один из "+strings.Join(operators, ", ")+"\n")

	}
}

func isCorrectNumberSystem(isRim1 bool, isRim2 bool) {
	if (isRim1 && !isRim2) || (!isRim1 && isRim2) {
		calculatorErrors = append(calculatorErrors, "Можно использовать одновременно одинаковые системы исчисления. \n")
	}
}

func checkNumberSize(number1 int, number2 int) {
	if number1 > 10 || number2 > 10 {
		calculatorErrors = append(calculatorErrors, "Введите число от 1 до 10. \n")

	}
}

func isDivisionByZero(number int, operator string) {
	if number == 0 && operator == "/" {
		calculatorErrors = append(calculatorErrors, "Делить на нуль нельзя \n")
	}
}

func main() {
	var enter1, enter2 string
	var number1, number2 int
	var operator string
	var isRim1 bool
	var isRim2 bool
	var result int
	var operators = []string{"+", "-", "/", "*"}
	var ErrRimResult = errors.New("При использовании римской системы исчисления, результат не может быть меньше I \n")

	mapConvertArabToRim := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}

	enter1, enter2, operator = getEnter()

	for key, value := range mapConvertArabToRim {
		if value == enter1 {
			number1 = key
			isRim1 = true
		}
		if value == enter2 {
			number2 = key
			isRim2 = true
		}
	}

	if !isRim1 {
		number1 = ParseString(enter1)
	}

	if !isRim2 {
		number2 = ParseString(enter2)
	}

	checkOperator(operators, operator)
	isCorrectNumberSystem(isRim1, isRim2)
	checkNumberSize(number1, number2)
	isDivisionByZero(number2, operator)

	if len(calculatorErrors) != 0 {
		fmt.Printf("При выполнении команды были выявлены ошибки \n %s", strings.Join(calculatorErrors, " "))
		return
	}

	result = getResult(number1, number2, operator)

	if isRim1 && isRim2 {
		if result < 1 {
			panic(ErrRimResult)
		}
		fmt.Print(mapConvertArabToRim[result])
	} else {
		fmt.Print(result)

	}
}
