package main

import (
	"context"
	"fmt"
	"github.com/pindarEng/stockSimulator/api"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006-01-02"))

	app := api.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
