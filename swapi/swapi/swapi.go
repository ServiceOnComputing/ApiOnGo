package swapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"ApiOnGo/swapi/film"
	"ApiOnGo/swapi/people"
	"ApiOnGo/swapi/planet"
	"ApiOnGo/swapi/root"
	"ApiOnGo/swapi/species"
	"ApiOnGo/swapi/starship"
	"ApiOnGo/swapi/vehicle"
)

/* Base URL used when communicating with the SWAPI RESTful services */
var BaseUrl string = "http://swapi.co/api"

/* Main structure housing API wrapper methods. */
type SWAPIClient struct {
}

/*
Returns a new instance of an SWAPIClient.
*/
func NewClient() *SWAPIClient {
	return &SWAPIClient{}
}

func (this *SWAPIClient) buildUrl(endpoint string) string {
	return fmt.Sprintf("%s%s", BaseUrl, endpoint)
}

func (this *SWAPIClient) get(endpoint string) (int, []byte, error) {
	response, err := http.Get(this.buildUrl(endpoint))
	if err != nil {
		return 0, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		return 0, nil, err
	}

	return response.StatusCode, body, nil
}

/*
Returns a FilmCollection structure which has a Results array
of all films available from SWAPI
*/
func (this *SWAPIClient) GetAllFilms() (*film.FilmCollection, int, error) {
	result := &film.FilmCollection{}

	status, body, err := this.get("/films/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a PeopleCollection structure which has a Results array
of all persons available from SWAPI.
*/
func (this *SWAPIClient) GetAllPeople() (*people.PeopleCollection, int, error) {
	result := &people.PeopleCollection{}

	status, body, err := this.get("/people/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a PlanetCollection structure which has a Results array
of all planets available from SWAPI.
*/
func (this *SWAPIClient) GetAllPlanets() (*planet.PlanetCollection, int, error) {
	result := &planet.PlanetCollection{}

	status, body, err := this.get("/planets/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a SpeciesCollection structure which has a Results array
of all the types of species in SWAPI.
*/
func (this *SWAPIClient) GetAllSpecies() (*species.SpeciesCollection, int, error) {
	result := &species.SpeciesCollection{}

	status, body, err := this.get("/species/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a StarshipCollection structure which has a Results array
of all the types of starships available in SWAPI.
*/
func (this *SWAPIClient) GetAllStarships() (*starship.StarshipCollection, int, error) {
	result := &starship.StarshipCollection{}

	status, body, err := this.get("/starships/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a VehicleCollection structure which has a Results array
of all the types of vehicles available in SWAPI.
*/
func (this *SWAPIClient) GetAllVehicles() (*vehicle.VehicleCollection, int, error) {
	result := &vehicle.VehicleCollection{}

	status, body, err := this.get("/vehicles/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a Root structure which contains a list of resources available
in the SWAPI.
*/
func (this *SWAPIClient) GetAvailableResources() (*root.RootCollection, int, error) {
	result := &root.RootCollection{}

	status, body, err := this.get("/")
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a Film structure matching the provided ID.
*/
func (this *SWAPIClient) GetFilmById(id int) (*film.Film, int, error) {
	result := &film.Film{}

	status, body, err := this.get(fmt.Sprintf("/films/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetfilmById(id int) (string) {
	_, body, err := this.get(fmt.Sprintf("/films/%d", id))
	if err != nil {
		return "error"
	}
	return string(body)
}


/*
Returns a People structure matching the provided ID.
*/
func (this *SWAPIClient) GetPersonById(id int) (*people.People, int, error) {
	result := &people.People{}

	status, body, err := this.get(fmt.Sprintf("/people/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

func (this *SWAPIClient) GetpersonById(id int) (string) {
	//result := &people.People{}

	_, body, err := this.get(fmt.Sprintf("/people/%d", id))
	if err != nil {
		return "error"
	}
	return string(body)

	//err = json.Unmarshal(body, result)
	//return result, status, err
}

/*
Returns a Planet structure matching the provided ID.
*/
func (this *SWAPIClient) GetPlanetById(id int) (*planet.Planet, int, error) {
	result := &planet.Planet{}

	status, body, err := this.get(fmt.Sprintf("/planets/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a Species structure matching the provided ID.
*/
func (this *SWAPIClient) GetSpeciesById(id int) (*species.Species, int, error) {
	result := &species.Species{}

	status, body, err := this.get(fmt.Sprintf("/species/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a Starship structure matching the provided ID.
*/
func (this *SWAPIClient) GetStarshipById(id int) (*starship.Starship, int, error) {
	result := &starship.Starship{}

	status, body, err := this.get(fmt.Sprintf("/starships/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}

/*
Returns a Vehicle structure matching the provided ID.
*/
func (this *SWAPIClient) GetVehicleById(id int) (*vehicle.Vehicle, int, error) {
	result := &vehicle.Vehicle{}

	status, body, err := this.get(fmt.Sprintf("/vehicles/%d", id))
	if err != nil {
		return result, 0, err
	}

	err = json.Unmarshal(body, result)
	return result, status, err
}
