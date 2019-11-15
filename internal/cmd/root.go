package cmd

import "savenotes/internal/api"

// Start prepare db and initlise router
func Start() {
	db := getPostgresdb()
	app := api.App{
		Router : 
	}
}
