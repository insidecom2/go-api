package configs

const (
	DBHOST     = "127.0.0.1"
	DBPORT     = "5432"
	DBNAME     = "db_go"
	DBUSER     = "root"
	DBPASSWORD = "root"
	DBTYPE     = "postgres"
)

func GetDBType() string {
	return DBTYPE
}
