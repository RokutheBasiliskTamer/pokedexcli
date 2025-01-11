package pokeapi

type PaginationResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []namedAPIResource
}

type namedAPIResource struct {
	Name string
	Url  string
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

type pokemonEncounter struct {
	Pokemon         namedAPIResource
	Version_details []versionEncounterDetails
}

type encounterMethodRate struct {
	Encounter_method namedAPIResource
	Version_details  []versionDetails
}

type versionDetails struct {
	Rate    int
	Version namedAPIResource
}

type versionEncounterDetails struct {
	Encounter_details []encounter
	Version           namedAPIResource
	Max_chance        int
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

type Config struct {
	Next     *string
	Previous *string
	Client   *Client
}
