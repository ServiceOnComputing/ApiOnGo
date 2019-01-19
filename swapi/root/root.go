package root

import "ApiOnGo/swapi/errors"

type RootCollection struct {
	errors.ErrorReport
	Films     string `json:"films"`
	People    string `json:"people"`
	Planets   string `json:"planets"`
	Species   string `json:"species"`
	Starships string `json:"starships"`
	Vehicles  string `json:"vehicles"`
}
