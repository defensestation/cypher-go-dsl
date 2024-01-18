package cypher

import (
	"errors"
	"fmt"
)

func AccountIdToLabel(accountId string) (string, error) {
	if accountId == "" {
		return "", errors.New("account Id can't be nil")
	}
	return fmt.Sprintf("%s%s", "dsc_", accountId), nil
}

type Node struct {
	accountId    string
	symbolicName SymbolicName
	labels       []NodeLabel
	properties   Properties
	key          string
	notNil       bool
	err          error
}


func NodeCreate(accountId string) Node {
	primaryLabel, err := AccountIdToLabel(accountId)
	if err != nil {
		return NodeError(err)
	}

	return NodeCreate2(primaryLabel)
}

func NodeCreate1(accountId string, properties Properties, additionalLabels ...string) Node {
	
	primaryLabel, err := AccountIdToLabel(accountId)
	if err != nil {
		return NodeError(err)
	}

	for _, label := range additionalLabels {
		if label == "" {
			return NodeError(errors.New("empty label is not allowed"))
		}
	}
	labels := make([]NodeLabel, 0)
	labels = append(labels, NodeLabel{value: primaryLabel})
	for _, label := range additionalLabels {
		labels = append(labels, NodeLabelCreate(label))
	}
	node := Node{
		accountId: accountId,
		notNil:     true,
		properties: properties,
		labels:     labels,
	}
	node.injectKey()
	return node
}

func NodeCreate2(accountId string) Node {
	primaryLabel, err := AccountIdToLabel(accountId)
	if err != nil {
		return NodeError(err)
	}
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabelCreate(primaryLabel))
	node := Node{
		accountId: accountId,
		labels: labels,
		notNil: true,
		properties: PropertiesCreate(NewMapExpression()),
	}
	node.injectKey()
	return node
}


func NodeCreate3(accountId string, additionalLabel ...string) Node {
	primaryLabel, err := AccountIdToLabel(accountId)
	if err != nil {
		return NodeError(err)
	}
	for _, label := range additionalLabel {
		if label == "" {
			return NodeError(errors.New("empty label is not allowed"))
		}
	}
	var labels = make([]NodeLabel, 0)
	labels = append(labels, NodeLabel{value: primaryLabel})
	for _, label := range additionalLabel {
		labels = append(labels, NodeLabelCreate(label))
	}
	node := Node{
		accountId: accountId,
		labels: labels,
		properties: PropertiesCreate(NewMapExpression()),
	}
	node.injectKey()
	return node
}

func NodeCreate4(newProperties MapExpression, node Node) Node {
	newNode := Node{symbolicName: node.symbolicName, labels: node.labels, notNil: true, properties: PropertiesCreate(newProperties)}
	newNode.injectKey()
	return newNode
}

func NodeCreate5(primaryLabel string, properties MapExpression, additionalLabel ...string) Node {
	return NodeCreate1(primaryLabel, PropertiesCreate(properties), additionalLabel...)
}

func (node *Node) injectKey() {
	node.key = getAddress(node)
}

func NodeError(err error) Node {
	return Node{
		err: err,
	}
}

func (node Node) GetRequiredSymbolicName() SymbolicName {
	if node.symbolicName.isNotNil() {
		return node.symbolicName
	}
	return SymbolicNameError(errors.New("node get symbolic name:  no name present"))
}

func (node Node) AddLabels(labels ...string) Node {
	for _, label := range labels {
		node.labels = append(node.labels, NodeLabelCreate(label))
	}
	return node
}

func (node Node) GetSymbolicName() SymbolicName {
	return node.symbolicName
}

func (node Node) GetSymbolicValue() string {
	return node.symbolicName.GetValue()
}

func (node Node) GetAccountId() string {
	return node.accountId
}

func (node Node) isNotNil() bool {
	return node.notNil
}

func (node Node) IsPatternElement() bool {
	return true
}

func (node Node) getKey() string {
	return node.key
}

func (node Node) hasSymbolic() bool {
	return node.symbolicName.isNotNil()
}

func (node Node) GetError() error {
	return node.err
}

