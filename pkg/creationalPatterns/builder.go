package creationalPatterns 

import "fmt"

type QueryBuilder interface {
	table(table string) QueryBuilder 
	select(cols string) QueryBuilder 
	limit(value int) QueryBuilder 
	where(col string, value int) QueryBuilder

	getQuery() string
}

// Concrete Builder
type MySqlQueryBuilder struct {
	query string 
}

func (s MySqlQueryBuilder) table(table string) MySqlQueryBuilder {
	// TODO
}

func (s MySqlQueryBuilder) select(cols string) MySqlQueryBuilder {
	// TODO
}

func (s MySqlQueryBuilder) limit(value int) MySqlQueryBuilder {
	// TODO 
}

func (s MySqlQueryBuilder) where(col string, val int) MySqlQueryBuilder {

}

func (s MySqlQueryBuilder) getQuery() string {
	return "This is MySQL Query"
}
