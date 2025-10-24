package simple

type Database struct {
	Name string
}

type DatabaseMongoDB Database
type DatabaseMySQL Database

func NewDatabaseMySQL() *DatabaseMySQL {
	return (*DatabaseMySQL)(&Database{Name: "MySQL"})
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatabaseMySQL   *DatabaseMySQL
	DatabaseMongoDB *DatabaseMongoDB
}

func NewDatabaseRepository(mySQL *DatabaseMySQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabaseMySQL:   mySQL,
		DatabaseMongoDB: mongoDB,
	}
}
