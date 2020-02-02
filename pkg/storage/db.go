package storage

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/andreylm/vm_test/pkg/ent"
	"github.com/andreylm/vm_test/pkg/ent/user"
	"github.com/andreylm/vm_test/pkg/models"

	// sqllite dialect
	_ "github.com/mattn/go-sqlite3"
)

// DB - db storage
type DB struct {
	cli *ent.Client
}

// CreateConnection - creates connection
func CreateConnection(conn string) (*DB, error) {
	client, err := ent.Open("sqlite3", conn)
	if err != nil {
		return nil, err
	}
	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		client.Close()
		return nil, err
	}

	return &DB{cli: client}, nil
}

// Load - loads data from jsonFile
func (db *DB) Load(filePath string) error {
	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	users := []*models.User{}
	if err := json.Unmarshal(jsonFile, &users); err != nil {
		log.Println(err)
	}

	for _, u := range users {
		db.Add(u)
	}

	return nil
}

// Add - adds user
func (db *DB) Add(u *models.User) error {
	uCreate := db.cli.User.Create()
	uCreate.SetID(int64(u.ID)).
		SetUserName(u.UserName).
		SetCity(u.City).
		SetBirthDate(u.BirthDate.Time).
		SetDepartment(u.Department).
		SetExperienceYears(u.ExperienceYears).
		SetFullName(u.FullName).
		SetGender(u.Gender)
	_, err := uCreate.Save(context.TODO())
	return err
}

// GetAll gets all user
func (db *DB) GetAll() ([]*models.User, error) {
	res, err := db.cli.User.Query().All(context.TODO())
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, 0, len(res))
	for _, v := range res {
		users = append(users, db.mapEntToModel(v))
	}

	return users, nil
}

// Get - gets user by id
func (db *DB) Get(ID int) (*models.User, error) {
	u, err := db.cli.User.Query().Where(user.IDEQ(int64(ID))).Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return db.mapEntToModel(u), nil
}

func (db *DB) mapEntToModel(u *ent.User) *models.User {
	bDay := models.CustomTime{}
	bDay.Time = u.BirthDate
	return &models.User{
		ID:              int(u.ID),
		FullName:        u.FullName,
		Gender:          u.Gender,
		ExperienceYears: u.ExperienceYears,
		UserName:        u.UserName,
		BirthDate:       bDay,
		City:            u.City,
		Department:      u.Department,
	}
}

// Close - closes connection
func (db *DB) Close() {
	db.cli.Close()
}
