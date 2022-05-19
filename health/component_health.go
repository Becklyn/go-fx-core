package health

type componentHealth struct {
	healthy bool
	reason  string
}

type componentChanged struct {
	name   string
	health componentHealth
}
