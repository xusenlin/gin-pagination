package gin_pagination

import "gorm.io/gorm"

type Config struct {
	PageSizeMaxVal     int
	PageSizeDefaultVal string
	DB                 *gorm.DB
}

var options *Config

func Init(c *Config) {
	options = c
}
