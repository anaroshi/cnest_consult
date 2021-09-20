package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_univ_susi_info_2019(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_2019;`)
	utils.HandleErr(err, "Drop univ_susi_info_2019 table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop univ_susi_info_2019 table")

	// Create Table
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS univ_susi_info_2019 (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			recruit_volume_2019 float,
			compete_ratio_2019 varchar(16),
			recruit_replace_2019 integer,
			nasin_best_2019 float,
			nasin_mean_2019 float,
			nasin_limit_2019 float,
			recruit_factor_mean_2019 float,
			inDate TIMESTAMP );
		`)

	utils.HandleErr(err, "Create univ_susi_info_2019 table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create univ_susi_info_2019 table")

}
