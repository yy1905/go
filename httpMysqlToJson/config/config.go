package config

type Conf struct {
	Mysql mysql
}

type mysql struct {
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
}

type QueryInfo struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
