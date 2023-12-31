package result

const (
	Success         State = 0b0001
	Failure         State = 0b0010
	ExpectedFailure State = 0b0011
)

type State int

func (s State) String() string {
	switch s {
	case Success:
		return "Success"

	case Failure:
		return "Failure"

	case ExpectedFailure:
		return "ExpectedFailure"

	default:
		return "Unknown"

	}
}
