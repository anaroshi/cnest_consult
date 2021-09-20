package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_univ_susi_info_2021(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_2021;`)
	utils.HandleErr(err, "Drop univ_susi_info_2021 table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop univ_susi_info_2021 table")

	// Create Table
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS univ_susi_info_2021 (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			recruit_volume_2021 float,
			compete_ratio_2021 varchar(16),
			recruit_replace_2021 integer,
			nasin_best_2021 float,
			nasin_mean_2021 float,
			nasin_limit_2021 float,
			recruit_factor_mean_2021 float,
			inDate TIMESTAMP );
		`)

	utils.HandleErr(err, "Create univ_susi_info_2021 table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create univ_susi_info_2021 table")

	//  Insert Data
	stmt, err = db.Prepare(`
		INSERT INTO univ_susi_info_2021 (id, formId, formId, sort, flag) VALUES
		(1, 'form_101', '학생부교과', 0, 0),
		(2, 'form_102', '학생부종합', 0, 0),
		(3, 'form_103', '논술', 0, 0),
		(4, 'form_104', '실기', 0, 0);
	`)
	
	utils.HandleErr(err, "Insert Data into univ_susi_info_2021 table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Insert Data into univ_susi_info_2021 table")
}
