package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	expression, _ := reader.ReadString('\n')

	result := calculate(expression)
	fmt.Printf(result)
}

func calculate(expression string) string {
	var romanflag bool
	var operand1 int64
	var operand2 int64
	var err error
	var calc int64
	var result string
	expression = strings.TrimSuffix(expression, "\r\n")
	expression = strings.TrimSuffix(expression, "\n")
	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		fmt.Println("Invalid expression")
		os.Exit(1)
	}

	// Cheking if roman
	switch tokens[0][0:1] {
	case "I", "V", "X":
		romanflag = true
	default:
		romanflag = false
	}

	switch tokens[2][0:1] {
	case "I", "V", "X":
		if !romanflag {
			fmt.Println("First operand is not roman")
			os.Exit(1)
		}

	default:
		if romanflag {
			fmt.Println("Second operand is not roman")
			os.Exit(1)
		}
	}
	// Saving operator
	operator := tokens[1]

	if romanflag {
		operand1 = romanToInt(tokens[0])
		operand2 = romanToInt(tokens[2])
	} else {
		operand1, err = strconv.ParseInt(tokens[0], 10, 32)
		if err != nil {
			fmt.Println("Invalid operand 1:", err)
			os.Exit(1)
		}

		operand2, err = strconv.ParseInt(tokens[2], 10, 32)
		if err != nil {
			fmt.Println("Invalid operand 2:", err)
			os.Exit(1)
		}
	}

	//Cheking if numbers are [0;10]
	if (operand1 > 10) || (operand1 < 1) {
		fmt.Println("Invalid range of operand 1")
		os.Exit(1)
	}
	if (operand2 > 10) || (operand2 < 1) {
		fmt.Println("Invalid range of operand 2")
		os.Exit(1)
	}

	//Calculation
	switch operator {
	case "+":
		calc = operand1 + operand2
	case "-":
		calc = operand1 - operand2
		if (calc < 1) && (romanflag) {
			fmt.Println("Invalid roman result")
			os.Exit(1)
		}
	case "*":
		calc = operand1 * operand2
	case "/":
		calc = operand1 / operand2
	default:
		fmt.Println("Invalid operator:", operator)
		os.Exit(1)
	}

	//Preparation for returning result

	if romanflag {
		result = intToRoman(calc)
	} else {
		result = strconv.FormatInt(calc, 10)
	}
	return result
}

// Roman converter
func romanToInt(s string) int64 {
	var sum int64
	sum = 0

	rom := map[string]int64{
		"I": 1,
		"V": 5,
		"X": 10,
	}

	for i, v := range s {
		sum += rom[string(v)]
		if i != 0 {
			if rom[string(s[i-1])] < rom[string(v)] {
				sum -= 2 * rom[string(s[i-1])]
			}
		}
	}

	return sum
}

func intToRoman(num int64) string {
	res := ""
	for num >= 10 {
		num -= 10
		res += "X"
	}
	for num >= 9 {
		num -= 9
		res += "IX"
	}
	for num >= 5 {
		num -= 5
		res += "V"
	}
	for num >= 4 {
		num -= 4
		res += "IV"
	}
	for num >= 1 {
		num -= 1
		res += "I"
	}
	return res
}
