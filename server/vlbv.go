package main

import (
	"log"

	. "github.com/lxn/walk/declarative"
)

func main() {
	MainWindow{
		Title:   "Title",
		MinSize: Size{480, 320},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "Button",
				OnClicked: func() {
					log.Println("clicked")
				},
			},
			ListBox{},
			ProgressBar{},
		},
	}.Run()
}
