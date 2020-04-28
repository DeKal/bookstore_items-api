package queries

import "github.com/olivere/elastic"

// EsQuery elastic search query
type EsQuery struct {
	ShouldEqual []EsFieldValue `json:"equal"`
}

// EsFieldValue elasticsearch field condition
type EsFieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

// Build build es query
func (q *EsQuery) Build() elastic.Query {
	equalQueries := make([]elastic.Query, 0)
	for _, eq := range q.ShouldEqual {
		matchQuery := elastic.NewMatchQuery(eq.Field, eq.Value)
		equalQueries = append(equalQueries, matchQuery)
	}

	query := elastic.NewBoolQuery()
	query.Must(equalQueries...)
	return query
}
