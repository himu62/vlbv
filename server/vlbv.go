package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	walk.MainWindow{
		Title:   "Title",
		MinSize: walk.Size{480, 320},
		Layout:  walk.VBox{},
		children: []walk.Widget{
			walk.PushButton{
				Text: "Button",
				OnClicked: func() {
					log.Println("clicked")
				},
			},
		},
	}.Run()
}
