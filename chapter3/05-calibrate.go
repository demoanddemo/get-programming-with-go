package main

import "fmt"

type Kelvin float64
type Sensor func() Kelvin

func realSensor() Kelvin {
	return 0 //#A
}

func calibrate(sensor Sensor, offset Kelvin) Sensor { //#B
	return func() Kelvin {
		return sensor() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor()) //#C
}

// #A TODO: implement a real sensor
// #B declare and return an anonymous function
// #C Print 5
