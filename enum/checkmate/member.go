package checkmate

type CheckMate int

const (
	NONE      = iota // 尚未出現勝負
	CHECKMATE        // 將軍
	BEEN_CHECKMATE
)
