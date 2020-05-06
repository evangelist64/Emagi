package mysql

import "database/sql"

type DBManager struct {
	conns   []*DB
	connMap map[string]*DB //简单处理，表名=>连接，同个表的query由同个连接操作
}

func (p *DBManager) Init() bool {
	//连接初始化
	return true
}

func (p *DBManager) Query(tableName string, sql string, cb func(*sql.Rows)) {
	if db, exist := p.connMap[tableName]; exist {
		db.Query(sql, cb)
	}
}
