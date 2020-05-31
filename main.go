package app

import (
	"time"

	mars "github.com/emxqm-11/go-mars-rover/pkg/marsrover"
)

func main() {
	r := mars.NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
}
