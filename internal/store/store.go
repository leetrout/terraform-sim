package store

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/leetrout/terraform-sim/internal/resources"
	"github.com/leetrout/terraform-sim/internal/ws"
	"github.com/schollz/jsonstore"
)

func NotifyClients(msg string) {
	ws.SOCKEX.Send(msg)
}

var Global = &State{
	Entities: map[string]*resources.Entity{},
	Groups:   map[string]*resources.Group{},
}

const StateFile = "./tfsim.json.gz"

func AddEntity(e *resources.Entity) {
	Global.Entities[e.ID] = e
	Global.Save()
	NotifyClients(fmt.Sprintf("ENTITY:ADDED:%s", e.ID))
}

func RemoveEntity(e *resources.Entity) {
	delete(Global.Entities, e.ID)
	Global.Save()
	NotifyClients(fmt.Sprintf("ENTITY:REMOVED:%s", e.ID))
}

func AddGroup(g *resources.Group) {
	Global.Groups[g.ID] = g
	Global.Save()
	NotifyClients(fmt.Sprintf("GROUP:ADDED:%s", g.ID))
}

func RemoveGroup(g *resources.Group) {
	delete(Global.Groups, g.ID)
	Global.Save()
	NotifyClients(fmt.Sprintf("GROUP:REMOVED:%s", g.ID))
}

func Initialize() {
	_, err := os.Stat(StateFile)
	if err == nil {
		Global.Load()
		return
	}
	if errors.Is(err, os.ErrNotExist) {
		Global.Save()
		return
	}
	panic("Could not initialize storage")
}

type State struct {
	Entities map[string]*resources.Entity
	Groups   map[string]*resources.Group
}

func (s *State) Save() {
	Persist(s)
}

func (s *State) Load() {
	loaded := Load()
	s.Entities = loaded.Entities
	s.Groups = loaded.Groups
}

// EntityList returns a list of entities instead of the mapping.
func (s *State) EntityList() []*resources.Entity {
	cnt := len(Global.Entities)
	entityList := make([]*resources.Entity, cnt)
	idx := 0
	for _, entity := range Global.Entities {
		entityList[idx] = entity
		idx += 1
	}

	sort.SliceStable(entityList, func(i int, j int) bool {
		return entityList[i].ID < entityList[j].ID
	})

	return entityList
}

// GroupList returns a list of groups instead of the mapping.
func (s *State) GroupList() []*resources.Group {
	cnt := len(Global.Groups)
	groupList := make([]*resources.Group, cnt)
	idx := 0
	for _, group := range Global.Groups {
		groupList[idx] = group
		idx += 1
	}

	sort.SliceStable(groupList, func(i int, j int) bool {
		return groupList[i].ID < groupList[j].ID
	})

	return groupList
}

func Persist(s *State) {
	ks := new(jsonstore.JSONStore)
	err := ks.Set("state", s)
	if err != nil {
		panic(err)
	}

	if err = jsonstore.Save(ks, "tfsim.json.gz"); err != nil {
		panic(err)
	}
}

func Load() *State {
	// TODO Check existance first
	ks, err := jsonstore.Open("tfsim.json.gz")
	if err != nil {
		panic(err)
	}

	// get the data back via an interface
	var s State
	err = ks.Get("state", &s)
	if s.Entities == nil {
		s.Entities = make(map[string]*resources.Entity)
	}
	if s.Groups == nil {
		s.Groups = make(map[string]*resources.Group)
	}
	if err != nil {
		panic(err)
	}
	return &s
}
