package stdout

import (
	"fmt"

	"github.com/JSila/packtFree/info"
	"github.com/JSila/packtFree/notify"
)

type Stdout struct {
	data *info.Info
}

func (s *Stdout) Notify(options *notify.User) error {
	fmt.Printf("Today's free Packt Publishing e-book: '%s'", s.data.Title)
	return nil
}

func (s *Stdout) Name() string {
	return "stdout"
}

func New(data *info.Info) *Stdout {
	return &Stdout{
		data: data,
	}
}
