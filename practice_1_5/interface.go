package main

import . "fmt"

type vehicle interface {
	move()
}

func drive(veh vehicle) {
	veh.move()
}

type Movable interface {
	moveX(distance int)
	moveY(distanse int)
}

type Rectangle struct{}

func (r Rectangle) moveX(distance int) {
	Println("Перемещаем прямоугольник на ", distance, "по оси Х")
}

func (r Rectangle) moveY(distance int) {
	Println("Перемещаем прямоугольник на ", distance, "по оси Y")
}

func move_object(obj Movable) {
	obj.moveX(10)
}

type Car struct{ model string }

type Aircraft struct{ model string }

func (c Car) move() {
	Println(c.model, "едет")
}
func (a Aircraft) move() {
	Println(a.model, "is flying")
}

func main() {

	tesla := Car{"Tesla"}
	volvo := Car{"Volvo"}
	boieng := Aircraft{"Boieng"}

	vehicles := []vehicle{tesla, volvo, boieng}

	for _, vehicle := range vehicles {
		vehicle.move()
	}

	// rect := Rectangle{}
	// move_object(rect)
	// tesla := Car{}
	// boing := Airplane{}
	// Println(tesla)

	// drive(tesla)
	// drive(boing)
	// tesla.move()
	// boing.move()
}
