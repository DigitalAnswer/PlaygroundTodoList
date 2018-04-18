package main

type dbSettings struct {
	ip     string
	port   string
	dbName string
}

var dbDevSettings = dbSettings{
	ip:     "localhost",
	port:   "6603",
	dbName: "TodoDev",
}

func (s *dbSettings) dataSource() string {
	return "root:admin@tcp(" + s.ip + ":" + s.port + ")/" + s.dbName
}
