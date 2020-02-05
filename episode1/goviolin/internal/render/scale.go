package render

import "strings"

// SetDefaultScaleOptions provide the defaults for rendering scales.
func SetDefaultScaleOptions() ([]ScaleOptions, []ScaleOptions, []ScaleOptions, []ScaleOptions) {

	// Set the default scaleOptions for scales and arpeggios.
	sOptions := []ScaleOptions{
		ScaleOptions{"Scalearp", "Scale", false, true, "Scales"},
		ScaleOptions{"Scalearp", "Arpeggio", false, false, "Arpeggios"},
	}

	// Set the default PitchOptions for scales and arpeggios.
	pOptions := []ScaleOptions{
		ScaleOptions{"Pitch", "Major", false, true, "Major"},
		ScaleOptions{"Pitch", "Minor", false, false, "Minor"},
	}

	// Set the default KeyOptions for scales and arpeggios.
	kOptions := []ScaleOptions{
		ScaleOptions{"Key", "A", false, true, "A"},
		ScaleOptions{"Key", "Bb", false, false, "Bb"},
		ScaleOptions{"Key", "B", false, false, "B"},
		ScaleOptions{"Key", "C", false, false, "C"},
		ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
		ScaleOptions{"Key", "D", false, false, "D"},
		ScaleOptions{"Key", "Eb", false, false, "Eb"},
		ScaleOptions{"Key", "E", false, false, "E"},
		ScaleOptions{"Key", "F", false, false, "F"},
		ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
		ScaleOptions{"Key", "G", false, false, "G"},
		ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
	}

	// Set the default OctaveOptions for scales and arpeggios.
	oOptions := []ScaleOptions{
		ScaleOptions{"Octave", "1", false, true, "1 Octave"},
		ScaleOptions{"Octave", "2", false, false, "2 Octave"},
	}

	return sOptions, pOptions, kOptions, oOptions
}

// SetKeyOptions sets the key options based on specified key.
func SetKeyOptions(key string) (kOptions []ScaleOptions) {
	switch key {
	case "A":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, true, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Bb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, true, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "B":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, true, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, true, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C#/Db":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, true, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "D":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, true, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Eb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, true, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "E":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, true, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, true, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F#/Gb":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, true, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "G":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, true, "G"},
			ScaleOptions{"Key", "G#/Ab", false, false, "G#/Ab"},
		}

	case "G#/Ab":
		kOptions = []ScaleOptions{
			ScaleOptions{"Key", "A", false, false, "A"},
			ScaleOptions{"Key", "Bb", false, false, "Bb"},
			ScaleOptions{"Key", "B", false, false, "B"},
			ScaleOptions{"Key", "C", false, false, "C"},
			ScaleOptions{"Key", "C#/Db", false, false, "C#/Db"},
			ScaleOptions{"Key", "D", false, false, "D"},
			ScaleOptions{"Key", "Eb", false, false, "Eb"},
			ScaleOptions{"Key", "E", false, false, "E"},
			ScaleOptions{"Key", "F", false, false, "F"},
			ScaleOptions{"Key", "F#/Gb", false, false, "F#/Gb"},
			ScaleOptions{"Key", "G", false, false, "G"},
			ScaleOptions{"Key", "G#/Ab", false, true, "G#/Ab"},
		}
	}

	return kOptions
}

// ChangeSharpToS WE DON'T KNOW WHY YET.
func ChangeSharpToS(path string) string {
	if strings.Contains(path, "#") {
		path = path[:len(path)-1]
		path += "s"
	}
	return path
}
