package main

import (
	"fmt"
	"math"
)

const underWeightMessage = "You are underweight."
const normalMessage = "You are healthy."
const overWeightMessage = "You are overweight."
const obeseIMessage = "You are classified as Class I obesity."
const obeseIIMessage = "You are classified as Class II obesity."
const obeseIIIMessage = "You are classified as Class III obesity."

type WeightAndHeight struct {
	weight float64
	height float64
}

func CalculateBMI(w WeightAndHeight) float64 {
	return w.weight / math.Pow(w.height, 2)
}

func ClassifyBMIToText(bmi float64) string {
	if bmi < 18.5 {
		return underWeightMessage
	} else if bmi >= 18.5 && bmi <= 24.9 {
		return normalMessage
	} else if bmi > 24.9 && bmi <= 29.9 {
		return overWeightMessage
	} else if bmi > 29.9 && bmi <= 34.9 {
		return obeseIMessage
	} else if bmi > 34.9 && bmi <= 39.9 {
		return obeseIIMessage
	} else if bmi > 39.9 {
		return obeseIIIMessage
	}
	return ""
}

func main() {
	var weight, height float64

	fmt.Println("Please enter your weight:")
	_, err := fmt.Scanf("%f\n", &weight)
	checkError(err)

	fmt.Println("Please enter your height:")
	_, err = fmt.Scanf("%f", &height)
	checkError(err)

	bmi := CalculateBMI(WeightAndHeight{weight, height})
	text := ClassifyBMIToText(bmi)
	fmt.Printf("Your BMI is %.2f, %v\n", bmi, text)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
