package creationalPatterns

import "fmt"

type QueryBuilder interface {
	Table(table string) QueryBuilder
	Select(cols []string) QueryBuilder
	Limit(value int) QueryBuilder
	Where(col string, value interface{}) QueryBuilder

	GetQuery() string
}

// Concrete Builder
type MySqlQueryBuilder struct {
	tableName string
	columns   []string
	whereCol  string
	whereVal  interface{}
	limitVal  int
}

// initialisation func that creates a memory record
func NewMySqlQueryBuilder() *MySqlQueryBuilder {
	return &MySqlQueryBuilder{}
}

func (s *MySqlQueryBuilder) Table(table string) QueryBuilder {
	s.tableName = table

	return s
}

func (s *MySqlQueryBuilder) Select(cols []string) QueryBuilder {
	s.columns = cols

	return s
}

func (s *MySqlQueryBuilder) Limit(value int) QueryBuilder {
	s.limitVal = value

	return s
}

// here the val interface is used to represent the union type between int and string.
// I dont think this is very good as it could create errors, probably better to have
// strict typing as it is safer, and there are only two expected types (into or string)
func (s *MySqlQueryBuilder) Where(col string, val interface{}) QueryBuilder {
	s.whereCol = col
	s.whereVal = val

	return s
}

func (s MySqlQueryBuilder) GetQuery() string {
	return "This is MySQL Query"
}

// Concrete Builder
type MongoDbQueryBuilder struct {
	tableName string
	columns   []string
	whereCol  string
	whereVal  interface{}
	limitVal  int
}

func NewMongoQueryBuilder() *MongoDbQueryBuilder {
	return &MongoDbQueryBuilder{}
}

func (m *MongoDbQueryBuilder) Table(table string) QueryBuilder {
	m.tableName = table

	return m
}

func (m *MongoDbQueryBuilder) Select(cols []string) QueryBuilder {
	m.columns = cols

	return m
}

func (m *MongoDbQueryBuilder) Limit(val int) QueryBuilder {
	m.limitVal = val

	return m
}

func (m *MongoDbQueryBuilder) Where(col string, val interface{}) QueryBuilder {
	m.whereCol = col
	m.whereVal = val

	return m
}

func (m *MongoDbQueryBuilder) GetQuery() string {
	return "This is MongoDbQuery"
}

// ClientCode
func Client(builder QueryBuilder) {
	query := builder.
		Table("posts").
		Select([]string{"id", "title"}).
		Limit(5).
		GetQuery()

	fmt.Println(query)
}

func ExampleBuilder() {
	mySqlBuilder := NewMySqlQueryBuilder()
	Client(mySqlBuilder)

	mongoBuilder := NewMongoQueryBuilder()
	Client(mongoBuilder)
}
