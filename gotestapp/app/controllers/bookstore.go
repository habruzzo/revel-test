package controllers

import (
	"github.com/revel/revel"
)

type BookStore struct {
	*revel.Controller
}

func (c BookStore) Index() revel.Result {
	return c.Render()
}