package main

import "fmt"


const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal uint16
	brakePedal uint16
	steeringWheel int16
	topSpeedKmh float64
}
// method of struct: value receiver
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh/usixteenbitmax)
}
// method of struct: value receiver
func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh/usixteenbitmax/kmhMultiple)
}

// method of struct: pointer receiver to change original struct value
func (c *car) newTopSpeed(newSpeed float64)  {
	c.topSpeedKmh = newSpeed
}


func main() {
	aCar := car{
		gasPedal: 62215,
		brakePedal: 10,
		steeringWheel: 1526,
		topSpeedKmh: 225,
	}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar.newTopSpeed(100)
	fmt.Println(aCar.topSpeedKmh)
}