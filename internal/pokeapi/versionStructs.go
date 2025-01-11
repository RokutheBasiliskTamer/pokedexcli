package pokeapi

type versionDetails struct {
	Rate    int
	Version namedAPIResource
}

type versionEncounterDetails struct {
	Encounter_details []encounter
	Version           namedAPIResource
	Max_chance        int
}

type versionGameIndex struct {
	Game_index int
	Version    namedAPIResource
}
