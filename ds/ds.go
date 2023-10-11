package ds

type DataSource interface {
	Ping() error
	Close() error
	ListDBs() ([]string, error)
	ListTables(db string) ([]string, error)
	ListRecords(table string) (data [][]*string, err error)
	ListSchemas(table string) (data [][]*string, err error)
	CustomQuery(query string) (data [][]*string, err error)
}
