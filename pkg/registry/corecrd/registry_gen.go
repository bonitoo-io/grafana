// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     CRDKindRegistryJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package corecrd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/grafana/dskit/services"
	"github.com/grafana/grafana/pkg/modules"
	"github.com/grafana/grafana/pkg/registry/corekind"
	"github.com/grafana/grafana/pkg/services/k8s/client"
	"github.com/grafana/grafana/pkg/services/k8s/crd"
	"github.com/grafana/grafana/pkg/services/k8s/informer"
	"github.com/grafana/grafana/pkg/services/k8s/resources/accesspolicy"
	"github.com/grafana/grafana/pkg/services/k8s/resources/dashboard"
	"github.com/grafana/grafana/pkg/services/k8s/resources/folder"
	"github.com/grafana/grafana/pkg/services/k8s/resources/librarypanel"
	"github.com/grafana/grafana/pkg/services/k8s/resources/playlist"
	"github.com/grafana/grafana/pkg/services/k8s/resources/preferences"
	"github.com/grafana/grafana/pkg/services/k8s/resources/publicdashboard"
	"github.com/grafana/grafana/pkg/services/k8s/resources/serviceaccount"
	"github.com/grafana/grafana/pkg/services/k8s/resources/team"
	"github.com/grafana/thema"
	"gopkg.in/yaml.v3"
)

// New constructs a new [Registry].
//
// All calling code within grafana/grafana is expected to use Grafana's
// singleton [thema.Runtime], returned from [cuectx.GrafanaThemaRuntime]. If nil
// is passed, the singleton will be used.
func New(
	rt *thema.Runtime,
	clientsetProvider client.ClientSetProvider,
	informerFactory informer.Informer,
	accesspolicyWatcher accesspolicy.Watcher,
	dashboardWatcher dashboard.Watcher,
	folderWatcher folder.Watcher,
	librarypanelWatcher librarypanel.Watcher,
	playlistWatcher playlist.Watcher,
	preferencesWatcher preferences.Watcher,
	publicdashboardWatcher publicdashboard.Watcher,
	serviceaccountWatcher serviceaccount.Watcher,
	teamWatcher team.Watcher,
) *Registry {
	breg := corekind.NewBase(rt)
	r := doNewRegistry(
		breg,
		clientsetProvider,
		informerFactory,
		accesspolicyWatcher,
		dashboardWatcher,
		folderWatcher,
		librarypanelWatcher,
		playlistWatcher,
		preferencesWatcher,
		publicdashboardWatcher,
		serviceaccountWatcher,
		teamWatcher,
	)
	r.BasicService = services.NewBasicService(r.start, r.run, nil).WithName(modules.KubernetesCRDs)
	return r
}

// All returns a slice of all core Grafana CRDs in the registry.
//
// The returned slice is guaranteed to be alphabetically sorted by kind name.
func (r *Registry) All() []crd.Kind {
	all := make([]crd.Kind, len(r.all))
	copy(all, r.all[:])
	return all
}

// Registry is a list of all of Grafana's core structured kinds, wrapped in a
// standard [crd.CRD] interface that makes them usable for interactions
// with certain Kubernetes controller and apimachinery libraries.
//
// There are two access methods: individually via literal named methods, or as
// a slice returned from All() method.
//
// Prefer the individual named methods for use cases where the particular kind(s)
// that are needed are known to the caller. Prefer All() when performing operations
// generically across all kinds.
type Registry struct {
	*services.BasicService
	all                    [9]crd.Kind
	clientsetProvider      client.ClientSetProvider
	informerFactory        informer.Informer
	accesspolicyWatcher    accesspolicy.Watcher
	dashboardWatcher       dashboard.Watcher
	folderWatcher          folder.Watcher
	librarypanelWatcher    librarypanel.Watcher
	playlistWatcher        playlist.Watcher
	preferencesWatcher     preferences.Watcher
	publicdashboardWatcher publicdashboard.Watcher
	serviceaccountWatcher  serviceaccount.Watcher
	teamWatcher            team.Watcher
}

