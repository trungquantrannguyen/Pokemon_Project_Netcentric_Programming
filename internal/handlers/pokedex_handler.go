package handlers

import (
	"encoding/json"

	"github.com/trungquantrannguyen/project_net_centric/internal/repositories"
	"github.com/trungquantrannguyen/project_net_centric/utils"
)

type PokeDexHandler struct {
	PokedexReposiotry *repositories.PokedexRepository
}

func NewPokeDexHandler(pokedexReposiotry *repositories.PokedexRepository) *PokeDexHandler {
	return &PokeDexHandler{
		PokedexReposiotry: pokedexReposiotry,
	}
}

func (s *PokeDexHandler) GetPokemon(name string) ([]byte, error) {
    // Retrieves the ID of the Pokémon from the utils.PokeMap using the provided name
    data, err := s.PokedexReposiotry.GetMonsterByID(utils.PokeMap[name])
    if err != nil {
        // If there is an error fetching the Pokémon data, return an empty byte slice and the error
        return []byte{}, err
    }

    // Marshal the Pokémon data into a JSON-formatted byte slice with indentation for readability
    jsonData, err := json.MarshalIndent(data, "", "    ")
    if err != nil {
        // If there is an error during marshaling, return an empty byte slice and the error
        return []byte{}, err
    }

    // Return the JSON-formatted Pokémon data and no error
    return jsonData, nil
}
