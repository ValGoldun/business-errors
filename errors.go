package business_errors

import "fmt"

type Error struct {
	Text       string
	IsCritical bool
	Metadata   Metadata
}

func New(err string, metadata Metadata) error {
	return Error{
		Text:       err,
		IsCritical: false,
		Metadata:   metadata,
	}
}

func NewCritical(err string, metadata Metadata) error {
	return Error{
		Text:       err,
		IsCritical: true,
		Metadata:   metadata,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s, %s", e.Text, e.Metadata.String())
}
