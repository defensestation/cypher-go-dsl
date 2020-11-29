package cypher_go_dsl

import "fmt"

type Skip struct {
	skipAmount NumberLiteral
	key        string
}

func (s Skip) getKey() string {
	return s.key
}

func (s Skip) accept(visitor *CypherRenderer) {
	s.key = fmt.Sprint(&s)
	(*visitor).enter(s)
	s.skipAmount.accept(visitor)
	(*visitor).Leave(s)
}

func CreateSkip(number int) *Skip {
	if number == 0 {
		return nil
	}
	literal := NumberLiteral{
		content: number,
	}
	return &Skip{skipAmount: literal}
}

func (s Skip) enter(renderer *CypherRenderer) {
	renderer.builder.WriteString(" SKIP ")
}

func (s Skip) leave(renderer *CypherRenderer) {
}
