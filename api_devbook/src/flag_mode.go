package flag_mode

import (
	"api/src/fakedata"
	"flag"
	"fmt"
)

func ProdMode() bool {
	prodMode := flag.Bool("prod", false, "Run in production mode")
	fakeData := flag.Bool("fake", false, "Insert fake data")

	flag.Parse()
	if *prodMode {
		fmt.Println("Running in production mode")
	} else {
		fmt.Println("Running in development mode")
		if *fakeData {
			fmt.Println("Inserting fake data")
			fakedata.FakeData()
		}
	}
	return *prodMode

}
