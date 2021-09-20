package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_univ_susi_info_2020(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_2020;`)
	utils.HandleErr(err, "Drop univ_susi_info_2020 table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop univ_susi_info_2020 table")

	// Create Table
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS univ_susi_info_2020 (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			recruit_volume_2020 float,
			compete_ratio_2020 varchar(16),
			recruit_replace_2020 integer,
			nasin_best_2020 float,
			nasin_mean_2020 float,
			nasin_limit_2020 float,
			recruit_factor_mean_2020 float,
			inDate TIMESTAMP );
		`)

	utils.HandleErr(err, "Create univ_susi_info_2020 table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create univ_susi_info_2020 table")

}
