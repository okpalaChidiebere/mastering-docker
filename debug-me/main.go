package main

import (
	"errors"
	"log"
	"os"
	"time"
)

var food = os.Getenv("FAVORITE_FOOD")


func favoriteFood(f string) error {
	if f != ""{
		log.Printf("My favorite food is %s", f)
		return nil
	}
	return errors.New("favorite food is undefined")
}

func sleep(ms int){
	time.Sleep(time.Duration(ms)*time.Millisecond)
}

	
func main() {

	for {
		err := favoriteFood(food)
		if err != nil {
			log.Println(err.Error())
		}
		sleep(5000)
	}
}