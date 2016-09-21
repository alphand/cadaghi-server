package db

//IDataStore - DataStore interface for DB Ops
type IDataStore interface {
	Create(interface{}) error
	GetByID(id string) (interface{}, error)
	GetAll(interface{}, []interface{}) error
	// Put(id string) (interface{}, error)
	// Delete(id string, soft bool)
}
