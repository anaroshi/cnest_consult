package app

import (
	"cnest_consult/univ/utils"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// 대학정보 whole
func GetDetailUniv(db *sql.DB, id string) items {

	// 대학정보1 Info
	var info1  detailUniv1
	info1 = univInfo1(info1, id)	
	suin_cd := info1.Suin_cd // 대학학과 cd	

	// 대학정보2 Info
	var info2  detailUniv2
	info2 = univInfo2(info2, suin_cd)

	// 2021학년도 Info
	var n21 value2021
	n21 = n21Info(n21, suin_cd)

	
	// 2020학년도 Info
	var n20 value2020
	n20 = n20Info(n20, suin_cd)


	// 2019학년도 Info
	var n19 value2019
	n19 = n19Info(n19, suin_cd)
		
	unviInfo := items{Info_1:info1, Info_2:info2, N_21:n21, N_20:n20, N_19:n19, }
	
	return unviInfo
}

// 대학정보1 Info
func univInfo1(info1 detailUniv1, id string) detailUniv1 {

	qry := `SELECT 
		u1.id, u1.suin_cd, 
		u2.univName, 
		u1.apply_term_id, u1.apply_form, u1.apply_formName, u1.apply_dept_id,	u1.apply_line_id, 
		u5.apSubjectName, 
		u1.recruit_volume, 
		u6.qualiName, 
		u3.recruitName, u3.recruit1st, u3.recruit1stVol, u3.recruitFinal, 
		u4.sulimitName, 
		u7.nasinmetName, 
		u8.docmetName, 
		u9.intmetName, 
		u10.nslmetName, 
		u11.etcmetName, 
		u12.doclstName, 
		u13.atwrtName 
		FROM univ_susi_info_fst u1 
		LEFT JOIN univ_info u2 ON u1.univ_id = u2.univId 
		LEFT JOIN recruit_method u3 ON u1.recruit_met_id=u3.recruitId 
		LEFT JOIN recruit_sununglimit u4 ON u1.recruit_su_lim_id = u4.sulimitId
		LEFT JOIN apply_subject u5 ON u1.apply_sub_id = u5.apSubjectId
		LEFT JOIN apply_qualification u6 ON u1.apply_quali_id = u6.qualiId
		LEFT JOIN nasin_method u7 ON u1.nasin_met_id = u7.nasinmetId
		LEFT JOIN document_method u8 ON u1.docu_met_id = u8.docmetId
		LEFT JOIN interview_method u9 ON u1.Nonsul_view_met_id = u9.intmetId
		LEFT JOIN nonsul_method u10 ON u1.nonsul_met_id = u10.nslmetId
		LEFT JOIN etc_method u11 ON u1.etc_met_id = u11.etcmetId
		LEFT JOIN document_list u12 ON u1.docu_list_id = u12.doclstId
		LEFT JOIN autowriting u13 ON u1.auto_writing_id = u13.atwrtId
		WHERE u1.id=? `

	rows, err := db.Query(qry, id)
	utils.HandleErr(err, "printUnivInfo_dbselect")
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&info1.Id,
			&info1.Suin_cd,
			&info1.UnivName,
			&info1.Apply_term_id,
			&info1.Apply_form,
			&info1.Apply_formName,
			&info1.Apply_dept_id,
			&info1.Apply_line_id,
			&info1.ApSubjectName,
			&info1.Recruit_volume,
			&info1.QualiName,
			&info1.RecruitName,
			&info1.Recruit1st,
			&info1.Recruit1stVol,
			&info1.RecruitFinal,
			&info1.SulimitName,
			&info1.NasinmetName,
			&info1.DocmetName,
			&info1.IntmetName,
			&info1.NslmetName,
			&info1.EtcmetName,
			&info1.DoclstName,
			&info1.AtwrtName,
		)
		utils.HandleErr(err, "printUnivInfo_loadDB")
	}
	return info1
}

// 대학정보2 Info
func univInfo2(info2 detailUniv2, suin_cd string) detailUniv2 {
	qry := `SELECT remarks, apply_datetime, apply_docu_datetime, recruit_select_date, recruit_test_date, 
	recruit_pass_date, result_remarks FROM univ_susi_info_snd WHERE suin_cd = ?`
	rows, err := db.Query(qry, suin_cd)
	utils.HandleErr(err, "dbselect2")
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan( 
			&info2.Remarks, &info2.ApplyDatetime, &info2.ApplyDocuDatetime, &info2.RecruitSelectDate, 
			&info2.RecruitTestDate, &info2.RecruitPassDate, &info2.ResultRemarks,
		)
		utils.HandleErr(err, "load detailUniv2 Value")
	}
	return info2
}


// 2021학년도 Info
func n21Info(n21 value2021, suin_cd string) value2021 {
	
	qry := `SELECT recruit_volume_2021, compete_ratio_2021, recruit_replace_2021, nasin_best_2021, nasin_mean_2021, 
	nasin_limit_2021, recruit_factor_mean_2021 FROM univ_susi_info_2021 WHERE suin_cd = ?`
	rows, err := db.Query(qry, suin_cd)
	utils.HandleErr(err, "dbselect3")
	defer rows.Close()
	
	for rows.Next() {
		err = rows.Scan( 
			&n21.Recruit_volume_2021, &n21.Compete_ratio_2021, &n21.Recruit_replace_2021, 
			&n21.Nasin_best_2021, &n21.Nasin_mean_2021, &n21.Nasin_limit_2021, &n21.Recruit_factor_mean_2021,
		)
		utils.HandleErr(err, "load 2020 Value")
	}
	return n21
}


// 2020학년도 Info
func n20Info(n20 value2020, suin_cd string) value2020 {
	
	qry := `SELECT recruit_volume_2020, compete_ratio_2020, recruit_replace_2020, nasin_best_2020, nasin_mean_2020, 
	nasin_limit_2020, recruit_factor_mean_2020 FROM univ_susi_info_2020 WHERE suin_cd = ?`
	rows, err := db.Query(qry, suin_cd)
	utils.HandleErr(err, "dbselect3")
	defer rows.Close()
	
	for rows.Next() {
		err = rows.Scan( 
			&n20.Recruit_volume_2020, &n20.Compete_ratio_2020, &n20.Recruit_replace_2020, 
			&n20.Nasin_best_2020, &n20.Nasin_mean_2020, &n20.Nasin_limit_2020, &n20.Recruit_factor_mean_2020,
		)
		utils.HandleErr(err, "load 2020 Value")
	}
	return n20
}


// 2019학년도 Info
func n19Info(n19 value2019, suin_cd string) value2019 {
	
	qry := `SELECT recruit_volume_2019, compete_ratio_2019, recruit_replace_2019, nasin_best_2019, nasin_mean_2019, 
	nasin_limit_2019, recruit_factor_mean_2019 FROM univ_susi_info_2019 WHERE suin_cd = ?`
	rows, err := db.Query(qry, suin_cd)
	utils.HandleErr(err, "dbselect3")
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan( 
			&n19.Recruit_volume_2019, &n19.Compete_ratio_2019, &n19.Recruit_replace_2019, 
			&n19.Nasin_best_2019, &n19.Nasin_mean_2019, &n19.Nasin_limit_2019, &n19.Recruit_factor_mean_2019,
		)
		utils.HandleErr(err, "load 2019 Value")
	}	
	return n19
}