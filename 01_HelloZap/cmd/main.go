package main

import "GoZapLearn/01_HelloZap/routes"

func main() {
	routes.RunRouter(routes.NewRouter())
}
