package util

func NewString(str string) *string {
	return &str
}

func BlankToNil(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func NilToBlank(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
