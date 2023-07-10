package models

type Board struct {
	boardContent [9][9]int
	valid        bool
	solved       bool
	difficulty   string
}
