package bridge

type Color interface {
	Hex() string
}

type Red struct{}

func (r Red) Hex() string {
	return "#E22E50"
}

type Green struct{}

func (r Green) Hex() string {
	return "#570658"
}
