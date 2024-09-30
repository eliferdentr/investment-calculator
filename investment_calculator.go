package main

import (
	"fmt"
	"math"
)

func main() {
	//investment amount, expected annual return and investment horizon are needed
	//	var investmentAmount, numOfYearsToHoldInvestment float64 = 1000 , 10
	var inflationRate  float64 
	var investmentAmount float64
	var expectedReturnRate float64
	var numOfYearsToHoldInvestment int64

	fmt.Print("Investment amount (in decimal): ")
	fmt.Scan(&investmentAmount)
	fmt.Print("\n Expected return rate (in decimal): ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("\nNumber of years to hold the investment: ")
	fmt.Scan(&numOfYearsToHoldInvestment)

	futureValue := investmentAmount * math.Pow(( 1 + expectedReturnRate/100 ), float64(numOfYearsToHoldInvestment))
    futureRealValue := futureValue / math.Pow(1+inflationRate/100 , float64(numOfYearsToHoldInvestment))
	fmt.Println("Future Value: ", futureValue )
	fmt.Println("Future Real Value: ", futureRealValue)
}