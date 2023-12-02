package utility

type Validator interface {
	Validate(*map[string]string) bool
}

type Game interface {
	Id() string
	Info() *map[string]string
}

type GameRound interface {
	Info() *map[string]string
}

type PowerBehavior interface {
	Power() int
}

/*
Utility function to implement contains for a list of strings.
*/
func ListContainsString(items *[]string, searchTerm string) bool {
	// iterate through item in list
	for _, item := range *items {
		if item == searchTerm {
			return true
		}
	}
	return false
}
