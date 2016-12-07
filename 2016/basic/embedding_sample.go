package main

type Sound struct {
	Filename string
}

func (s Sound) Play() {}

type Music struct {
	Sound
	Artist string
	Genre  string
}

func main() {
	m := Music{
		Sound:  Sound{Filename: "muse.mp3"},
		Artist: "Muse",
		Genre:  "Rock",
	}
	m.Play()
}
