package exceptions

type MissingField string

// Error{ Kind: Validation, Param: "title", Inner: MissingField("title") }
func (e MissingField) Error() string {
	return string(e) + " is required"
}

// Error{ Kind: Validation, Code:"invalid_date_format", Param: "release_date", Inner: err}
