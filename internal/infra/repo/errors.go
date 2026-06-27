package repo

import (
	"fmt"
)

// ErrReadFile represents an infrastructure error that occurs when the system
// fails to physically open or read the embedded dictionary asset file from disk.
type ErrReadFile struct {
	FileName string
	Err      error
}

func (e *ErrReadFile) Error() string {
	return fmt.Sprintf("failed to read dict file %q: %v", e.FileName, e.Err)
}

// ErrUnmarshalJSON represents an error that occurs when the read dictionary asset data
// string payload cannot be successfully parsed due to corrupted or invalid JSON structural formatting.
type ErrUnmarshalJSON struct {
	FileName string
	Err      error
}

func (e *ErrUnmarshalJSON) Error() string {
	return fmt.Sprintf("failed to unmarshal dict json from %q: %v", e.FileName, e.Err)
}
