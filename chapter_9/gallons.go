package main

import "fmt"

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%02.f gal", g)
}

type Milliliters float64

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Liters {
	return Liters(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	//var carFuel Gallons
	//var busFuel Liters

	carFuel := Gallons(40.0)
	busFuel := Liters(30.0)

	fmt.Println(carFuel, busFuel)

	//carFuel = Gallons(Liters(40) * 0.264)
	//busFuel = Liters(Gallons(63.0) * 3.785)

	//fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", ToGallons(busFuel), ToLiters(carFuel))

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Println(water)
	fmt.Printf("%0.2f milliliters equals %0.3f gallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters", milk, milk.ToMilliliters())
}