func (r *Registry) start(ctx context.Context) error {
	var (
		err error
		b   []byte
	)
	clientSet := r.clientsetProvider.GetClientset()

	/************************ AccessPolicy ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map0 := make(map[string]any)
	err = yaml.Unmarshal(accesspolicy.CRDYaml, map0)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for AccessPolicy failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map0)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for AccessPolicy: %s", err))
	}
	err = json.Unmarshal(b, &accesspolicy.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for AccessPolicy: %s", err))
	}

	err = clientSet.RegisterKind(ctx, accesspolicy.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for AccessPolicy failed to register: %s\n", err))
	}

	watcherWrapper0 := accesspolicy.NewWatcherWrapper(r.accesspolicyWatcher)
	r.informerFactory.AddWatcher(accesspolicy.CRD, watcherWrapper0)

	/************************ Dashboard ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map1 := make(map[string]any)
	err = yaml.Unmarshal(dashboard.CRDYaml, map1)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for Dashboard failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map1)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for Dashboard: %s", err))
	}
	err = json.Unmarshal(b, &dashboard.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for Dashboard: %s", err))
	}

	err = clientSet.RegisterKind(ctx, dashboard.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for Dashboard failed to register: %s\n", err))
	}

	watcherWrapper1 := dashboard.NewWatcherWrapper(r.dashboardWatcher)
	r.informerFactory.AddWatcher(dashboard.CRD, watcherWrapper1)

	/************************ Folder ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map2 := make(map[string]any)
	err = yaml.Unmarshal(folder.CRDYaml, map2)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for Folder failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map2)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for Folder: %s", err))
	}
	err = json.Unmarshal(b, &folder.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for Folder: %s", err))
	}

	err = clientSet.RegisterKind(ctx, folder.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for Folder failed to register: %s\n", err))
	}

	watcherWrapper2 := folder.NewWatcherWrapper(r.folderWatcher)
	r.informerFactory.AddWatcher(folder.CRD, watcherWrapper2)

	/************************ LibraryPanel ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map3 := make(map[string]any)
	err = yaml.Unmarshal(librarypanel.CRDYaml, map3)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for LibraryPanel failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map3)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for LibraryPanel: %s", err))
	}
	err = json.Unmarshal(b, &librarypanel.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for LibraryPanel: %s", err))
	}

	err = clientSet.RegisterKind(ctx, librarypanel.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for LibraryPanel failed to register: %s\n", err))
	}

	watcherWrapper3 := librarypanel.NewWatcherWrapper(r.librarypanelWatcher)
	r.informerFactory.AddWatcher(librarypanel.CRD, watcherWrapper3)

	/************************ Playlist ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map4 := make(map[string]any)
	err = yaml.Unmarshal(playlist.CRDYaml, map4)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for Playlist failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map4)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for Playlist: %s", err))
	}
	err = json.Unmarshal(b, &playlist.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for Playlist: %s", err))
	}

	err = clientSet.RegisterKind(ctx, playlist.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for Playlist failed to register: %s\n", err))
	}

	watcherWrapper4 := playlist.NewWatcherWrapper(r.playlistWatcher)
	r.informerFactory.AddWatcher(playlist.CRD, watcherWrapper4)

	/************************ Preferences ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map5 := make(map[string]any)
	err = yaml.Unmarshal(preferences.CRDYaml, map5)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for Preferences failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map5)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for Preferences: %s", err))
	}
	err = json.Unmarshal(b, &preferences.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for Preferences: %s", err))
	}

	err = clientSet.RegisterKind(ctx, preferences.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for Preferences failed to register: %s\n", err))
	}

	watcherWrapper5 := preferences.NewWatcherWrapper(r.preferencesWatcher)
	r.informerFactory.AddWatcher(preferences.CRD, watcherWrapper5)

	/************************ PublicDashboard ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map6 := make(map[string]any)
	err = yaml.Unmarshal(publicdashboard.CRDYaml, map6)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for PublicDashboard failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map6)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for PublicDashboard: %s", err))
	}
	err = json.Unmarshal(b, &publicdashboard.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for PublicDashboard: %s", err))
	}

	err = clientSet.RegisterKind(ctx, publicdashboard.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for PublicDashboard failed to register: %s\n", err))
	}

	watcherWrapper6 := publicdashboard.NewWatcherWrapper(r.publicdashboardWatcher)
	r.informerFactory.AddWatcher(publicdashboard.CRD, watcherWrapper6)

	/************************ ServiceAccount ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map7 := make(map[string]any)
	err = yaml.Unmarshal(serviceaccount.CRDYaml, map7)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for ServiceAccount failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map7)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for ServiceAccount: %s", err))
	}
	err = json.Unmarshal(b, &serviceaccount.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for ServiceAccount: %s", err))
	}

	err = clientSet.RegisterKind(ctx, serviceaccount.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for ServiceAccount failed to register: %s\n", err))
	}

	watcherWrapper7 := serviceaccount.NewWatcherWrapper(r.serviceaccountWatcher)
	r.informerFactory.AddWatcher(serviceaccount.CRD, watcherWrapper7)

	/************************ Team ************************/
	// TODO Having the committed form on disk in YAML is worth doing this for now...but fix this silliness
	map8 := make(map[string]any)
	err = yaml.Unmarshal(team.CRDYaml, map8)
	if err != nil {
		panic(fmt.Sprintf("generated CRD YAML for Team failed to unmarshal: %s", err))
	}
	b, err = json.Marshal(map8)
	if err != nil {
		panic(fmt.Sprintf("could not re-marshal CRD JSON for Team: %s", err))
	}
	err = json.Unmarshal(b, &team.CRD.Schema)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal CRD JSON for Team: %s", err))
	}

	err = clientSet.RegisterKind(ctx, team.CRD)
	if err != nil {
		panic(fmt.Sprintf("generated CRD for Team failed to register: %s\n", err))
	}

	watcherWrapper8 := team.NewWatcherWrapper(r.teamWatcher)
	r.informerFactory.AddWatcher(team.CRD, watcherWrapper8)

	return nil
}

