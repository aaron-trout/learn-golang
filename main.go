package main

import (
	"fmt"
	"time"
)

func dayOfWeek() time.Weekday {
	return time.Now().Weekday()
}

func main() {
	day := dayOfWeek()
	fmt.Printf("Happy %s!", day)
}
