package samplegobots

import (
	"github.com/Radisovik/gobot"
)

type Killer struct {
	*gobot.BaseBot
}

func (k Killer) Start() {
	k.DriveTowards(10, 100)
}

func (k Killer) Bb() *gobot.BaseBot {
	return k.BaseBot
}

func main() {
	gobot.Run(Killer{}, Killer{})
}
