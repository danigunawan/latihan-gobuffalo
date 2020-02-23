package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

func DaniHandler(c buffalo.Context) error {
	c.Set("name", "Dani Handler")
	return c.Render(200, r.String("Hi <%= name %>"))
}
