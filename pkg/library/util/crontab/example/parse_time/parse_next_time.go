package main

import (
	"fmt"
	"time"

	"goaway/pkg/library/util/crontab"
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
