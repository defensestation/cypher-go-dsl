package cypher_go_dsl

import "errors"

type ListExpression struct {
	content ExpressionList
	err     error
	key     string
	notNil  bool
}

func ListExpressionCreate(content ExpressionList) ListExpression {
	if content.getError() != nil {
		return ListExpressionError(content.getError())
	}
	list := ListExpression{
		content: content,
		notNil:  true,
	}
	list.key = getAddress(&list)
	return list
}

func ListExpressionCreate1(expressions ...Expression) ListExpression {
	return ListExpressionCreate(ExpressionListCreate(expressions))
}

func ListExpressionError(err error) ListExpression {
	return ListExpression{
		err: err,
	}
}

func ListOrSingleExpression(expressions ...Expression) Expression {
	if expressions == nil || len(expressions) == 0 {
		ListExpressionError(errors.New("expressions are required"))
	}
	if len(expressions) == 1 {
		return expressions[0]
	} else {
		return ListExpressionCreate1(expressions...)
	}
}

func (l ListExpression) getError() error {
	return l.err
}

func (l ListExpression) accept(visitor *CypherRenderer) {
	visitor.enter(l)
	l.content.accept(visitor)
	visitor.leave(l)
}

func (l ListExpression) enter(renderer *CypherRenderer) {
	renderer.append("[")
}

func (l ListExpression) leave(renderer *CypherRenderer) {
	renderer.append("]")
}

func (l ListExpression) getKey() string {
	return l.key
}

func (l ListExpression) isNotNil() bool {
	return l.notNil
}

func (l ListExpression) GetExpressionType() ExpressionType {
	return "ListExpression"
}