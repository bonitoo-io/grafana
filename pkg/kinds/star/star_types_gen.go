// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     GoTypesJenny
//     LatestJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package star

// Star defines model for star.
type Star struct {
	// The dashboard which is starred.
	DashboardId int64 `json:"dashboardId"`

	// ID of the star.
	Id int64 `json:"id"`

	// UserID is the ID of an user the star belongs to.
	UserId int64 `json:"userId"`
}
