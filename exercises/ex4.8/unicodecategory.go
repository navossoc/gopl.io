package main

type UnicodeCategory int

//go:generate stringer -type=UnicodeCategory

const (
	control UnicodeCategory = iota
	digit
	graphic
	letter
	lower
	mark
	number
	print
	punct
	space
	symbol
	title
	upper
	invalid
)
