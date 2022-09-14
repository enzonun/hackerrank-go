package plusminus

import (
	"math"
)

func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

// func main() {
// var positive float64
// positive += 1
// fmt.Println(RoundDown(positive/6, 6))
// }
