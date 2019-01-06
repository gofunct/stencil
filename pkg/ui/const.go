package ui

var (
	creatableStatusSet = map[status]struct{}{
		StatusCreate: {},
		StatusForce:  {},
	}
)

type status int

const (
	StatusCreate status = iota
	StatusDelete
	StatusExist
	StatusIdentical
	StatusConflicted
	StatusForce
	StatusSkipped
)
