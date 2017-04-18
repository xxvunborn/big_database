package schemes

// PersonService represents a service that handles all database actions
// relateds with person's table
type PersonService struct {
	*Datastore
}

// Person represents the persons1/2 database table
type Person struct {
	Rut     int    `db:"rut"`
	Name    string `db:"name"`
	Age     int    `db:"age"`
	Address string `db:"address"`
}

// Persons return all the matches with table persons1/2 given by the query
func (ps *PersonService) Persons(query string) ([]*Person, error) {
	var pp = make([]*Person, 0, 0)
	if err := ps.SQL(query).QueryStructs(&pp); err != nil {
		return nil, err
	}

	return pp, nil
}
