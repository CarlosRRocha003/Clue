package player

type Player struct {
	Id   int
	Name string
}

var noone = Player{Id: 0, Name: "No one"}
var someone = Player{Id: 1, Name: "Someone"}
var players = []Player{noone, someone}

func GetAllPlayers() []Player {
	return players
}

func fillPlayers(filledPlayers []Player) []Player {
	players = filledPlayers
	return players
}

func AddPlayer(player Player) {
	if player.Id == 0 {
		player.Id = len(players)
	}
	players = append(players, player)
}

func removeAllPlayers() {
	players = []Player{}
}
