package views

import "github.com/gin-contrib/multitemplate"

func LoadTemplates() multitemplate.Render {
	r := multitemplate.New()

	r.AddFromFiles("welcome",
		"templates/master.html",
		"templates/welcome.html",
	)
	r.AddFromFiles("about",
		"templates/master.html",
		"templates/about.html",
	)

	return r
}