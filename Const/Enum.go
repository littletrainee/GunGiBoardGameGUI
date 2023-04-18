package Const

type Color int
type Level int

const (
	White Color = iota
	Black
)

const (
	Beginner Level = iota
	Elementary
	Intermediate
	Advanced
)

const ALLKOMA string = "帥大中小侍忍槍砦馬兵弓砲筒謀"

const MAXROW int = 10
