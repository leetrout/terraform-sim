package resources

// Group allows grouping entities together in
// an EntitySet.
type Group struct {
	Name      string   `validate:"required"`
	EntitySet []string `validate:"required,dive,uuid"`
}
