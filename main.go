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
	query := `select age from persons2 group by age`
	start := time.Now()
	_, err = ps.Persons(query)
	if err != nil {
		_ = fmt.Errorf("error %+v\n", err)
	}

	//groupByAge(pp)

	elapsed := time.Since(start)
	fmt.Println("query time: ", elapsed)
}

func groupByAge(pp []*schemes.Person) {
	group := make([]int, 100, 100)
	for _, p := range pp {
		group[p.Age]++
	}
}
