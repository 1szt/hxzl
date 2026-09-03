package static

import (
	web "mcup-server/gin"
)

func Run() {
	web.Router.Static("/static", "./static")
}
