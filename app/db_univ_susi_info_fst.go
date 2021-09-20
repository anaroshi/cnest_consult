package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_univ_susi_info_fst(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_fst;`)
	utils.HandleErr(err, "Drop univ_susi_info_fst table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop univ_susi_info_fst table")

	// Create Table
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS univ_susi_info_fst (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			univ_id varchar(16) NOT NULL,
			apply_term_id varchar(16),
			apply_form varchar(32),
			apply_formName varchar(32),
			apply_dept_id varchar(16),
			apply_line_id varchar(16),
			apply_sub_id varchar(16),
			recruit_volume integer,
			apply_quali_id varchar(128),
			recruit_met_id varchar(128),
			recruit_su_lim_id varchar(356),
			nasin_met_id varchar(356),
			docu_met_id varchar(128),
			nonsul_view_met_id varchar(128),
			nonsul_met_id varchar(128),
			etc_met_id varchar(128),
			docu_list_id varchar(128),
			auto_writing_id varchar(128),
			inDate TIMESTAMP );
		`)

	utils.HandleErr(err, "Create univ_susi_info_fst table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create univ_susi_info_fst table")
}
