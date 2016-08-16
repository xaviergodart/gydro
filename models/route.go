package models

type Route struct {
	id       int
	Pattern  string
	Backends []string
}
