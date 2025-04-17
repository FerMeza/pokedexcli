package pokedex

var Pokedex map[string]Pokemon = map[string]Pokemon{}

type Pokemon struct {
	Name           string
	BaseExperience int
}
