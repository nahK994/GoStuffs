package composition

import (
	"fmt"
)

type Caller struct {
	number string
}

func (c Caller) MakeCall(number string) {
	fmt.Println("Calling from", c.number, "to", number)
}

func (c Caller) PrintCallerNumber() {
	fmt.Println("The caller number is", c.number)
}

type Browser struct{}

func (b Browser) BrowseInternet() {
	fmt.Println("Browsing the internet")
}

type Camera struct{}

func (c Camera) TakePhoto() {
	fmt.Println("Taking a photo")
}

type MediaPlayer struct{}

func (m MediaPlayer) WatchMovie() {
	fmt.Println("Watching a movie")
}

type Smartphone struct {
	Caller
	Browser
	Camera
}

type Tablet struct {
	Caller
	Browser
	MediaPlayer
}

func Composition() {
	phone := Smartphone{
		Caller:  Caller{number: "01676498001"},
		Browser: Browser{},
		Camera:  Camera{},
	}

	tablet := Tablet{
		Caller:      Caller{number: "01176498001"},
		Browser:     Browser{},
		MediaPlayer: MediaPlayer{},
	}

	phone.PrintCallerNumber()
	phone.MakeCall("123-456-7890")
	phone.BrowseInternet()
	phone.TakePhoto()
	fmt.Println()
	tablet.PrintCallerNumber()
	tablet.MakeCall("098-765-4321")
	tablet.BrowseInternet()
	tablet.WatchMovie()
}
