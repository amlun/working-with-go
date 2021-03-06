/**
 * object-oriented-example.go
 *
 * A basic example create on object and defining a method
 */

package main

import "fmt"

// define Dog object type
type Dog struct {
	Name  string
	Color string
	Owner string
}

// Method for object specify the object the refer to and then
// the method name and rest of normal function definition
func (d Dog) Call() {
	fmt.Printf("Here comes a %s dog, %s.\n", d.Color, d.Name)
}

func (d Dog) Shout() {
	fmt.Println("Awang! Awang!")
}

func (d Dog) BelongsTo() {
	fmt.Printf("This dog %s is belongs to %s!\n", d.Name, d.Owner)
}

func main() {

	// create instance of object and set properties
	Spot := Dog{Name: "Spot", Color: "brown"}

	// or, try it like this...
	var Rover Dog
	Rover.Name = "Rover"
	Rover.Color = "light blondish with some dark patches and very scruffy"

	// call object method
	Spot.Call()
	Rover.Call()

	Spot.Shout()
	Rover.Shout()

	Spot.Owner = "lunweiwei"
	Spot.BelongsTo()
}
