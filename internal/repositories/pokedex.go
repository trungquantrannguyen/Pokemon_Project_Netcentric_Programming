package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/trungquantrannguyen/project_net_centric/internal/models"
	"github.com/trungquantrannguyen/project_net_centric/utils"
)

type PokedexRepository struct {
	BasePath string
}

func NewPokedexRepository(baseFilePath string) *PokedexRepository {
	return &PokedexRepository{
		BasePath: baseFilePath,
	}
}

type Pokemon struct {
	Monster             *models.Monster             `json:"monster"`
	Description         []*models.Descriptions      `json:"description"`
	Types               []*models.Types             `json:"types"`
	MonsterMoves        []*models.Move              `json:"monster_moves"`
}

func (p *PokedexRepository) GetMonsterMovesByID(id string) ([]*models.Move, error) {
	var data []*models.Move
	pathFile := fmt.Sprintf("%s/monster_moves/data/%s.json", p.BasePath, id)
	file, err := ioutil.ReadFile(pathFile)
	if err != nil {
		return nil, err
	}
	var monsterMove models.MonsterMove
	err = json.Unmarshal(file, &monsterMove)
	if err != nil {
		return nil, err
	}

	requestMoves := monsterMove.Move
	for _, move := range requestMoves {
		pathFile = fmt.Sprintf("%s/moves/data/%d.json", p.BasePath, move.Id)
		file, err = ioutil.ReadFile(pathFile)
		if err != nil {
			return nil, err
		}
		var m models.Move
		err = json.Unmarshal(file, &m)
		if err != nil {
			return nil, err
		}
		data = append(data, &m)
	}
	return data, nil

}

func (p *PokedexRepository) GetMonsterTypeByID(path []models.ListMapObject) ([]*models.Types, error) {
	var data []*models.Types
	for _, id := range path {
		pathFile := fmt.Sprintf("%s/api/v1/type/%s/poke.json", p.BasePath, id.Name)
		file, err := ioutil.ReadFile(pathFile)
		if err != nil {
			return nil, err

		}
		var t models.Types
		err = json.Unmarshal(file, &t)
		if err != nil {
			return nil, err

		}
		data = append(data, &t)
	}
	return data, nil
}

func (p *PokedexRepository) GetMonsterDescription(path []models.ListMapObject) ([]*models.Descriptions, error) {
	var data []*models.Descriptions
	for _, id := range path {
		pathFile := fmt.Sprintf("%s%spoke.json", p.BasePath, id.ResourceURI)
		file, err := ioutil.ReadFile(pathFile)
		if err != nil {
			return nil, err
		}
		var desc models.Descriptions
		err = json.Unmarshal(file, &desc)
		if err != nil {
			return nil, err
		}
		data = append(data, &desc)
	}

	return data, nil

}

func (p *PokedexRepository) GetMonsterByID(id string) (*Pokemon, error) {
	// Construct the path to the skim monster data file
	pathFile := fmt.Sprintf("%s/skim_monsters/data/%s.json", p.BasePath, id)
	
	// Read the skim monster data file
	file, err := ioutil.ReadFile(pathFile)
	if err != nil {
		// Return nil and the error if reading the file fails
		return nil, err
	}
	
	// Unmarshal the JSON data into a SkimMonster struct
	var monster models.SkimMonster
	err = json.Unmarshal(file, &monster)
	if err != nil {
		// Return nil and the error if unmarshaling fails
		return nil, err
	}

	// Construct the path to the evolution data file
	pathFile = fmt.Sprintf("%s/evolutions/data/%s.json", p.BasePath, id)
	
	// Read the evolution data file
	file, err = ioutil.ReadFile(pathFile)
	if err != nil {
		// Return nil and the error if reading the file fails
		return nil, err
	}
	
	// Construct the path to the monster supplemental data file
	pathFile = fmt.Sprintf("%s/monster_supplementals/data/%s.json", p.BasePath, id)
	
	// Read the monster supplemental data file
	file, err = ioutil.ReadFile(pathFile)
	if err != nil {
		// Return nil and the error if reading the file fails
		return nil, err
	}
	
	// Retrieve the monster's moves by its ID
	monsterMoves, err := p.GetMonsterMovesByID(id)
	if err != nil {
		// Return nil and the error if retrieving the moves fails
		return nil, err
	}

	// Retrieve the monster's descriptions
	desc, _ := p.GetMonsterDescription(monster.Description)

	// Retrieve the monster's types by its type IDs
	types, err := p.GetMonsterTypeByID(monster.Type)
	if err != nil {
		// Return nil and the error if retrieving the types fails
		return nil, err
	}

	// Return the constructed Pokemon struct
	return &Pokemon{
		Monster:             monster.ToMonster(),            // Convert SkimMonster to Monster struct
		MonsterMoves:        monsterMoves,                   // Set monster moves
		Description:         desc,                           // Set descriptions
		Types:               types,                          // Set types
	}, nil
}


func crawl() string {

	var pokedex PokedexRepository

	data, err := pokedex.GetMonsterByID(utils.PokeMap["Ivysaur"])
	if err != nil {
		return ""
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Print JSON data
	return string(jsonData)

}
