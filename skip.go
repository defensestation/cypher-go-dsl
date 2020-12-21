package cypher

type Skip struct {
	skipAmount NumberLiteral
	key        string
	notNil     bool
	err        error
}

func SkipCreate(number int) Skip {
	literal := NumberLiteralCreate(number)
	skip := Skip{skipAmount: literal}
	skip.key = getAddress(&skip)
	return skip
}

func (s Skip) getError() error {
	return s.err
}

func (s Skip) isNotNil() bool {
	return s.notNil
}

func (s Skip) getKey() string {
	return s.key
}

func (s Skip) accept(visitor *CypherRenderer) {
	(*visitor).enter(s)
	s.skipAmount.accept(visitor)
	(*visitor).leave(s)
}

func (s Skip) enter(renderer *CypherRenderer) {
	renderer.append(" SKIP ")
}

func (s Skip) leave(renderer *CypherRenderer) {
}
