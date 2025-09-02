package main

import (
	"fmt"
)

// ===== Abstraction =====
type Vehicle interface {
	GetName() string
	GetSpeed() int
	SetSpeed(speed int)
	IsRunning() bool
	Drive()
}

// ===== Encapsulation =====
type BaseVehicle struct {
	name  string
	speed int
	fuel  int
}

func (v *BaseVehicle) GetName() string {
	return v.name
}

func (v *BaseVehicle) GetSpeed() int {
	return v.speed
}

func (v *BaseVehicle) SetSpeed(speed int) {
	v.speed = speed
}

func (v *BaseVehicle) IsRunning() bool {
	return v.fuel > 0
}

// ===== Inheritance + Polymorphism =====

// Car
type Car struct {
	BaseVehicle
}

func (c *Car) Drive() {
	fmt.Printf("%s drives smoothly at speed %d km/h.\n", c.name, c.speed)
}

// Truck
type Truck struct {
	BaseVehicle
}

func (t *Truck) Drive() {
	fmt.Printf("%s hauls heavy load at speed %d km/h.\n", t.name, t.speed)
}

// Motorcycle
type Motorcycle struct {
	BaseVehicle
}

func (m *Motorcycle) Drive() {
	fmt.Printf("%s zooms through traffic at speed %d km/h.\n", m.name, m.speed)
}

// ElectricCar
type ElectricCar struct {
	BaseVehicle
}

func (e *ElectricCar) Drive() {
	fmt.Printf("%s glides silently at speed %d km/h.\n", e.name, e.speed)
}

// ===== Main =====
func main() {
	car := &Car{BaseVehicle{"Sedan Car", 80, 50}}
	truck := &Truck{BaseVehicle{"Big Truck", 60, 70}}
	bike := &Motorcycle{BaseVehicle{"Sport Bike", 120, 20}}
	ev := &ElectricCar{BaseVehicle{"Tesla Model X", 100, 90}}

	// Demonstrate polymorphism
	var vehicle Vehicle

	vehicle = car
	vehicle.Drive()

	vehicle = truck
	vehicle.Drive()

	vehicle = bike
	vehicle.Drive()

	vehicle = ev
	vehicle.Drive()
}
