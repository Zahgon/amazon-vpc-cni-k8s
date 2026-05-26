package datastore

// Checkpointer can persist data and (hopefully) restore it later
type Checkpointer interface {
	Checkpoint(data interface{}) error
	Restore(into interface{}) error
}

// NullCheckpoint discards data and always returns "not found". For testing only!
type NullCheckpoint struct{}

// Checkpoint implements the Checkpointer interface in the most
// trivial sense, by just discarding data.
func (c NullCheckpoint) Checkpoint(data interface{}) error {
	_ = "STUB: not implemented"

	// Restore implements the Checkpointer interface in the most trivial
	// sense, by always returning "not found".
	return nil
}

func (c NullCheckpoint) Restore(into interface{}) error { _ = "STUB: not implemented"; return nil }

// TestCheckpoint maintains a snapshot in memory.
type TestCheckpoint struct {
	Error error
	Data  interface{}
}

// NewTestCheckpoint creates a new TestCheckpoint.
func NewTestCheckpoint(data interface{}) *TestCheckpoint { _ = "STUB: not implemented"; return nil }

// Checkpoint implements the Checkpointer interface.
func (c *TestCheckpoint) Checkpoint(data interface{}) error { _ = "STUB: not implemented"; return nil }

// Restore implements the Checkpointer interface.
func (c *TestCheckpoint) Restore(into interface{}) error { _ = "STUB: not implemented"; return nil }

// `into` is always a pointer to interface{}, but we can't
// actually make the Restore() function *interface{}, because
// that doesn't match the (widely used) `encoding.Unmarshal`
// interface :(
// Round trip through json strings instead because copying is
// hard.

// JSONFile is a checkpointer that writes to a JSON file
type JSONFile struct {
	path string
}

// NewJSONFile creates a new JsonFile
func NewJSONFile(path string) *JSONFile { _ = "STUB: not implemented"; return nil }

// Checkpoint implements the Checkpointer interface
func (c *JSONFile) Checkpoint(data interface{}) error { _ = "STUB: not implemented"; return nil }

// Restore implements the Checkpointer interface
func (c *JSONFile) Restore(into interface{}) error { _ = "STUB: not implemented"; return nil }
