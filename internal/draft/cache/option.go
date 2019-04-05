package cache

const (
	Second Duration = 1
	Minute          = 60 * Second
	Hour            = 60 * Minute
	Day             = 24 * Hour
	Week            = 7 * Day

	Any       State = 0
	Exists          = 1
	NotExists       = 2
)

type Duration int

type State uint8

type Get struct{}

type GetOption func(*Get)

type Set struct {
	State State
	TTL   Duration
}

type SetOption func(*Set)

func SetIfExists() SetOption {
	return func(option *Set) { option.State = Exists }
}

func SetIfNotExists() SetOption {
	return func(option *Set) { option.State = NotExists }
}

func SetWithTimeout(d Duration) SetOption {
	return func(option *Set) { option.TTL = d }
}
