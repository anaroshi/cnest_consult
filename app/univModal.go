package app

import (
	"cnest_consult/utils"
	"net/http"
)

type detailUniv1 struct {
	Id int
	Suin_cd *string
	UnivName *string
	Apply_term_id *string
	Apply_form *string
	Apply_formName *string
	Apply_dept_id *string
	Apply_line_id *string
	ApSubjectName *string
	Recruit_volume int
	QualiName *string
	RecruitName *string
	Recruit1st *string
	Recruit1stVol *string
	RecruitFinal *string
	SulimitName *string
	NasinmetName *string
	DocmetName *string
	IntmetName *string
	NslmetName *string
	EtcmetName *string
	DoclstName *string
	AtwrtName *string
}	

type detailUniv2 struct {
	Remarks  string
	ApplyDatetime  string
	ApplyDocuDatetime  string
	RecruitSelectDate  string
	RecruitTestDate  string
	RecruitPassDate  string
	ResultRemarks  string
}

type value2021 struct {
	Recruit_volume_2021 float64
	Compete_ratio_2021 string
	Recruit_replace_2021 int
	Nasin_best_2021 float64
	Nasin_mean_2021 float64
	Nasin_limit_2021 float64
	Recruit_factor_mean_2021 float64
}

type value2020 struct {
	Recruit_volume_2020 float64
	Compete_ratio_2020 string
	Recruit_replace_2020 int
	Nasin_best_2020 float64
	Nasin_mean_2020 float64
	Nasin_limit_2020 float64
	Recruit_factor_mean_2020 float64
}

type value2019 struct {
	Recruit_volume_2019 float64
	Compete_ratio_2019 string
	Recruit_replace_2019 int
	Nasin_best_2019 float64
	Nasin_mean_2019 float64
	Nasin_limit_2019 float64
	Recruit_factor_mean_2019 float64
}

type items struct {
	Info_1 detailUniv1
	Info_2 detailUniv2
	N_21 value2021
	N_20 value2020
	N_19 value2019
}

func univDetail(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")	

	db = connDB()
	defer db.Close()

	// 대학정보1
	//
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
	WHERE u1.id=?
	`
	rows, err := db.Query(qry,id)
	utils.HandleError(w, err, "dbselect")
	defer rows.Close()


	var a detailUniv1
	
	for rows.Next() {
		err = rows.Scan(
			&a.Id, &a.Suin_cd, &a.UnivName, &a.Apply_term_id, &a.Apply_form, 
			&a.Apply_formName, &a.Apply_dept_id, &a.Apply_line_id, &a.ApSubjectName, &a.Recruit_volume, 
			&a.QualiName, &a.RecruitName, &a.Recruit1st, &a.Recruit1stVol, &a.RecruitFinal, 
			&a.SulimitName, &a.NasinmetName, &a.DocmetName, &a.IntmetName, &a.NslmetName, 
			&a.EtcmetName, &a.DoclstName, &a.AtwrtName,
		)
		utils.HandleError(w, err, "load detailUniv1 Value")
	}

	// 대학정보2
	//
	qry = `SELECT remarks, apply_datetime, apply_docu_datetime, recruit_select_date, recruit_test_date, 
	recruit_pass_date, result_remarks FROM univ_susi_info_snd WHERE suin_cd = ?`
	rows, err = db.Query(qry, &a.Suin_cd)
	utils.HandleError(w, err, "dbselect2")
	defer rows.Close()

	var b detailUniv2
	
	for rows.Next() {
		err = rows.Scan( 
			&b.Remarks, &b.ApplyDatetime, &b.ApplyDocuDatetime, &b.RecruitSelectDate, 
			&b.RecruitTestDate, &b.RecruitPassDate, &b.ResultRemarks,
		)
		utils.HandleError(w, err, "load detailUniv2 Value")
	}

	// 2021학년도
	//
	qry = `SELECT recruit_volume_2021, compete_ratio_2021, recruit_replace_2021, nasin_best_2021, nasin_mean_2021, 
	nasin_limit_2021, recruit_factor_mean_2021 FROM univ_susi_info_2021 WHERE suin_cd = ?`
	rows, err = db.Query(qry, &a.Suin_cd)
	utils.HandleError(w, err, "dbselect3")
	defer rows.Close()

	var n21 value2021
	
	for rows.Next() {
		err = rows.Scan( 
			&n21.Recruit_volume_2021, &n21.Compete_ratio_2021, &n21.Recruit_replace_2021, 
			&n21.Nasin_best_2021, &n21.Nasin_mean_2021, &n21.Nasin_limit_2021, &n21.Recruit_factor_mean_2021,
		)
		utils.HandleError(w, err, "load 2021 Value")
	}
	
	// 2020학년도
	// 
	qry = `SELECT recruit_volume_2020, compete_ratio_2020, recruit_replace_2020, nasin_best_2020, nasin_mean_2020, 
	nasin_limit_2020, recruit_factor_mean_2020 FROM univ_susi_info_2020 WHERE suin_cd = ?`
	rows, err = db.Query(qry, &a.Suin_cd)
	utils.HandleError(w, err, "dbselect3")
	defer rows.Close()

	var n20 value2020
	
	for rows.Next() {
		err = rows.Scan( 
			&n20.Recruit_volume_2020, &n20.Compete_ratio_2020, &n20.Recruit_replace_2020, 
			&n20.Nasin_best_2020, &n20.Nasin_mean_2020, &n20.Nasin_limit_2020, &n20.Recruit_factor_mean_2020,
		)
		utils.HandleError(w, err, "load 2020 Value")
	}


	// 2019학년도
	//
	qry = `SELECT recruit_volume_2019, compete_ratio_2019, recruit_replace_2019, nasin_best_2019, nasin_mean_2019, 
	nasin_limit_2019, recruit_factor_mean_2019 FROM univ_susi_info_2019 WHERE suin_cd = ?`
	rows, err = db.Query(qry, &a.Suin_cd)
	utils.HandleError(w, err, "dbselect3")
	defer rows.Close()

	var n19 value2019
	
	for rows.Next() {
		err = rows.Scan( 
			&n19.Recruit_volume_2019, &n19.Compete_ratio_2019, &n19.Recruit_replace_2019, 
			&n19.Nasin_best_2019, &n19.Nasin_mean_2019, &n19.Nasin_limit_2019, &n19.Recruit_factor_mean_2019,
		)
		utils.HandleError(w, err, "load 2019 Value")
	}
		
	data := items{Info_1:a, Info_2:b, N_21:n21, N_20:n20, N_19:n19, }

	tpl.ExecuteTemplate(w, "univDetail.gohtml", data)
}
