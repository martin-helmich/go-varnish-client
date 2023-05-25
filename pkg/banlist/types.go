package banlist

import "time"

type BanList []Ban

const (
	BanActive BanStatus = iota
	BanGone
	BanComplete
)

type BanStatus int

func (v BanStatus) String() string {
	return [...]string{"active", "gone", "complete"}[v]
}

type Ban struct {
	Time    time.Time
	Objects int64
	Status  BanStatus
	Spec    string
}

func (b Ban) Equals(ban Ban) bool {
	if b.Spec != ban.Spec || b.Objects != ban.Objects || b.Status != ban.Status || b.Time != ban.Time {
		return false
	}
	return true
}
