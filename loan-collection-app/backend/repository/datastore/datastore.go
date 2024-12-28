package datastore

import (
	"fmt"
	"os"
)

func Get() {
	if PostgeSQLConn == nil {
		err := ConnectPostgeSQL()
		if err != nil {
			fmt.Println("Error connecting to MySQL:", err)
			os.Exit(1)
		}
	}
	if RateLimiter == nil {
		InitializeRateLimiter()
	}
}
