package action

type Action int

const (
	NONE               Action = iota // 無
	MOVE                             // 直接移動
	ILLEGAL_MOVE                     // 不可移動
	CAPTURE_ONLY                     // 僅可俘獲
	CAPTURE_OR_CONTROL               // 俘獲或控制
)
