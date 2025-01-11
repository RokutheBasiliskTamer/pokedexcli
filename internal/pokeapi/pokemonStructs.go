package pokeapi

import "encoding/json"

type Pokemon struct {
	Id                       int
	Name                     string
	Base_experience          int
	Height                   int
	Is_default               bool
	Order                    int
	Weight                   int
	Abilities                []pokemonAbility
	Forms                    []namedAPIResource
	Game_indices             []versionGameIndex
	Held_items               []pokemonHeldItem
	Location_area_encounters string
	Moves                    []pokemonMove
	Past_types               []pokemonTypePast
	Sprites                  pokemonSprites
	Cries                    pokemonCries
	Species                  namedAPIResource
	Stats                    []pokemonStat
	Types                    []pokemonType
}

func (p *Pokemon) UnByteify(byteData []byte) error {
	//unmarshal byte data into struct for pagination requests
	if err := json.Unmarshal(byteData, p); err != nil {
		return err
	}
	return nil
}

type pokemonAbility struct {
	Is_hidden bool
	Slot      int
	Ability   namedAPIResource
}

type pokemonType struct {
	Slot int
	Type namedAPIResource
}

type pokemonTypePast struct {
	Generation namedAPIResource
	Types      []pokemonType
}

type pokemonHeldItem struct {
	Item           namedAPIResource
	VersionDetails []pokemonHeldItemVersion
}

type pokemonHeldItemVersion struct {
	Version namedAPIResource
	Rarity  int
}

type pokemonMove struct {
	Move                  namedAPIResource
	Version_group_details []pokemonMoveVersion
}

type pokemonMoveVersion struct {
	Move_learn_method namedAPIResource
	Version_group     namedAPIResource
	Level_learned_at  int
}

type pokemonStat struct {
	Stat      namedAPIResource
	Effort    int
	Base_stat int
}

type pokemonSprites struct {
	Front_default      string
	Front_shiny        string
	Front_female       string
	Front_shiny_female string
	Back_default       string
	Back_shiny         string
	Back_female        string
	Back_shiny_female  string
}

type pokemonCries struct {
	Latest string
	Legacy string
}
