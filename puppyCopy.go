package puppycopy

import (
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
