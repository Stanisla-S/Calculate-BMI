package main

import (
	"errors"
	"fmt"
	"math"
)

const IBMPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")
	for {
		userKg, userHeight := getUserInput()
		IBM, err := calculateIBM(userKg, userHeight)
		if err != nil {
			// fmt.Println("Не заданы параметры для расчета")
			// continue
			panic("Не заданы параметры для расчета")
		}
		outputResult(IBM)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}

}
func outputResult(ibm float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", ibm)
	fmt.Println(result)
	switch {
	case ibm < 16:
		fmt.Println("У вас сильный дефицит массы тела")

	case ibm < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case ibm < 25:
		fmt.Println("У вас нормальный вес")
	case ibm < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}

}

func calculateIBM(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IBM := userKg / math.Pow(userHeight/100, IBMPower)
	return IBM, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчет? (y/n):  ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
