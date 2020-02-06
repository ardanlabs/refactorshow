package handlers

import (
	"log"
	"net/http"
)

// NewMux constructs and mux with all route predefined.
func NewMux(log *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	// Serve everything in the css folder, the img folder and mp3 folder as a file.
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	mux.Handle("/mp3/", http.StripPrefix("/mp3/", http.FileServer(http.Dir("mp3"))))

	base := Base{log}
	mux.HandleFunc("/", base.Home)
	mux.HandleFunc("/scale", base.Scale)
	mux.HandleFunc("/scaleshow", base.ScaleShow)
	mux.HandleFunc("/duets", base.Duet)
	mux.HandleFunc("/duetshow", base.DuetShow)

	return mux
}
