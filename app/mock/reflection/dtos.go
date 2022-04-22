package reflection

import "reflect"

type PropertyKind int64

const (
	single PropertyKind = iota
	slice
)

func (pk PropertyKind) String() string {
	return [...]string{"", "single", "slice"}[pk]
}

type PropertyPattern struct {
	Reflect         reflect.Kind `json:"-"`
	ReflectionLabel string       `json:"type"`
	Kind            PropertyKind `json:"kind"`
	Len             int          `json:"length, omitempty"`
}
