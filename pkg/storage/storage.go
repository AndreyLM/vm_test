package storage

import "github.com/andreylm/vm_test/pkg/models"

// IStorage - storage interface
type IStorage interface {
	Load(jsonFile string) error
	Add(u *models.User) error
	Get(ID int) (*models.User, error)
	GetAll() ([]*models.User, error)
	Close()
}

// GetStorage - gets storage
func GetStorage(useDB bool, conn string) (IStorage, error) {
	if !useDB {
		return CreateCache(), nil
	}
	return CreateConnection(conn)
}
