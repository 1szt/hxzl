package template

import (
	web "mcup-server/gin"
)

func Run() {
	web.Router.LoadHTMLGlob("template/*")
}
