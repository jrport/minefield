package pool

type MatchPool struct {
	maxSize int
}

func NewMatchPool(ms int) *MatchPool{
	return &MatchPool{
		maxSize: ms,
	}
}
