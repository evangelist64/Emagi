package mysql

import (
	"Emagi/log"
	"database/sql"
	"fmt"
)

type DB struct {
	conn *sql.DB     //mysql连接
	ch   chan DBTask //需要执行的sql
}

type DBTask struct {
	sql string
	cb  func(*sql.Rows)
}

func (p *DB) Query(sql string, cb func(*sql.Rows)) {
	p.ch <- DBTask{sql: sql, cb: cb}
}

func (p *DB) Run() {
	for {
		task := <-p.ch
		rows, err := p.conn.Query(task.sql)
		if err != nil {
			log.Error(fmt.Sprintf("run sql failed, sql:%s", task.sql))
		}
		if task.cb != nil {
			task.cb(rows)
		}
	}
}
