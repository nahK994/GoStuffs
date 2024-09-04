package engine

type Engine struct {
	horsepower int
}

func (e *Engine) SetHorsepower(hp int) {
	e.horsepower = hp
}

func (e Engine) Horsepower() int {
	return e.horsepower
}
