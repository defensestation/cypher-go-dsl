package cypher

type ExposesWhere interface {
	Where(condition Condition) OngoingReadingWithWhere
	WherePattern(pattern RelationshipPattern) OngoingReadingWithWhere
}
