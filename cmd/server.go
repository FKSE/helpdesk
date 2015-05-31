package main

import "github.com/fkse/helpdesk"

func main() {
	// Make config
	config := &helpdesk.Configuration{
		Port: 3000,
		Database: helpdesk.ConfigDatabase{
			Driver:   "mysql",
			Host:     "localhost",
			Database: "helpdesk",
			User:     "root",
		},
	}
	helpdesk.Run(config)
}
