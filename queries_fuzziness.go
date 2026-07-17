package goesdsl

import "fmt"

// QueriesFuzziness https://www.elastic.co/docs/reference/elasticsearch/rest-apis/common-options#fuzziness
const (
	QueriesFuzzinessAuto     = "AUTO" // Preferred value; equivalent to AUTO:3,6
	QueriesFuzzinessExact    = "0"    // Must match exactly
	QueriesFuzzinessOneEdit  = "1"    // One edit allowed
	QueriesFuzzinessTwoEdits = "2"    // Two edits allowed
)

func CustomQueriesFuzzinessAuto(low, high int) string {
	return fmt.Sprintf("AUTO:%d,%d", low, high)
}
