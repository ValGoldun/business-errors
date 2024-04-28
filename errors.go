package business_errors

type Error struct {
	Text       string
	IsCritical bool
}

func New(err string) error {
	return Error{
		Text:       err,
		IsCritical: false,
	}
}

func NewCritical(err string) error {
	return Error{
		Text:       err,
		IsCritical: true,
	}
}

func (e Error) Error() string {
	return e.Text
}
