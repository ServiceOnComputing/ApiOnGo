package people

import (
	"ApiOnGo/swapi/collection"
	"ApiOnGo/swapi/errors"
)

type People struct {
	errors.ErrorReport
	BirthYear string   `json:"birth_year"`
	EyeColor  string   `json:"eye_color"`
	Films     []string `json:"films"`
	Gender    string   `json:"gender"`
	HairColor string   `json:"hair_color"`
	Height    string   `json:"height"`
	Homeworld string   `json:"homeworld"`
	Mass      string   `json:"mass"`
	Name      string   `json:"name"`
	SkinColor string   `json:"skin_color"`
	Created   string   `json:"created"`
	Edited    string   `json:"edited"`
	Species   []string `json:"species"`
	Starships []string `json:"starships"`
	Url       string   `json:"url"`
	Vehicles  []string `json:"vehicles"`
}
type PeopleCollection struct {
	collection.ResultCollection
	errors.ErrorReport
	Results []People `json:"results"`
}