func (node Node) accept(visitor *CypherRenderer) {
	visitor.enter(node)
	VisitIfNotNull(node.symbolicName, visitor)
	for _, label := range node.labels {
		label.accept(visitor)
	}
	VisitIfNotNull(node.properties, visitor)
	visitor.leave(node)
}

func (node Node) enter(renderer *CypherRenderer) {
	renderer.append("(")
	if !node.hasSymbolic() {
		return
	}
	_, renderer.skipNodeContent = renderer.visitedNamed[node.key]
	renderer.visitedNamed[node.key] = 1
	if renderer.skipNodeContent {
		renderer.append(node.symbolicName.value)
	}
}

func (node Node) leave(renderer *CypherRenderer) {
	renderer.append(")")
	renderer.skipNodeContent = false
}

func (node Node) RelationshipTo(other Node, types ...string) Relationship {
	return RelationshipCreate(node, LTR(), other, types...)
}

func (node Node) RelationshipFrom(other Node, types ...string) Relationship {
	return RelationshipCreate(node, RTL(), other, types...)
}

func (node Node) RelationshipBetween(other Node, types ...string) Relationship {
	return RelationshipCreate(node, UNI(), other, types...)
}

func (node Node) WithRawProperties(keysAndValues ...interface{}) Node {
	properties := MapExpression{}
	if keysAndValues != nil && len(keysAndValues) != 0 {
		properties = NewMapExpression(keysAndValues...)
		if properties.GetError() != nil {
			return NodeError(properties.GetError())
		}
	}
	return node.WithProperties(properties)
}

func (node Node) WithProperties(newProperties MapExpression) Node {
	return NodeCreate4(newProperties, node)
}

func (node Node) Property(name string) Property {
	return PropertyCreate(node, name)
}

func (node Node) NamedByString(name string) Node {
	node.symbolicName = SymbolicNameCreate(name)
	return node
}

func (node Node) Named(name SymbolicName) Node {
	if name.GetError() != nil {
		return NodeError(name.GetError())
	}
	if !name.isNotNil() {
		return NodeError(errors.New("node named: symbolic name is required"))
	}
	node.symbolicName = name
	return node
}

func (node Node) As(alias string) AliasedExpression {
	symbolicName := node.GetRequiredSymbolicName()
	if symbolicName.GetError() != nil {
		return AliasedExpressionError(symbolicName.GetError())
	}
	return ExpressionWrap(symbolicName).As(alias).Get().(AliasedExpression)
}

func (node Node) InternalId() FunctionInvocation {
	return IdByNode(node)
}

func (node Node) HasLabels(labelsToQuery ...string) Condition {
	return HasLabelConditionCreate(node.GetRequiredSymbolicName(), labelsToQuery...)
}

func (node Node) Labels() FunctionInvocation {
	return Labels(node)
}

func (node Node) AddProperties(entries []Expression) Node {
	node.properties = node.properties.AddProperties(entries)
	return node
}

func (node Node) Project(entries ...interface{}) MapProjection {
	return MapProjectionCreate(node.GetRequiredSymbolicName(), entries...)
}

type NodeLabel struct {
	value  string
	key    string
	notNil bool
	err    error
}

func NodeLabelCreate(value string) NodeLabel {
	n := NodeLabel{
		value:  value,
		notNil: true,
	}
	n.key = getAddress(&n)
	return n
}

func NodeLabelError(err error) NodeLabel {
	return NodeLabel{err: err}
}

func (n NodeLabel) GetError() error {
	return n.err
}

func (n NodeLabel) GetValue() string {
	return n.value
}

func (n NodeLabel) isNotNil() bool {
	return n.notNil
}

func (n NodeLabel) getKey() string {
	return n.key
}

func (n NodeLabel) accept(visitor *CypherRenderer) {
	visitor.enter(n)
	visitor.leave(n)
}

func (n NodeLabel) enter(renderer *CypherRenderer) {
	if n.value == "" {
		return
	}
	renderer.append(NodeLabelStart)
	renderer.append(escapeName(n.value))
}

func (n NodeLabel) leave(renderer *CypherRenderer) {
}
