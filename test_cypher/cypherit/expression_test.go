package cypherit

import (
	"github.com/defensestation/cypher-go-dsl"
	"testing"
)

func TestShouldRenderParameters(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		Match(userNode).
		WhereConditionContainer(userNode.Property("a").IsEqualTo(cypher.AParam("aParameter"))).
		DetachDeleteByNamed(userNode).
		ReturningByNamed(userNode)
	Assert(t, builder, "MATCH (u:`dsc_User` {}) WHERE u.a = $aParameter DETACH DELETE u RETURN u")
}

func TestShouldRenderMap(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	builder = cypher.
		Match(cypher.AnyNodeNamed("123456789","n")).
		Returning(cypher.Point(cypher.MapOf(
			"latitude", cypher.AParam("latitude"),
			"longitude", cypher.AParam("longitude"),
			"crs", cypher.LiteralOf(4326))))
	Assert(t, builder, "MATCH (n) RETURN point({latitude: $latitude, longitude: $longitude, crs: 4326})")
}

func TestShouldRenderPointFunction(t *testing.T) {
	var builder cypher.BuildableStatement
	//
	n := cypher.AnyNodeNamed("123456789","n")
	builder = cypher.
		Match(n).
		WhereConditionContainer(cypher.Distance(n.Property("location"),
			cypher.PointByParameter(cypher.AParam("point.point"))).
			Gt(cypher.AParam("point.distance"))).
		ReturningByNamed(n)
	Assert(t, builder, "MATCH (n) WHERE distance(n.location, point($point.point)) > $point.distance RETURN n")
}
