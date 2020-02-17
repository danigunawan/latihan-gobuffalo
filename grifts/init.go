package grifts

import (
	"github.com/Danigunawan/latihan/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
