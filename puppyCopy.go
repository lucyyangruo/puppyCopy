package puppycopy

import (
	"fmt"

	dogcopy "github.com/lucyyangruo/dogCopy"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark(s string) string {
	return dogcopy.WhenGrowUp(Bark())
}
func BigBarks(s string) string {
	return dogcopy.WhenGrowUp(Barks())
}

func showVersion() {
	fmt.Println("I'm form version v1.0.0")
}

func showVersionAgain() {
	fmt.Println("I'm form version v1.2.0")
}

func showVersionAgainAndAgain() {
	fmt.Println("I'm form version v1.3.0")
}

func ShowEver() {
	showVersion()
	showVersionAgain()
	showVersionAgainAndAgain()
}
