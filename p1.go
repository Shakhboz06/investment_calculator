package main

import(
	"fmt"
	"math"
)

const inflation float64 = 4.3 
func main() {
	fmt.Println("Hello World")
 
	// var investmentAmount, years float64= 2000, 7 
	// var investmentAmount float64 = 2000 // or in case we use scan we can do below:
	var investmentAmount float64; 
	// years := 7.0
	var years float64; 
	// expectedReturn := 5.5
	var expectedReturn float64; 
	// difference between var(can be changed) and const (unchanged)
	investmentAmount = 2200

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years: ")
	fmt.Scan(&years)
	fmt.Print("Expected return rate: ")	
	fmt.Scan(&expectedReturn)

	calcfutureValues(investmentAmount, expectedReturn, years)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturn / 100, (years))
	// futureRealValue :=  futureValue / math.Pow(1 + inflation / 100, years)
	futureValue, futureRealValue := calcfutureValues(investmentAmount, expectedReturn, years)
	// futureRealValue :=  futureValue / math.Pow(1 + inflation / 100, years)


	// fmt.Printf("Future value: %v\nFuture Real Value: %v",futureValue, futureRealValue) 
	fmt.Printf("Future value: %.2f\n", futureValue) 
	// fmt.Print("Future real value:", futureRealValue)
	
	
	// Storing into the variable
	varFutureRealvaLue := fmt.Sprintf("Future real value: %.2f\n", futureRealValue)
	fmt.Println((varFutureRealvaLue))
	
	// with backsticks
	// fmt.Printf(`Future value: %.2f 
	// Future Real Value: %.2f`,futureValue, futureRealValue) 

}

// func main() {
// 	fmt.Println("Hello World")


// 	investmentAmount, years, expectedReturn := 2000.0, 7.0, 5.5
// 	futureValue := investmentAmount * math.Pow(1 + expectedReturn / 100, (years))

// 	fmt.Println(futureValue)

// }

func calcfutureValues (investmentAmount, expectedReturn, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturn / 100, (years))	
	rfv := fv / math.Pow(1 + inflation / 100, years)
	return fv, rfv
	
}

// alternative approach
// func calcfutureValues (investmentAmount, expectedReturn, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1 + expectedReturn / 100, (years))	
// 	rfv = fv / math.Pow(1 + inflation / 100, years)
// 	return 

// }