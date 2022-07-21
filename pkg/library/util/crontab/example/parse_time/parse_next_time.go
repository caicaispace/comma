package main

import (
	"comma/pkg/library/util/crontab"
	"fmt"
	"time"
)

func main() {
	t, err := crontab.Parse("0 0 0 */1 * *")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(" now: ", time.Now())
	next := t.Next(time.Now())
	fmt.Println("next: ", next)
}
