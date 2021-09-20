package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_univ_susi_info_snd(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_snd;`)
	utils.HandleErr(err, "Drop univ_susi_info_snd table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop univ_susi_info_snd table")

	// Create Table
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS univ_susi_info_snd (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			remarks varchar(256) NOT NULL,
			apply_datetime varchar(128),
			apply_docu_datetime varchar(128),
			recruit_select_date varchar(128),
			recruit_test_date varchar(128),
			recruit_pass_date varchar(128),
			result_remarks varchar(256),
			inDate TIMESTAMP );
		`)

	utils.HandleErr(err, "Create univ_susi_info_snd table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create univ_susi_info_snd table")
}
