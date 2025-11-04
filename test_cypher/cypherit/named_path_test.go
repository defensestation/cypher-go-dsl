package cypherit

import (
	"github.com/defensestation/cypher-go-dsl"
	"testing"
)

func TestDoc3148(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	namePath := cypher.
		APath("p").
		DefinedBy(cypher.AnyNodeNamed("123456789","michael").
			WithRawProperties("name", cypher.LiteralOf("Michael Douglas")).
			RelationshipTo(cypher.AnyNode("123456789")))
	builder = cypher.Match(namePath).
		ReturningByNamed(namePath)

	Assert(t, builder, "MATCH p = (michael {name: 'Michael Douglas'})-->() RETURN p")
}

func TestShouldWorkInListComprehensions(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	namePath := cypher.
		APath("p").
		DefinedBy(cypher.AnyNodeNamed("123456789","n").
			RelationshipTo(cypher.AnyNode("123456789"), "LIKES", "OWNS").
			Unbounded())
	builder = cypher.CypherReturning(cypher.ListBasedOnNamed(namePath).ReturningByNamed(namePath))

	Assert(t, builder, "RETURN [p = (n)-[:`LIKES`|`OWNS`*]->() | p]")
}
