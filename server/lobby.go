package main

import (
	"github.com/jak103/uno/db"
)

type Game struct {
	Name        string
	NumPlayers  uint32
	MaxPlayers  uint32
	HasStarted  bool
	HasPassword bool
}

// TODO: This function will need to query the DB for the server list.
// For now, just return a dummy list of what we think we will need
func getGames() []Game {
	games := []Game{}

	games = append(games, Game{"Game1", 1, 10, false, false})
	games = append(games, Game{"Game2", 10, 10, true, true})

	return games
}

func createGameFromLobby(gameName string, gamePassword string) (map[string]interface{}, error) {
	database, err := db.GetDb()

	if err != nil {
		return nil, err
	}

	// TODO: We need some way to specify the game parameters
	// This will require changing the UnoDB interface
	game, err := database.CreateGame()

	if err != nil {
		return nil, err
	}

	jsonRsp := map[string]interface{}{
		// TODO: Games do not have names, but IDs. We should add names to the model
		"name": game.ID,
		"id":   game.ID,
	}

	return jsonRsp, nil
}
