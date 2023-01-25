// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     BaseCoreRegistryJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package corekind

import (
	"fmt"

	"github.com/grafana/grafana/pkg/kinds/dashboard"
	"github.com/grafana/grafana/pkg/kinds/playlist"
	"github.com/grafana/grafana/pkg/kinds/star"
	"github.com/grafana/grafana/pkg/kinds/team"
	"github.com/grafana/grafana/pkg/kindsys"
	"github.com/grafana/thema"
)

// Base is a registry of kindsys.Interface. It provides two modes for accessing
// kinds: individually via literal named methods, or as a slice returned from
// an All*() method.
//
// Prefer the individual named methods for use cases where the particular kind(s) that
// are needed are known to the caller. For example, a dashboard linter can know that it
// specifically wants the dashboard kind.
//
// Prefer All*() methods when performing operations generically across all kinds.
// For example, a validation HTTP middleware for any kind-schematized object type.
type Base struct {
	all       []kindsys.Core
	dashboard *dashboard.Kind
	playlist  *playlist.Kind
	star      *star.Kind
	team      *team.Kind
}

// type guards
var (
	_ kindsys.Core = &dashboard.Kind{}
	_ kindsys.Core = &playlist.Kind{}
	_ kindsys.Core = &star.Kind{}
	_ kindsys.Core = &team.Kind{}
)

// Dashboard returns the [kindsys.Interface] implementation for the dashboard kind.
func (b *Base) Dashboard() *dashboard.Kind {
	return b.dashboard
}

// Playlist returns the [kindsys.Interface] implementation for the playlist kind.
func (b *Base) Playlist() *playlist.Kind {
	return b.playlist
}

// Star returns the [kindsys.Interface] implementation for the star kind.
func (b *Base) Star() *star.Kind {
	return b.star
}

// Team returns the [kindsys.Interface] implementation for the team kind.
func (b *Base) Team() *team.Kind {
	return b.team
}

func doNewBase(rt *thema.Runtime) *Base {
	var err error
	reg := &Base{}

	reg.dashboard, err = dashboard.NewKind(rt)
	if err != nil {
		panic(fmt.Sprintf("error while initializing the dashboard Kind: %s", err))
	}
	reg.all = append(reg.all, reg.dashboard)

	reg.playlist, err = playlist.NewKind(rt)
	if err != nil {
		panic(fmt.Sprintf("error while initializing the playlist Kind: %s", err))
	}
	reg.all = append(reg.all, reg.playlist)

	reg.star, err = star.NewKind(rt)
	if err != nil {
		panic(fmt.Sprintf("error while initializing the star Kind: %s", err))
	}
	reg.all = append(reg.all, reg.star)

	reg.team, err = team.NewKind(rt)
	if err != nil {
		panic(fmt.Sprintf("error while initializing the team Kind: %s", err))
	}
	reg.all = append(reg.all, reg.team)

	return reg
}
