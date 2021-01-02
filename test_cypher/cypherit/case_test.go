package cypherit

import (
	"github.com/manhcuongbk56/cypher-go-dsl"
	"testing"
)

func TestSimpleCase(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNode("a").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Where(cypher.CaseExpression(node.Property("value")).
			When(cypher.LiteralOf("blubb")).
			Then(cypher.CypherLiteralTrue())).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`) WHERE CASE n.value WHEN 'blubb' THEN true END RETURN n")
}

func TestSimpleCaseWithElse(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNode("a").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Where(cypher.CaseExpression(node.Property("value")).
			When(cypher.LiteralOf("blubb")).
			Then(cypher.CypherLiteralTrue()).
			ElseDefault(cypher.CypherLiteralFalse())).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`) WHERE CASE n.value WHEN 'blubb' THEN true ELSE false END RETURN n")
}

func TestSimpleCaseWithMultipleWhenThen(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNode("a").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Where(cypher.CaseExpression(node.Property("value")).
			When(cypher.LiteralOf("blubb")).
			Then(cypher.CypherLiteralTrue()).
			When(cypher.LiteralOf("bla")).
			Then(cypher.CypherLiteralFalse())).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`) WHERE CASE n.value WHEN 'blubb' THEN true WHEN 'bla' THEN false END RETURN n")
}

func TestSimpleCaseWithMultipleWhenThenAndElse(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNode("a").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Where(cypher.CaseExpression(node.Property("value")).
			When(cypher.LiteralOf("blubb")).
			Then(cypher.CypherLiteralTrue()).
			When(cypher.LiteralOf("bla")).
			Then(cypher.CypherLiteralFalse()).
			ElseDefault(cypher.LiteralOf(1))).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`) WHERE CASE n.value WHEN 'blubb' THEN true WHEN 'bla' THEN false ELSE 1 END RETURN n")
}

func TestGenericCase(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNode("a").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Where(cypher.GenericCaseExpression().
			When(node.Property("value").IsEqualTo(cypher.LiteralOf("blubb")).Get()).
			Then(cypher.CypherLiteralTrue())).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`) WHERE CASE WHEN n.value = 'blubb' THEN true END RETURN n")
}

func TestGenericCaseWithElse(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNode("a").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Where(cypher.GenericCaseExpression().
			When(node.Property("value").IsEqualTo(cypher.LiteralOf("blubb")).Get()).
			Then(cypher.CypherLiteralTrue()).
			ElseDefault(cypher.CypherLiteralFalse())).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`) WHERE CASE WHEN n.value = 'blubb' THEN true ELSE false END RETURN n")
}

func TestGenericCaseWithMultipleWhenThen(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNode("a").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Where(cypher.GenericCaseExpression().
			When(node.Property("value").IsEqualTo(cypher.LiteralOf("blubb")).Get()).
			Then(cypher.CypherLiteralTrue()).
			When(node.Property("value").IsEqualTo(cypher.LiteralOf("bla")).Get()).
			Then(cypher.CypherLiteralFalse())).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`) WHERE CASE WHEN n.value = 'blubb' THEN true WHEN n.value = 'bla' THEN false END RETURN n")
}

func TestGenericCaseWithMultipleWhenThenAndElse(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.NewNode("a").NamedByString("n")
	//
	builder = cypher.
		MatchElements(node).
		Where(cypher.GenericCaseExpression().
			When(node.Property("value").IsEqualTo(cypher.LiteralOf("blubb")).Get()).
			Then(cypher.CypherLiteralTrue()).
			When(node.Property("value").IsEqualTo(cypher.LiteralOf("bla")).Get()).
			Then(cypher.CypherLiteralFalse()).
			ElseDefault(cypher.LiteralOf(1))).
		ReturningByNamed(node)

	Assert(t, builder, "MATCH (n:`a`) WHERE CASE WHEN n.value = 'blubb' THEN true WHEN n.value = 'bla' THEN false ELSE 1 END RETURN n")
}

// from https://neo4j.com/docs/cypher-manual/current/syntax/expressions/#syntax-simple-case
func TestCanGetAliasedInReturn(t *testing.T) {
	var builder cypher.BuildableStatement
	node := cypher.AnyNodeNamed("n")
	//
	builder = cypher.
		MatchElements(node).
		Returning(cypher.CaseExpression(node.Property("eyes")).
			When(cypher.LiteralOf("blue")).
			Then(cypher.LiteralOf(1)).
			When(cypher.LiteralOf("brown")).
			Then(cypher.LiteralOf(2)).
			ElseDefault(cypher.LiteralOf(3)).
			As("result").Get())

	Assert(t, builder, "MATCH (n) RETURN CASE n.eyes WHEN 'blue' THEN 1 WHEN 'brown' THEN 2 ELSE 3 END AS result")
}
