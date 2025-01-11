package pokeapi

import "encoding/json"

type PaginationResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []namedAPIResource
}

func (p *PaginationResponse) UnByteify(byteData []byte) error {
	//unmarshal byte data into struct for pagination requests
	if err := json.Unmarshal(byteData, p); err != nil {
		return err
	}
	return nil
}

type LocationArea struct {
	Encounter_method_rates []encounterMethodRate
	Game_index             int
	Id                     int
	Location               namedAPIResource
	Name                   string
	Names                  []name
	Pokemon_encounters     []pokemonEncounter
}

func (l *LocationArea) UnByteify(byteData []byte) error {
	//unmarshal byte data into struct for pagination requests
	if err := json.Unmarshal(byteData, l); err != nil {
		return err
	}
	return nil
}

type namedAPIResource struct {
	Name string
	Url  string
}

type pokemonEncounter struct {
	Pokemon         namedAPIResource
	Version_details []versionEncounterDetails
}

type encounterMethodRate struct {
	Encounter_method namedAPIResource
	Version_details  []versionDetails
}

type encounter struct {
	Chance           int
	Condition_values []namedAPIResource
	Max_level        int
	Method           namedAPIResource
	Min_level        int
}

type name struct {
	Name     string
	Language namedAPIResource
}
