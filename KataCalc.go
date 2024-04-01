package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"errors"
)

func Input() string {
	str := bufio.NewReader(os.Stdin)
	text, _ := str.ReadString('\n')
	return text
}

func Output(s string) {
	fmt.Println(s)
}

func FromIntToRom(i int) string {
	var s = []rune("")
	for i > 0 {
		if i == 100 {
			s = append(s, 'C')
			i -= 100
		} else if i > 89 {
			s = append(s, 'X')
			s = append(s, 'C')
			i -= 90
		} else if i > 49 {
			s = append(s, 'L')
			i -= 50
		} else if i > 39 {
			s = append(s, 'X')
			s = append(s, 'L')
			i -= 40
		} else if i > 9 {
			s = append(s, 'X')
			i -= 10
		} else if i == 9 {
			s = append(s, 'I')
			s = append(s, 'X')
			i -= 9
		} else if i > 4 {
			s = append(s, 'V')
			i -= 5
		} else if i == 4 {
			s = append(s, 'I')
			s = append(s, 'V')
			i -= 4
		} else if i >= 1 {
			s = append(s, 'I')
			i -= 1
		} else {
			break
		}
	}
	return string(s)
}

var ToIntFromRom = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func RomOrArabAndCheckForCorrect(s string) string {
	words := strings.Fields(s)
	var RomCnt, ArabCnt, ErrCnt, incorrect int = 0, 0, 0, 0
	for _, word := range words {
		word = strings.TrimSpace(word)
		if strings.Contains(word, "I") == true || strings.Contains(word, "V") == true || strings.Contains(word, "X") == true {
			if _, ok := ToIntFromRom[word]; ok {
				RomCnt++
			} else {
				incorrect++
				break
			}
		} else if strings.ContainsAny(word, "0123456789") == true {
			if dig, _ := strconv.Atoi(word); dig > 0 && dig < 11 && reflect.TypeOf(dig) == reflect.TypeOf(ArabCnt) {
				ArabCnt++
			} else {
				incorrect++
				break
			}
		} else {
			ErrCnt++
			break
		}
	}
	if RomCnt == 2 && ErrCnt == 0 && ArabCnt == 0 {
		return "Rom"
	} else if ArabCnt == 2 && ErrCnt == 0 && RomCnt == 0 {
		return "Arab"
	} else if incorrect > 0 {
		return "incorrect"
	} else {
		return "Err"
	}
}

func EliminationFromOperation(s string) string {
	operation, countOfOperation := CheckTypeOfOperation(s)
	if countOfOperation == 1 {
		a, b, _ := strings.Cut(s, operation)
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)
		text := []rune(a + " " + b)
		return string(text)
	}
	return "Err"
}

func CheckTypeOfOperation(s string) (string, int) {
	var TotalCnt int = 0
	var TypeOfOperation = []rune("")
	if strings.Contains(s, "+") && strings.Count(s, "+") == 1 {
		TypeOfOperation = append(TypeOfOperation, '+')
		TotalCnt++
	}
	if strings.Count(s, "+") > 1 {
		TotalCnt++
	}
	if strings.Contains(s, "-") && strings.Count(s, "-") == 1 {
		TypeOfOperation = append(TypeOfOperation, '-')
		TotalCnt++
	}
	if strings.Count(s, "-") > 1 {
		TotalCnt++
	}
	if strings.Contains(s, "/") && strings.Count(s, "/") == 1 {
		TypeOfOperation = append(TypeOfOperation, '/')
		TotalCnt++
	}
	if strings.Count(s, "/") > 1 {
		TotalCnt++
	}
	if strings.Contains(s, "*") && strings.Count(s, "*") == 1 {
		TypeOfOperation = append(TypeOfOperation, '*')
		TotalCnt++
	}
	if strings.Count(s, "+") > 1 {
		TotalCnt++
	}
	if TotalCnt == 1 {
		return string(TypeOfOperation), TotalCnt
	} else {
		return "Err", TotalCnt
	}
}

func RomOperation(operands, operation string) string {
	words := strings.Fields(operands)
	num1 := ToIntFromRom[strings.TrimSpace(words[0])]
	num2 := ToIntFromRom[strings.TrimSpace(words[1])]
	if operation == "+" {
		return FromIntToRom(num1 + num2)
	}
	if operation == "-" {
		if num1-num2 >= 1 {
			return FromIntToRom(num1 - num2)
		} else {
			panic(errors.New("В римской системе нет отрицательных чисел."))
		}
	}
	if operation == "*" {
		return FromIntToRom(num1 * num2)
	}
	if operation == "/" {
		if num1/num2 >= 1 {
			return FromIntToRom(num1 / num2)
		} else {
			panic(errors.New("Не сказано что выводить в случае, если результат деления римских цифр меньше единицы."))
		}
	}
	return "You are never to see this string in console, it seems it called \"syntaxyx sugar\":)"
}

func ArabOperation(operands, operation string) string {
	words := strings.Fields(operands)
	num1, _ := strconv.Atoi(strings.TrimSpace(words[0]))
	num2, _ := strconv.Atoi(strings.TrimSpace(words[1]))
	if operation == "+" {
		return strconv.Itoa(num1 + num2)
	}
	if operation == "-" {
		return strconv.Itoa(num1 - num2)
	}
	if operation == "*" {
		return strconv.Itoa(num1 * num2)
	}
	if operation == "/" {
		return strconv.Itoa(num1 / num2)
	}
	return "You are never to see this string in console, it seems it called \"syntaxyx sugar\":)"
}

func main() {

	for {
		goal := Input()
		operation, countOfOperation := CheckTypeOfOperation(goal)
		if operation == "Err" && countOfOperation == 0 {
			panic(errors.New("Cтрока не является математической операцией."))
		}
		if countOfOperation > 1 {
			panic(errors.New("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."))
		}
		typeOfSys := RomOrArabAndCheckForCorrect(EliminationFromOperation(goal))
		if typeOfSys == "Rom" {
			operands := EliminationFromOperation(goal)
			Output(RomOperation(operands, operation))
		} else if typeOfSys == "Arab" {
			operands := EliminationFromOperation(goal)
			Output(ArabOperation(operands, operation))
		} else if typeOfSys == "incorrect" {
			panic(errors.New("Некорректный ввод числа."))
		} else if typeOfSys == "Err" {
			panic(errors.New("Используются одновременно разные системы счисления."))
		}
	}
}
