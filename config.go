package gin_pagination

import "gorm.io/gorm"

type Config struct {
	PageSizeMaxVal     int
	PageSizeDefaultVal int
	DB                 *gorm.DB
}

var options *Config

func Init(c *Config) {
	if c.PageSizeDefaultVal > c.PageSizeMaxVal {
		c.PageSizeDefaultVal = c.PageSizeMaxVal
	}
	options = c
}
