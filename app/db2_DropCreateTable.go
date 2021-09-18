package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db2_DropCreateTable(db *sql.DB) {
	// Drop TABLE univ_susi_info_snd
	statement, err := db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_snd;`)

	utils.HandleErr(err, "drop table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table11 exec")


	// CREATE TABLE univ_susi_info_snd
	statement, err = db.Prepare(
		`CREATE TABLE univ_susi_info_snd (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			remarks varchar(256) NOT NULL,
			apply_datetime varchar(128),
			apply_docu_datetime varchar(128),
			recruit_select_date varchar(128),
			recruit_test_date varchar(128),
			recruit_pass_date varchar(128),
			result_remarks varchar(256),
			inDate TIMESTAMP
		);`)
	utils.HandleErr(err, "create table2")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table12 exec")


}
