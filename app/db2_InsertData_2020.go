package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db2_InsertData_2020(db *sql.DB) {

	stmt, err := db.Prepare(
		`INSERT INTO univ_susi_info_2021 (id, suin_cd, recruit_volume_2021, compete_ratio_2021, recruit_replace_2021, nasin_best_2021, nasin_mean_2021, nasin_limit_2021, recruit_factor_mean_2021, inDate) VALUES
		(1, 'univ_109_001', 6, '3.8:1', 4, 2.55, 2.93, 3.16, 0, '2021-08-25 08:18:46'),
		
		;`)
	utils.HandleErr(err, "create table1")
	_, err = stmt.Exec()
	utils.HandleErr(err, "create table11 exec")

	stmt, err = db.Prepare(
		`INSERT INTO univ_susi_info_2021 (id, suin_cd, recruit_volume_2021, compete_ratio_2021, recruit_replace_2021, nasin_best_2021, nasin_mean_2021, nasin_limit_2021, recruit_factor_mean_2021, inDate) VALUES
	;`)
	utils.HandleErr(err, "create table2")
	_, err = stmt.Exec()
	utils.HandleErr(err, "create table12 exec")

	stmt, err = db.Prepare(
		`INSERT INTO univ_susi_info_2021 (id, suin_cd, recruit_volume_2021, compete_ratio_2021, recruit_replace_2021, nasin_best_2021, nasin_mean_2021, nasin_limit_2021, recruit_factor_mean_2021, inDate) VALUES
	
		;`)
	utils.HandleErr(err, "create table3")
	_, err = stmt.Exec()
	utils.HandleErr(err, "create table13 exec")

	stmt, err = db.Prepare(
		`INSERT INTO univ_susi_info_2021 (id, suin_cd, recruit_volume_2021, compete_ratio_2021, recruit_replace_2021, nasin_best_2021, nasin_mean_2021, nasin_limit_2021, recruit_factor_mean_2021, inDate) VALUES
	
		;`)
	utils.HandleErr(err, "create table4")
	_, err = stmt.Exec()
	utils.HandleErr(err, "create table14 exec")
}
