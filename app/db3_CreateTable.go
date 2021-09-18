package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db3_CreateTable(db *sql.DB) {
			
	statement, err := db.Prepare(
		`CREATE TABLE univ_info (
			id SERIAL PRIMARY KEY,
			univId varchar(16) NOT NULL,
			univName varchar(32) NOT NULL,
			univFullName varchar(64),
			location varchar(32),
			address varchar(128),
			gender varchar(14),
			phone varchar(16),
			fax varchar(16),
			email varchar(32),
			website varchar(32),
			remark varchar(128),
			inUser varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP
		);`)

	utils.HandleErr(err, "create table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table11 exec")

	statement, err = db.Prepare(
		`CREATE TABLE recruit_method (
			id SERIAL PRIMARY KEY,
			recruitId varchar(16) NOT NULL,
			recruitName varchar(16) NOT NULL,
			recruit1st varchar(120),
			recruit1stVol varchar(120),
			recruitFinal varchar(256),
			sort integer,
			flag integer,
			inUser varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP
		);`)
	utils.HandleErr(err, "create table2")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table12 exec")

	statement, err = db.Prepare(
		`CREATE TABLE recruit_sununglimit (
			id SERIAL PRIMARY KEY,
			sulimitId varchar(16) NOT NULL,
			sulimitName varchar(500) NOT NULL,
			sort integer,
			flag integer,
			inUser varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP
		);`)
	utils.HandleErr(err, "create table3")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table13 exec")
}
