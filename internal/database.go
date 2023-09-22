package internal

type Database struct {
}

func (db *Database) Initialize() (bool, error) {
	return true, nil
}
