package species

import (
	"ApiOnGo/swapi/collection"
	"ApiOnGo/swapi/errors"
)

type Species struct {
	errors.ErrorReport
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairColors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	AverageLifespan string   `json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	People          []string `json:"people"`
	Films           []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	Url             string   `json:"url"`
}
type SpeciesCollection struct {
	errors.ErrorReport
	collection.ResultCollection
	Results []Species `json:"results"`
}
