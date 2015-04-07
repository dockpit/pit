package model

type Milestone struct {
	ID         string            `json:"id"`
	IsUnlocked func(*Stats) bool `json:"-"`
	Shown      bool              `json:"shown"`
}

var (
	FirstDep = Milestone{"first_dep", func(s *Stats) bool {
		return s.NrOfDepsCreated > 0
	}, false}
)

func Milestones() []Milestone {
	return []Milestone{
		FirstDep,
	}
}
