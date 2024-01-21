package requests

type AddUserMatchRequest struct {
	Name   string    `json:"name"  binding:"required"`
	Height uint64    `json:"height" binding:"required"`
	Gender *uint32   `json:"gender" binding:"required"`
	Times  uint64    `json:"times"binding:"required"`
	Rule   MatchRule `json:"rule" binding:"required"`
}

type SingleUserRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type MatchRule struct {
	Range struct {
		Start uint64 `json:"start" binding:"required"`
		End   uint64 `json:"end" binding:"required"`
	} `json:"range"binding:"required"`
	Gender uint32 `json:"gender"binding:"required"`
}
