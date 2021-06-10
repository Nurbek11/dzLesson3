package dz3Lesson

import (
	"log"
)

func RuneByIndex(s *string, i *int) (rune, error) {
	runes := []rune(*s)
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered: %v", r)
		}
	}()
	return runes[*i], nil

}
