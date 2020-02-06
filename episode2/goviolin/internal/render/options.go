package render

import (
	"strings"
)

// Option represents the options for generating content.
type Option struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

// PageVars represents the input for generating a web page.
type PageVars struct {
	Title         string
	Scalearp      string
	Key           string
	Pitch         string
	DuetImgPath   string
	ScaleImgPath  string
	GifPath       string
	AudioPath     string
	AudioPath2    string
	DuetAudioBoth string
	DuetAudio1    string
	DuetAudio2    string
	LeftLabel     string
	RightLabel    string
	Scales        []Option
	Duets         []Option
	Pitches       []Option
	Keys          []Option
	Octaves       []Option
}

// SetDefaultOptions provide the defaults for rendering scales.
func SetDefaultOptions() ([]Option, []Option, []Option, []Option) {
	scale := []Option{
		Option{"Scale", "Scale", false, true, "Scales"},
		Option{"Scale", "Arpeggio", false, false, "Arpeggios"},
	}

	pitch := []Option{
		Option{"Pitch", "Major", false, true, "Major"},
		Option{"Pitch", "Minor", false, false, "Minor"},
	}

	key := []Option{
		Option{"Key", "A", false, true, "A"},
		Option{"Key", "Bb", false, false, "Bb"},
		Option{"Key", "B", false, false, "B"},
		Option{"Key", "C", false, false, "C"},
		Option{"Key", "C#/Db", false, false, "C#/Db"},
		Option{"Key", "D", false, false, "D"},
		Option{"Key", "Eb", false, false, "Eb"},
		Option{"Key", "E", false, false, "E"},
		Option{"Key", "F", false, false, "F"},
		Option{"Key", "F#/Gb", false, false, "F#/Gb"},
		Option{"Key", "G", false, false, "G"},
		Option{"Key", "G#/Ab", false, false, "G#/Ab"},
	}

	octive := []Option{
		Option{"Octave", "1", false, true, "1 Octave"},
		Option{"Octave", "2", false, false, "2 Octave"},
	}

	return scale, pitch, key, octive
}

// SetOctaveOptions sets the octave options based on the specified octave.
func SetOctaveOptions(octave string) []Option {
	var options []Option

	switch octave {
	case "1":
		options = []Option{
			Option{"Octave", "1", false, true, "1 Octave"},
			Option{"Octave", "2", false, false, "2 Octave"},
		}
	case "2":
		options = []Option{
			Option{"Octave", "1", false, false, "1 Octave"},
			Option{"Octave", "2", false, true, "2 Octave"},
		}
	}

	return options
}

// SetPitchOptions sets the pitch options based on the specified pitch.
func SetPitchOptions(pitch string) []Option {
	var options []Option

	switch pitch {
	case "Major":
		options = []Option{
			Option{"Pitch", "Major", false, true, "Major"},
			Option{"Pitch", "Minor", false, false, "Minor"},
		}
	case "Minor":
		options = []Option{
			Option{"Pitch", "Major", false, false, "Major"},
			Option{"Pitch", "Minor", false, true, "Minor"},
		}
	}

	return options
}

// SetScaleOptions sets the scale options based on the specified scale.
func SetScaleOptions(scale string) []Option {
	var options []Option

	switch scale {
	case "Scale":
		options = []Option{
			Option{Name: "Scale", Value: "Scale", IsDisabled: false, IsChecked: true, Text: "Scales"},
			Option{Name: "Scale", Value: "Arpeggio", IsDisabled: false, IsChecked: false, Text: "Arpeggios"},
		}

	case "Arpeggio":
		options = []Option{
			Option{Name: "Scale", Value: "Scale", IsDisabled: false, IsChecked: false, Text: "Scales"},
			Option{Name: "Scale", Value: "Arpeggio", IsDisabled: false, IsChecked: true, Text: "Arpeggios"},
		}
	}

	return options
}

