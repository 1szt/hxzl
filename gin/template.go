package gin

func Template() {
	Router.LoadHTMLGlob("template/*.html")
}
