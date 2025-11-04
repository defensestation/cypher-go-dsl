package cypher

import "testing"

// func TestGh48(t *testing.T) {
// 	n := ANode("Label").NamedByString("n")
// 	statement, err := Match(n).
// 		SetWithNamed(n, MapOf("a", StringLiteralCreate("bar"), "b", StringLiteralCreate("baz"))).
// 		ReturningByNamed(n).
// 		Build()
// 	if err != nil {
// 		t.Errorf("error When build query: %s", err)
// 	}
// 	query, _ := NewRenderer().Render(statement)
// 	if query != "MATCH (n:`Label`) SET n = {`a`: 'bar', `b`: 'baz'} RETURN n" {
// 		t.Errorf("Query is not MatchPhrase: %s", query)
// 	}
// }

func TestComplexCypherQuery(t *testing.T) {
	// Construct the various Cypher components
	endNode := ANode("EndLabel").NamedByString("endNode")
	deleteNodes := ANode("DeleteLabel").NamedByString("deleteNodes")
	startNode := ANode("StartLabel").NamedByString("startNode")
	relation := startNode.RelationshipTo(endNode).NamedByString("relation")
	conditions := endNode.InternalId().IsEqualTo(LiteralOf("dfg")) // Replace with your actual conditions
	documentToRoleCounter := "documentToRoleCounter"
	test := []Expression{
		relation.Property(documentToRoleCounter),
		LiteralOf(LiteralOf(1).Add(ASymbolic("asd")).Get())}

	t.Logf("Generated query: %+v\n", len(test))

	// Build the statement
	statement, err := Match(endNode).
		Match(deleteNodes).
		WhereConditionContainer(conditions).
		Merge(startNode).
		Merge(relation).
		// OnCreate().
		// Set(
		// 	relation.Property(documentToRoleCounter),  // property name
		// 	LiteralOf(1),                              // value
		// ).
		OnMatch().
		Set(
			// relation.Property(documentToRoleCounter),
			OperationSet(relation.Property(documentToRoleCounter), relation.Property(documentToRoleCounter).Add(LiteralOf(1)).Get()),
		    // LiteralOf(1).Add(ASymbolic("asd")).Get(),
		    // LiteralOf(1),
		).
		DetachDeleteByNamed(deleteNodes).
		Returning(startNode.GetSymbolicName()).
		Build()

	if err != nil {
		t.Fatalf("Failed to build query: %v", err)
	}

	// Render the query
	query, renderErr := NewRenderer().Render(statement)
	if renderErr != nil {
		t.Fatalf("Failed to render query: %v", renderErr)
	}

	// Print the query to see the output
	// t.Logf("Generated query: %s", query)


	// Assert the expected query
	expectedQuery := "MATCH (endNode:`EndLabel`), (deleteNodes:`DeleteLabel`) " +
		"WHERE <conditions> " +
		"MERGE (startNode:`StartLabel`) " +
		"MERGE (relation:`RELATES_TO`) " +
		"ON CREATE SET relation.`documentToRoleCounter` = 1 " +
		"ON MATCH SET relation.`documentToRoleCounter` = relation.`documentToRoleCounter` + 1 " +
		"DETACH DELETE deleteNodes " +
		"RETURN startNode"

	if query != expectedQuery {
		t.Errorf("Query does not match expected output.\nExpected: %s\nGot: %s", expectedQuery, query)
	}
}




//func TestGh51(t *testing.T) {
//	n := CypherAnyNode1("n")
//	foobarProp := proper
//	statement, err := Match(n).
//		SetWithNamed(n, MapOf("a", StringLiteralCreate("bar"), "b", StringLiteralCreate("baz"))).
//		ReturningByNamed(n).
//		Build()
//	if err != nil {
//		t.Errorf("error When build query: %s", err)
//	}
//	query := NewRenderer().Render(statement)
//	if query != "MATCH (n:`Label`) SET n = {`a`: 'bar', `b`: 'baz'} RETURN n" {
//		t.Errorf("Query is not MatchPhrase: %s", query)
//	}
//}
