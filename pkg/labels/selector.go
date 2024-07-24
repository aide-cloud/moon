package labels

type Selector interface {
	// Matches returns true if the specified Labels match this Selector.
	Matches(labels Labels) bool

	// AddRequirement adds a new Requirement to this Selector.
	AddRequirement(r ...Requirement) Selector
}

var _ Selector = &emptySelector{}

type emptySelector struct{}

func (e *emptySelector) Matches(labels Labels) bool {
	return false
}

func (e *emptySelector) AddRequirement(r ...Requirement) Selector {
	return nil
}

type Requirements []Requirement

func NewSelector() Selector {
	return Requirements(nil)
}

func (x Requirements) Matches(labels Labels) bool {
	for i := range x {
		if !x[i].Matches(labels) {
			return false
		}
	}
	return true
}

func (x Requirements) AddRequirement(r ...Requirement) Selector {
	req := make(Requirements, 0, len(x)+len(r))
	req = append(req, x...)
	req = append(req, r...)
	return req
}
