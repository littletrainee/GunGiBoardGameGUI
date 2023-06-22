// 對弈期間的狀態列舉
package state

type State int

const (
	PREPARING State = iota
	ARRANGEMENT
	DUELING
)
