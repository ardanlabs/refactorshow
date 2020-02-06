package render

import (
	"fmt"
	"net/http"
	"text/template"
)

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
