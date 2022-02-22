package doener

import (
	"math/rand"
)

/*
	choose chooses one or more items at random from a given array
*/
func choose(array []string) []string {
	numItems := rand.Intn(len(array)) + 1

	rand.Shuffle(numItems, func(i, j int) {
		array[i], array[j] = array[j], array[i]
	})

	return array[0:numItems]
}

/*
	randSauces chooses one or more of the available sauces
*/
func randSauces() []string {
	return choose(Sauces)
}

/*
	randToppings chooses one or more of the of the available toppings
*/
func randToppings() []string {
	return choose(Toppings)
}

/*
	randFilling chooses one of the available fillings
*/
func randFilling() string {
	index := rand.Intn(len(Fillings))
	return Fillings[index]
}

/*
	String returns a textual description of the Döner Kebab.
*/
func (d Doener) String() string {
	// Basic type
	result := d.Filling + " Döner mit "

	// list sauces
	sauceCount := len(d.Sauces)
	for i := 0; i < sauceCount-1; i++ {
		result += d.Sauces[i] + ", "
	}

	result += d.Sauces[sauceCount-1] + " und "

	// list toppings
	toppingCount := len(d.Toppings)
	for i := 0; i < toppingCount-1; i++ {
		result += d.Toppings[i] + ", "
	}

	result += d.Toppings[toppingCount-1]

	return result
}

/*
	WithAll returns a Doener with the first standard filling, all toppings and sauces
*/
func WithAll() *Doener {
	return &Doener{Fillings[0], Sauces, Toppings}
}

/*
	Random returns a randomly selected Döner Kebab.
	One filling protein, and any combination of toppings and sauces is chosen.
*/
func Random() *Doener {
	sauces := randSauces()
	toppings := randToppings()
	filling := randFilling()

	return &Doener{filling, sauces, toppings}
}
