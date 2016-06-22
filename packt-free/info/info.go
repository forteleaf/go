package info

import (
	"fmt"
	"time"
)

// Info represent informations about free offer
type Info struct {
	Title, Img  string
	Description []string
	EndTime     time.Time
}

// String formats string representation of Info
func (i Info) String() string {
	return fmt.Sprintf("Info about '%s'", i.Title)
}
