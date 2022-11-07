package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type VeggeMania struct {
}

type TomatoTopping struct {
	pizza IPizza
}

type CheeseTopping struct {
	pizza IPizza
}

func (p *VeggeMania) getPrice() int {
	return 15
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func main() {

	pizza := &VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}
	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
	//Price of veggeMania with tomato and cheese topping is 32
}
