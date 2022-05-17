package crontab

import "github.com/vela-security/vela-public/lua"

func (c *Cron) Header(out lua.Console) {
	out.Printf("type: %s", c.Type())
	out.Printf("uptime: %s", c.uptime.Format("2006-01-02 15:04:06"))
	out.Println("version:  v1.0.0")
	out.Println("")
}

func (c *Cron) Show(out lua.Console) {
	c.Header(out)
	for id, val := range c.masks {
		out.Printf("%-4d%-15s %s", id, val.spec, val.label)
	}
}

func (c *Cron) Help(out lua.Console) {
}
