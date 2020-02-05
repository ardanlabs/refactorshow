package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// ScaleOptions represents the options for generating content.
type ScaleOptions struct {
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
	ScaleOptions  []ScaleOptions
	DuetOptions   []ScaleOptions
	PitchOptions  []ScaleOptions
	KeyOptions    []ScaleOptions
	OctaveOptions []ScaleOptions
}

// Render generates the html for any given web page.
func Render(w http.ResponseWriter, tmpl string, pageVars PageVars) error {

	// Prefix the name passed in with templates.
	tmpl = fmt.Sprintf("templates/%s", tmpl)

	// Parse the template file held in the templates folder.
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	// Execute the template and pass in the variables to fill the gaps.
	if err := t.Execute(w, pageVars); err != nil {
		return err
	}

	return nil
}
