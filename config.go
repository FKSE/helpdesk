package helpdesk

import "fmt"

type ConfigDatabase struct {
	Driver   string
	Host     string
	Database string
	User     string
	Password string
}

type Configuration struct {
	Port     int
	Database ConfigDatabase
}

func (c *ConfigDatabase) ConnectionString() (s string) {
	switch c.Driver {
	case "mysql":
		// Default
		s = fmt.Sprintf("%s:%s@/%s?charset=utf8", c.User, c.Password, c.Database)
	}

	return s
}
