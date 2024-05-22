package main

import (
	"net/http"
)

func (app *application) healthCheckerHandler(w http.ResponseWriter, r *http.Request) {

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, 200, env, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server ecountered a problem and could not process it", http.StatusInternalServerError)
	}

}
