package gin

func Static() {
	Router.Static("/static", "./static")
}
