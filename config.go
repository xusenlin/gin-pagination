package gin_pagination

import "gorm.io/gorm"

type Config struct {
	pageSizeDefaultVal string
	db                 *gorm.DB
}

var options *Config

func Init(c *Config) {
	options = c
}
