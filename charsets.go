package main


type charsets struct {
	option string
	min int
	max int
}

func (c charsets) getRandomChar () string {
	number := c.min + r.Intn(c.max +1 -c.min)
	return string(number)
}

