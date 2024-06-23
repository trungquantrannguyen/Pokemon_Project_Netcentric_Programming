# NET-CENTRIC REPORT

## POKEMON GAME

[Project Youtube link](https://youtu.be/jikYgbk13Io)

## I. Members and contributions

| Name                   | Student ID  | Contribution           |
| ---------------------- | ----------- | ---------------------- |
| Bùi Thị Cẩm Vân        | ITITIU20111 | 40%                    |
|                        |             | • Models, handlers     |
|                        |             | • Pokedex              |
|                        |             | • README.file          |
|                        |             | • Presentation content |
| Trần Nguyễn Trung Quân | ITITDK21071 | 60%                    |
|                        |             | • Models               |
|                        |             | • Battle               |
|                        |             | • Catch                |
|                        |             | • Presentation speaker |

## II. Acknowledgment

We would like to express our deepest gratitude to Msc. Le Thanh Son and Msc. Nguyen Trung Nghia for their invaluable guidance and support throughout this course.

## III. Project structure

### 1. How to run

- **First, download packages:**
  - `cd project_net_centric >> go run main.go`
- **Pokedex feature:**
  - `cd project_net_centric/pokedexServer >> go run main.go`
  - `cd project_net_centric/pokedexClient >> go run main.go`
- **Battle feature:**
  - `cd project_net_centric/battleServer >> go run main.go`
  - `cd project_net_centric/battleClient >> go run main.go`
- **Catch feature:**
  - `cd project_net_centric/catch >> go run main.go`

### 2. Modules:

- **Handlers**
  - Handlers manage incoming requests and coordinate with repositories to fetch or manipulate data. They are responsible for the business logic.
- **Repositories**
  - Repositories handle data access and persistence. They provide methods to interact with the data source, whether it's a file system, a database, or an API.
- **Models**
  - Models define the structure of the data used in the application. They include various JSON schemas and other data representations.
- **Utilities**
  - Utility functions and helpers provide common functionality used across different application parts.

### 3. Structure:

project_net_centric/
├── .idea
├── battleClient
│ └── main.go
├── battleServer
│ ├── usermanager
│ │ ├── usermanager.go
│ │ └── pokemonNames.json
│ └── main.go
├── catch
│ ├── main.go
│ └── pokemon_image.png
├── internal
│ ├── handlers
│ │ └── pokedex_handler.go
│ ├── models
│ └── repositories
├── pokedexClient
│ └── client.go
├── pokedexServer
│ └── server.go
├── utils
│ └── poke_map.go
├── .gitignore
├── go.mod
├── main.go
└── README.me

- **battleClient/**: Contains the client-side code and resources for handling battle-related functionalities.
  - `main.go`: The main client program for interacting with the battle server.
- **battleServer/**: Houses the server-side code and logic for managing battles.
  - **usermanager/**: Contains components related to user management within the battle server.
    - `usermanager.go`: Handles user management functions.
    - `pokemonNames.json`: JSON file containing Pokémon names and their corresponding IDs.
  - `main.go`: The main server program handling battle logic and player interactions.
- **catch/**: Contains code for catching Pokémon.
  - `main.go`: The main program for catching Pokémon.
  - `pokemon_image.png`: Image resource for the catch functionality.
- **internal/**: Contains core internal components shared across different parts of the application.
- **handlers/**: Holds request handlers or controllers that manage incoming requests and responses.
  - `pokedex_handler.go`: Handles requests related to Pokémon data.
- **models/**: Includes data models and schemas used throughout the application. All the data is stored in this folder.
  - `monster_moves`: Moves corresponding to Pokémon.
  - `monsters`: The final Pokémon information.
  - `moves`: The details of moves.
  - `skim_monsters`: The general information of Pokémon.
- **repositories/**: Contains data access layers and repository patterns for interacting with databases or file systems.
- **pokedexClient/**: Includes the client-side code and assets for interacting with the Pokédex.
  - `client.go`: The main client program for querying the Pokédex.
- **pokedexServer/**: Comprises the server-side code and logic for managing the Pokédex.
  - `server.go`: The main server program handling Pokédex data.
- **utils/**: Stores utility functions and helper modules used across the application for various purposes.
  - `poke_map.go`: Utility functions for mapping Pokémon names to IDs.
- `.gitignore`: Specifies files and directories to be ignored by Git.
- `go.mod`: Go module file for dependency management.
- `main.go`: Entry point of the project.
- `README.me`: Documentation file for the project.

## IV. Key functions

### 1. Models

Main models move, monsters, skim_monsters, monster_moves

### 2. API

Have type and description data get from the Pokemon API

### 3. POKEDEX

- **Handler GetPokemon**: get Pokémon data and marshalize it into JSON file
- **Repository GetMonsterById**: get information of Pokémon monster by its ID

### 4. BATTLE

- **Usermanager**: handle function to manage user information
  - `NewUserManager`: create a new instance of User Manager
  - `AddUser`: add user
  - `UpdatePokemonData`: update Pokemon data when user input pokemon from terminal (client side)
  - `StartBattle`: start battle when both users have been added to the server, and both inputed enough 3 Pokemons in terminal
  - `AllPokemonProvided`: check if there are exactly 2 connected players and enough Pokemon chosen by users

### 5. CATCH

- **get**: get data by pass in URL parameter from http
- **getRegionDetails, getLocationDetails**: get data about region and location of the pokemon in the map by passing in pokemon id
- **getPokemon**: get pokemon data based on the pokemon name
- **getRandomEncounter**: choose a random pokemon from the selected location and get data of pokemon by `getPokemon()`
- **downloadImage, openImage**: download image based on URL and open that Pokemon image.
