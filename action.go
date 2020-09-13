package gardener

import "github.com/bitfield/script"

type Act func() error

type Script struct {
	Directory string
	Action    []Act
}

func (s *Script) Run() error {
	p := script.NewPipe()
	return p.Error()
}
