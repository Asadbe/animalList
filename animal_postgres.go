package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "asadbek"
	password = "1"
	dbname   = "mydb"
)

type sqlxDB struct {
	connectDB *sqlx.DB
}

var err error

//NewAnimalList ...
func NewAnimalList() (AnimalListint, error) {
	cm := sqlxDB{}
	psqlInfo := fmt.Sprintf(` user=%s dbname=%s password=%s`, user, dbname, password)
	cm.connectDB, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &cm, nil
}
func (s *sqlxDB) Add(a AnimalList) error {
	insertionQuery := `insert into animals (name, voice, leg) values ($1, $2, $3);`

	_, err := s.connectDB.Exec(insertionQuery, a.name, a.voice, a.leg)

	if err != nil {
		return err
	}

	return nil
}
func (s *sqlxDB) Update(name string, voice string, leg int, id int) error {
	updatingQuery := `update animals set name=$1,voice=$2, leg=$3
	where id =$4`

	_, err := s.connectDB.Exec(updatingQuery, name, voice, leg, id)

	if err != nil {
		fmt.Println("Can't update")
		return err
	}

	return nil
}
func (s *sqlxDB) Delete(id int) error {
	deletingQuery := `delete from animals where id=$1;`

	_, err := s.connectDB.Exec(deletingQuery, id)

	if err != nil {
		fmt.Println("Can't delete")
		return err
	}
	return nil
}
func (s *sqlxDB) GetTask(id int) error {
	listTaskQuery := `select id,name,voice,leg from animals where id=$1;`

	rows, err := s.connectDB.Queryx(listTaskQuery, id)

	if err != nil {
		fmt.Println("Can't print task list")
		return err
	}
	var ts AnimalList
	for rows.Next() {
		err = rows.Scan(&ts.id, &ts.name, &ts.voice, &ts.leg)
		if err != nil {
			fmt.Println("Can't scan struct")
			return err
		}
	}
	fmt.Println(ts)
	return nil

}
