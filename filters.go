package cfclient

import "fmt"

// Filter represents a query filter
// Parameters used to filter the result set.
// Valid ops: : >= <= < > IN
type Filter struct {
	Filter string
	Ops    string
	Value  string
}

type Filters []Filter

// ToParam formats a query as <filter><op><value>
func (filter Filter) ToParam() string {
	return fmt.Sprintf("q=%s%s%s", filter.Filter, filter.Ops, filter.Value)
}

// ToParam formats queries as <filter><op><value>&<filter><op><value>
func (filters *Filters) ToParam() (query string) {
	for _, filter := range *filters {
		if len(query) > 0 {
			query = query + "&"
		}
		query = query + filter.ToParam()
	}
	return
}
