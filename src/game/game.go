package game

import (
	"./actors"
	"./field"
	"./views"
	"github.com/dradtke/gopher"
	"github.com/dradtke/gopher/config"
)

const (
	TILE_SIZE = 70
)

func Play() {
	gopher.NewState(&PlayingState{}, new(views.HumanView))
}

type PlayingState struct {
	gopher.BaseState
}

func (s *PlayingState) InitState() {
	gopher.AddActor(actors.NewBox(0, 4, actors.NormalBox))
	gopher.AddActor(actors.NewBox(1, 4, actors.ExplosiveBox))
	gopher.AddActor(actors.NewBox(2, 4, actors.AltBox))
	gopher.AddActor(actors.NewBox(3, 4, actors.NormalBox))
	gopher.AddActor(actors.NewBox(4, 4, actors.EmptyBox))
	gopher.AddActor(actors.NewBox(5, 4, actors.NormalBox))
	gopher.AddActor(actors.NewBox(6, 4, actors.EmptyBox))
	gopher.AddActor(actors.NewBox(7, 4, actors.NormalBox))
	gopher.AddActor(actors.NewBox(8, 4, actors.WarningBox))
	gopher.AddActor(actors.NewBox(9, 4, actors.NormalBox))

	_, sh := config.DisplaySize()
	field.Player = new(actors.Player)
	field.Player.X = 100
	field.Player.Y = float32(sh - (TILE_SIZE * 2.5))
	gopher.AddActor(field.Player)
}