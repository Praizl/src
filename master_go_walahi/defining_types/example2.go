package main

import "fmt"

type duration int

func main() {
	var hour duration = 3600
	minute := duration(60)

	fmt.Println(hour != minute)
	musicApps := [4]string{"apple music", "spotify", "youtube", "tidal"}
	aroSubs := [1]string{"spotify"}
label:
	for index, music := range musicApps {
		for _, aroSubs := range aroSubs {
			if music == aroSubs {
				fmt.Printf("Aro uses %q no %d out of all the music apps", music, index)
				break label
			}

		}
	}
	fmt.Println("")
	//goto
	var d int = 1
weOutside:
	if d < 10 {
		fmt.Println(d)
	}
	d++
	goto weOutside
}
