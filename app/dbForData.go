package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func dbForData(db *sql.DB) {
			
	statement, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS apply_subject (
			id SERIAL PRIMARY KEY,
			apSubjectId VARCHAR(16) NOT NULL ,
			apSubjectName VARCHAR(32) NOT NULL,
			sort integer,
			flag integer,
			user VARCHAR(24),
			inDate TIMESTAMP,
			reUser VARCHAR(24),
			reDate TIMESTAMP,
		);`)

	utils.HandleErr(err, "create table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table11 exec")

	statement, err = db.Prepare(
		`CREATE TABLE apply_line (
			id SERIAL PRIMARY KEY,
			lineId varchar(16) NOT NULL,
			lineName varchar(32) NOT NULL,
		);`)
	utils.HandleErr(err, "create table2")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table12 exec")

	statement, err = db.Prepare(
		`CREATE TABLE apply_form (
			id SERIAL PRIMARY KEY,
			formId varchar(16) NOT NULL,
			formName varchar(32) NOT NULL,
		);`)
	utils.HandleErr(err, "create table3")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table13 exec")

	statement, err = db.Prepare(
		`CREATE TABLE apply_dept (
			id SERIAL PRIMARY KEY,
			deptId varchar(16) NOT NULL,
			deptName varchar(32) NOT NULL,
			sort integer,
			flag iinteger,
			user varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP,
		);`)
	utils.HandleErr(err, "create table4")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table14 exec")


}
