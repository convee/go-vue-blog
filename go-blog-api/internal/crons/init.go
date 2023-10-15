package crons

import (
	"log"

	"github.com/convee/go-blog-api/configs"
	"github.com/robfig/cron"
)

func Init() {
	// second     = field(fields[0], seconds)
	// minute     = field(fields[1], minutes)
	// hour       = field(fields[2], hours)
	// dayofmonth = field(fields[3], dom)
	// month      = field(fields[4], months)
	// dayofweek  = field(fields[5], dow)
	c := cron.New()
	if configs.Conf.Cron.Push {

	}

	c.Start()
	log.Println("cron work starting...")

}
