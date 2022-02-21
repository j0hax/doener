package doener

// Represents a Döner Kebab
type Doener struct {
	Filling  string
	Sauces   []string
	Toppings []string
}

// Standard sauces
var Sauces = []string{"scharf", "Tsatsiki", "Cocktail Soße", "Gelbe Soße", "Joghurt"}

// Standard toppings
var Toppings = []string{
	"Gurken",
	"Zwiebeln",
	"Schafskäse",
	"Tomaten",
	"Koriander",
	"Paprika",
	"Rotkohl",
	"Weißkohl",
	"Karotten",
	"Salat",
	"Mais",
}

// Standard fillings
var Fillings = []string{"Hähnchen", "Kalb", "Veganer", "Lamm", "Falafel"}
