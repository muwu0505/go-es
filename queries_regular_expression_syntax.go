package goesdsl

// QueriesFlags https://www.elastic.co/docs/reference/query-languages/query-dsl/regexp-syntax
const (
	QueriesFlagsAll          = "ALL" // default
	QueriesFlagsComplement   = "COMPLEMENT"
	QueriesFlagsEmpty        = "EMPTY"
	QueriesFlagsInterval     = "INTERVAL"
	QueriesFlagsIntersection = "INTERSECTION"
	QueriesFlagsAnystring    = "ANYSTRING"
	QueriesFlagsNone         = "NONE"
)
