package dashingtypes

// An Event contains the widget ID, a body of data,
// and an optional target (only "dashboard" for now).
type Event struct {
	ID     string
	Body   map[string]interface{}
	Target string
}
