package controllers

import (
	"github.com/revel/revel"
)

type TicTacToe struct {
	*revel.Controller
}

func (c TicTacToe) Index() revel.Result {
	return c.Render()
}