/*
https://leetcode-cn.com/problems/design-parking-system/
 */
package main

import "fmt"

func main() {
	var (
		system = Constructor(1,1,0)
	)
	fmt.Printf("%t\n", system.AddCar(1))
	fmt.Printf("%t\n", system.AddCar(2))
	fmt.Printf("%t\n", system.AddCar(3))
	fmt.Printf("%t\n", system.AddCar(1))
}

type ParkingSystem struct {
	big, medium, small int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}


func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big == 0 {
			return false
		} else {
			this.big--
			return true
		}
	case 2:
		if this.medium == 0 {
			return false
		} else {
			this.medium--
			return true
		}
	case 3:
		if this.small == 0 {
			return false
		} else {
			this.small--
			return true
		}
	}
	return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */