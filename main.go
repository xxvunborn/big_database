package main

import (
	"fmt"
	"time"

	"github.com/mtavano/big_db/schemes"
)

const (
	postgresDSN = "postgres://localhost/big_database?sslmode=disable"
)

func main() {
	fmt.Println("big databases task 5")
	ds, err := schemes.NewDatastore(postgresDSN)
	if err != nil {
		_ = fmt.Errorf("error %+v\n", err)
	}

	fmt.Println("database connection established")

	ps := &schemes.PersonService{Datastore: ds}
	query := `select * from persons1`
	start := time.Now()
	pp, err := ps.Persons(query)
	if err != nil {
		_ = fmt.Errorf("error %+v\n", err)
	}

	groupByAge(pp)

	elapsed := time.Since(start)
	fmt.Printf("query time: %d\n", elapsed)
}

func groupByAge(pp []*schemes.Person) {
	group := make([]int, 100, 100)
	for _, p := range pp {
		group[p.Age]++
	}
}
