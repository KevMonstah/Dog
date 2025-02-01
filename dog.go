package Dog

import (
	"log"
)

func Sleeps(s string) string {
	log.Println(s + " the dog sleeps")
	return "My Dog sleeps"
}

func Runs(s string) string {
	log.Println(s + " the dog runs")
	return "My Dog Runs"
}

func Chills(s string) string {
	log.Println(s + "the dog is chilling")
	return "My Dog Chills"
}