// SetKeyOptions sets the key options based on specified key.
func SetKeyOptions(key string) []Option {
	var options []Option

	switch key {
	case "A":
		options = []Option{
			Option{"Key", "A", false, true, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Bb":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, true, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "B":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, true, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, true, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C#/Db":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, true, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "D":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, true, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Eb":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, true, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "E":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, true, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, true, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F#/Gb":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, true, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "G":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, true, "G"},
			Option{"Key", "G#/Ab", false, false, "G#/Ab"},
		}

	case "G#/Ab":
		options = []Option{
			Option{"Key", "A", false, false, "A"},
			Option{"Key", "Bb", false, false, "Bb"},
			Option{"Key", "B", false, false, "B"},
			Option{"Key", "C", false, false, "C"},
			Option{"Key", "C#/Db", false, false, "C#/Db"},
			Option{"Key", "D", false, false, "D"},
			Option{"Key", "Eb", false, false, "Eb"},
			Option{"Key", "E", false, false, "E"},
			Option{"Key", "F", false, false, "F"},
			Option{"Key", "F#/Gb", false, false, "F#/Gb"},
			Option{"Key", "G", false, false, "G"},
			Option{"Key", "G#/Ab", false, true, "G#/Ab"},
		}
	}

	return options
}

// SetActualKey transforms the provided key based on the selected pitch.
// For major scales if the key is longer than 2 characters, we only care about
// the last 2 characters.
// For minor scales if the key is longer than 2 characters, we only care about
// the first 2 characters.
func SetActualKey(pitch string, key string) string {
	switch pitch {
	case "Major":

		// Only select last two characters for keys which contain two possible
		// names e.g. C#/Db.
		if len(key) > 2 {
			key = key[3:]
		}

	case "Minor":

		// Only select first two characters for keys which contain two possible
		// names e.g. C#/Db.
		if len(key) > 2 {
			key = key[:2]
		}
	}

	return key
}

// SetMusicLabels sets the text for the music players. Major have a scale and
// a drone, while minor have melodic and harmonic minor scales.
func SetMusicLabels(pitch string, scale string) (string, string) {
	var left string
	var right string

	switch pitch {
	case "Major":
		left = "Listen to Major "
		right = "Listen to Drone"
		if scale == "Scale" {
			left += "Scale"
			break
		}
		left += "Arpeggio"

	case "Minor":
		switch scale {
		case "Scale":
			left = "Listen to Harmonic Minor Scale"
			right = "Listen to Melodic Minor Scale"
		case "Arpeggio":
			left = "Listen to Minor Arpeggio"
			right = "Listen to Drone"
		}
	}

	return left, right
}

// SetAssetPaths builds paths to img and mp3 files that correspond to user
// selection.
func SetAssetPaths(pitch string, scale string, key string, octave string) (string, string, string) {
	imgPath, audioPath, audioPath2 := "img/", "mp3/", "mp3/"

	switch scale {
	case "Scale":
		imgPath += "scale/"
		audioPath += "scale/"
	case "Arpeggio":
		imgPath += "arps/"
		audioPath += "arps/"
	}

	switch pitch {
	case "Major":
		imgPath += "major/"
		audioPath += "major/"
	case "Minor":
		imgPath += "minor/"
		audioPath += "minor/"
	}

	audioPath += strings.ToLower(key)
	imgPath += strings.ToLower(key)

	imgPath = changeSharpToS(imgPath)
	audioPath = changeSharpToS(audioPath)

	switch octave {
	case "1":
		imgPath += "1"
		audioPath += "1"
	case "2":
		imgPath += "2"
		audioPath += "2"
	}

	audioPath += ".mp3"
	imgPath += ".png"

	// Audio path2 can either be a melodic minor scale or a drone note.
	// Set to melodic minor scale - if the first 16 characters of audio path.
	switch {
	case audioPath[:16] == "mp3/scale/minor/":

		// Set audioPath2 to the original audioPath.
		audioPath2 = audioPath

		// Chop off the last 4 characters, this removes .mp3 extension.
		audioPath2 = audioPath2[:len(audioPath2)-4]

		// Then add m for melodic and the .mp3 suffix.
		audioPath2 += "m.mp3"

	default:
		audioPath2 += "drone/"
		audioPath2 += strings.ToLower(key)

		// May have just added a # to the path, so use the function to
		// change # to s.
		audioPath2 = changeSharpToS(audioPath2)

		switch octave {
		case "1":
			audioPath2 += "1.mp3"
		case "2":
			audioPath2 += "2.mp3"
		}
	}

	return imgPath, audioPath, audioPath2
}

// changeSharpToS WE DON'T KNOW WHY YET.
func changeSharpToS(path string) string {
	if strings.Contains(path, "#") {
		path = path[:len(path)-1]
		path += "s"
	}
	return path
}