func (r *Registry) run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}

// AccessPolicy returns the [crd.Kind] instance for the AccessPolicy kind.
func (r *Registry) AccessPolicy() crd.Kind {
	return r.all[0]
}

// Dashboard returns the [crd.Kind] instance for the Dashboard kind.
func (r *Registry) Dashboard() crd.Kind {
	return r.all[1]
}

// Folder returns the [crd.Kind] instance for the Folder kind.
func (r *Registry) Folder() crd.Kind {
	return r.all[2]
}

// LibraryPanel returns the [crd.Kind] instance for the LibraryPanel kind.
func (r *Registry) LibraryPanel() crd.Kind {
	return r.all[3]
}

// Playlist returns the [crd.Kind] instance for the Playlist kind.
func (r *Registry) Playlist() crd.Kind {
	return r.all[4]
}

// Preferences returns the [crd.Kind] instance for the Preferences kind.
func (r *Registry) Preferences() crd.Kind {
	return r.all[5]
}

// PublicDashboard returns the [crd.Kind] instance for the PublicDashboard kind.
func (r *Registry) PublicDashboard() crd.Kind {
	return r.all[6]
}

// ServiceAccount returns the [crd.Kind] instance for the ServiceAccount kind.
func (r *Registry) ServiceAccount() crd.Kind {
	return r.all[7]
}

// Team returns the [crd.Kind] instance for the Team kind.
func (r *Registry) Team() crd.Kind {
	return r.all[8]
}

func doNewRegistry(
	breg *corekind.Base,
	clientsetProvider client.ClientSetProvider,
	informerFactory informer.Informer,
	accesspolicyWatcher accesspolicy.Watcher,
	dashboardWatcher dashboard.Watcher,
	folderWatcher folder.Watcher,
	librarypanelWatcher librarypanel.Watcher,
	playlistWatcher playlist.Watcher,
	preferencesWatcher preferences.Watcher,
	publicdashboardWatcher publicdashboard.Watcher,
	serviceaccountWatcher serviceaccount.Watcher,
	teamWatcher team.Watcher,
) *Registry {
	reg := &Registry{}
	reg.clientsetProvider = clientsetProvider
	reg.informerFactory = informerFactory

	reg.accesspolicyWatcher = accesspolicyWatcher
	reg.all[0] = accesspolicy.CRD

	reg.dashboardWatcher = dashboardWatcher
	reg.all[1] = dashboard.CRD

	reg.folderWatcher = folderWatcher
	reg.all[2] = folder.CRD

	reg.librarypanelWatcher = librarypanelWatcher
	reg.all[3] = librarypanel.CRD

	reg.playlistWatcher = playlistWatcher
	reg.all[4] = playlist.CRD

	reg.preferencesWatcher = preferencesWatcher
	reg.all[5] = preferences.CRD

	reg.publicdashboardWatcher = publicdashboardWatcher
	reg.all[6] = publicdashboard.CRD

	reg.serviceaccountWatcher = serviceaccountWatcher
	reg.all[7] = serviceaccount.CRD

	reg.teamWatcher = teamWatcher
	reg.all[8] = team.CRD

	return reg
}
