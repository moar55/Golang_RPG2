package scan

import "Golang_RPG/models"

type Battle struct {
	Message string         `json:"message"`
	Type    string         `json:"Type"`
	Enemy   models.Enemies `json:"Enemy"`
}

func EnterBattle(message string, _type string, enemy models.Enemies) *Battle {
	return &Battle{Message: message + enemy.Name, Type: _type, Enemy: enemy}
}
