package goesdsl

// QueriesRewriteParameter https://www.elastic.co/docs/reference/query-languages/query-dsl/query-dsl-multi-term-rewrite
const (
	QueriesRewriteParameterConstantScoreBlended  = "constant_score_blended" // Default
	QueriesRewriteParameterConstantScore         = "constant_score"
	QueriesRewriteParameterConstantScoreBoolean  = "constant_score_boolean"
	QueriesRewriteParameterScoringBoolean        = "scoring_boolean"
	QueriesRewriteParameterTopTermsBlendedFreqsN = "top_terms_blended_freqs_N"
	QueriesRewriteParameterTopTermsBoostN        = "top_terms_boost_N"
	QueriesRewriteParameterTopTermsN             = "top_terms_N"
)
