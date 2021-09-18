package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func dbForDataUnvi(db *sql.DB) {
			
	statement, err := db.Prepare(
		`CREATE TABLE univ_susi_info_fst (
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
			inDate TIMESTAMP
		);`)

	utils.HandleErr(err, "create table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table11 exec")

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

	statement, err = db.Prepare(
		`CREATE TABLE univ_susi_info_2021 (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			recruit_volume_2021 float,
			compete_ratio_2021 varchar(16),
			recruit_replace_2021 integer,
			nasin_best_2021 float,
			nasin_mean_2021 float,
			nasin_limit_2021 float,
			recruit_factor_mean_2021 float,
			inDate TIMESTAMP
		);`)
	utils.HandleErr(err, "create table3")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table13 exec")

	statement, err = db.Prepare(
		`CREATE TABLE univ_susi_info_2020 (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			recruit_volume_2020 float,
			compete_ratio_2020 varchar(16),
			recruit_replace_2020 integer,
			nasin_best_2020 float,
			nasin_mean_2020 float,
			nasin_limit_2020 float,
			recruit_factor_mean_2020 float,
			inDate TIMESTAMP
		);`)
	utils.HandleErr(err, "create table4")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table14 exec")

	statement, err = db.Prepare(
		`CREATE TABLE univ_susi_info_2019 (
			id SERIAL PRIMARY KEY,
			suin_cd varchar(16) NOT NULL,
			recruit_volume_2019 float,
			compete_ratio_2019 varchar(16),
			recruit_replace_2019 integer,
			nasin_best_2019 float,
			nasin_mean_2019 float,
			nasin_limit_2019 float,
			recruit_factor_mean_2019 float,
			inDate TIMESTAMP
		);`)
	utils.HandleErr(err, "create table4")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table14 exec")


}
