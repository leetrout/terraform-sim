package store

import (
	"github.com/leetrout/terraform-sim/internal/resources"
	"github.com/schollz/jsonstore"
)

var Global = &State{}

func AddEntity(e *resources.Entity) {
	Global.Entities = append(Global.Entities, e)
	Global.Save()
}

func AddGroup(g *resources.Group) {
	Global.Groups = append(Global.Groups, g)
	Global.Save()
}

func Initialize() {
	Global.Load()
}

type State struct {
	Entities []*resources.Entity
	Groups   []*resources.Group
}

func (s *State) Save() {
	Persist(s)
}

func (s *State) Load() {
	loaded := Load()
	s.Entities = loaded.Entities
	s.Groups = loaded.Groups
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
	if err != nil {
		panic(err)
	}
	return &s
}
