package movie

import "fmt"

func Review(name string, rating float64) {
	fmt.Printf("I reviewed %s and it's rating is %f !!!\n", name, rating)
}

func FindName(imdbID string) string {
	switch imdbID {
	case "tt4154769":
		return "Avergers: Endgame"
	case "tt1825683":
		return "Black Panther"
	}
	return "Not Found!"
}
