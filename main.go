package main

import (
	"github.com/vishalvivekm/pyqqserver/app"
)

  
func main() {
	application := app.App{}
	application.Init()
	application.Run()
}
