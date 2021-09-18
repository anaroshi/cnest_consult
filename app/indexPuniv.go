package app

import (
	"cnest_consult/univ/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var myValueStr string
var myValueFlt = 2.7

type univ struct {
	No int	`json:"no"`
	Id int	`json:"id"`
	UnivName string  `json:"univ_name"`
	Apply_formName string `json:"apply_formName"`
	Apply_dept_id string	`json:"apply_dept_id"`
	Apply_line_id string	`json:"apply_line_id"`
	Recruit_volume int	`json:"recruit_volume"`		
	RecruitName string	`json:"recruit_name"`
	Recruit1st string	`json:"recruit_1st"`
	Recruit1stVol string	`json:"recruit_1stVol"`
	RecruitFinal string	`json:"recruit_final"`	
	SulimitName string	`json:"sulimitName"`	
}

type applyData1 struct {	
	UnivInfo []univ
	ApplyDept []applyDeptList	
}

type applyData2 struct {	
	ApplyOpt option
	ApplyDept []applyDeptList	
}

type scatterData struct {
	Sq string
	UnivDeptNm string
	Nasin_mean_2021 float32
	SelfValue float32
}

// chart 생성
// 수정함
func indexPuniv(w http.ResponseWriter, r *http.Request) {
	
	// Required if you don't call r.FormValue()
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v", err)
	}
	fmt.Println("r.Form : ", r.Form)

	db = connDB()
	defer db.Close()

	univs := []univ{}
	applyDeptLists := []applyDeptList{}	
	
	dept := r.Form["selectedItem[]"]
	//fmt.Println(dept)
	
	if dept == nil || dept[0] == "" {		
		js := `<script type="text/javascript" charset="utf-8">
		alert("모집단위를 선택하세요!");
		</script>`

		w.Write([]byte(js))
		return
	}

	myValue := r.Form["myValue"]
	fAvgValue := r.Form["fAvgValue"]
	lAvgValue := r.Form["lAvgValue"]
	fmt.Println("myValue : ", myValue)
	fmt.Println("fAvgValue : ", fAvgValue)
	fmt.Println("lAvgValue : ", lAvgValue)
		
	formQry := []string{}
	// 모집단위 (쿼리조건1)
	for _, v := range dept {
		str := fmt.Sprintf(` u1.apply_dept_id = '%v' `,v)
		formQry = append(formQry, str)
	}
	formQry4 := strings.Join(formQry,"OR")
	qryStr := aaaa.getSrhOpt2(formQry4)

	// 내점수
	myValueFlt = 2.7
	if myValue != nil {
		myValueStr = myValue[0]
		myValueFlt, _ = strconv.ParseFloat(myValueStr, 64)
		qryStr = getMyOpts(myValueFlt, qryStr)
	}
	

	// 2021 내신 평균
	var fAvgValueInt float64
	var lAvgValueInt float64
	if fAvgValue != nil {		
		fAvgValueInt, _ = strconv.ParseFloat(fAvgValue[0], 64)
		lAvgValueInt, _ = strconv.ParseFloat(lAvgValue[0], 64)
		qryStr = getAvgOpts(fAvgValueInt, lAvgValueInt, qryStr)
	}

	// ----------- 조건에 부합하는 대학들
	qry := `SELECT count(*) OVER (ORDER BY A.id ) no, A.id, A.univName, A.apply_formName, A.apply_dept_id, A.apply_line_id, 
	A.recruit_volume, A.recruitName, A.recruit1st, A.recruit1stVol, A.recruitfinal, A.sulimitName  FROM
	( SELECT u1.id, u2.univName, u1.apply_formName, u1.apply_dept_id, u1.apply_line_id, u1.recruit_volume, 
	u3.recruitName, u3.recruit1st, u3.recruit1stVol, u3.recruitfinal, u4.sulimitName 
	FROM univ_susi_info_fst u1 
	INNER JOIN univ_info u2 ON u1.univ_id = u2.univId 
	INNER JOIN recruit_method u3 ON u1.recruit_met_id=u3.recruitId 
	INNER JOIN recruit_sununglimit u4 ON u1.recruit_su_lim_id = u4.sulimitId
	INNER JOIN univ_susi_info_2021 u5 ON u1.suin_cd=u5.suin_cd
	WHERE ` + qryStr + `ORDER BY u2.univName, u1.apply_formName, u1.apply_dept_id ) A`
	fmt.Println("qry :",qry )
	rows, err := db.Query(qry)
	utils.HandleError(w, err, "dbselect")

	var s univ
	
	for rows.Next() {
			
		err = rows.Scan(
			&s.No, &s.Id, &s.UnivName, &s.Apply_formName, &s.Apply_dept_id, &s.Apply_line_id, &s.Recruit_volume, &s.RecruitName, 
			&s.Recruit1st, &s.Recruit1stVol, &s.RecruitFinal, &s.SulimitName,
		)
		utils.HandleError(w, err, "loadValue")
		univs = append(univs, s)	
	}	
	data := applyData1 { UnivInfo: univs, ApplyDept: applyDeptLists, }	
	//fmt.Println(data)
	
	var sd scatterData

	// Chart Data	
	qry = `
		SELECT count(*) OVER (ORDER BY A.id ) no, A.subjNm, A.nasinMean FROM
		( SELECT CONCAT(u2.univName, '-', u1.apply_dept_id) subjNm, u5.nasin_mean_2021 nasinMean
		FROM univ_susi_info_fst u1 
		INNER JOIN univ_info u2 ON u1.univ_id = u2.univId 
		INNER JOIN univ_susi_info_2021 u5 ON u1.suin_cd=u5.suin_cd
		WHERE ` + qryStr + ` ORDER BY u2.univName, u1.apply_formName, u1.apply_dept_id) A`

		
	fmt.Println("qry :",qry )

	rows, err = db.Query(qry)
	utils.HandleError(w, err, "dbselect for chart")
	defer rows.Close()

	var (
		No []string
		SubjNm []string
		AvgValue []float32
		StdValue []float32
	)

	for rows.Next() {
			
		err = rows.Scan(
			&sd.Sq, &sd.UnivDeptNm, &sd.Nasin_mean_2021,
		)
		utils.HandleError(w, err, "load ChartValue")
		No = append(No, sd.Sq)
		SubjNm = append(SubjNm, sd.UnivDeptNm)
		AvgValue = append(AvgValue, sd.Nasin_mean_2021)
		StdValue = append(StdValue, float32(myValueFlt))
	}	
	
	// No   		= []int{1,2,3,4,5,6,7,8,9}
	// SubjNm   = []string{"Swimming", "Surfing", "Shooting ", "Skating", "Wrestling", "Diving", "Wrestling", "Diving", "Skating"}
	// AvgValue = []float32{2.9, 3.28, 2.73, 2.63, 3.39, 4.94, 4.27, 3.6, 1.85}
	//StdValue  = []float32{3, 3, 3, 3, 3, 3, 3, 3, 3}
	
	wg.Add(1)
	go chart(w, No, SubjNm, AvgValue, StdValue)
	wg.Wait()

	tpl.ExecuteTemplate(w, "indexPuniv.html", data)
}

func getMyOpts(myValueFlt float64, qryStr string) string {
	fMyValue := myValueFlt - 0.3
	lMyValue := myValueFlt + 0.3
	formQry5 := fmt.Sprintln(" AND u5.nasin_mean_2021 BETWEEN ", fMyValue, " AND ", lMyValue)
	qryStr += formQry5
	return qryStr
}

func getAvgOpts(fAvgValueInt float64, lAvgValueInt float64, qryStr string) string {
	if fAvgValueInt > lAvgValueInt {
		fAvgValueInt, lAvgValueInt = lAvgValueInt, fAvgValueInt
	} 
	formQry5 := fmt.Sprintln(" AND u5.nasin_mean_2021 BETWEEN ", fAvgValueInt, " AND ", lAvgValueInt)
	qryStr += formQry5
	return qryStr
}


// SELECT u2.univName,  u1.apply_dept_id, u3.nasin_mean_2021
//         FROM univ_susi_info_fst u1 
//         INNER JOIN univ_info u2 ON u1.univ_id = u2.univId 
//         INNER JOIN univ_susi_info_2021 u3 ON u1.suin_cd=u3.suin_cd
//         WHERE  ( apply_form = '학생부종합' ) AND apply_line_id = '인문' AND ( apply_sub_id = 'subj_1001' OR apply_sub_id = 'subj_1013' ) AND ( u1.apply_dept_id = '간호학과(인문)' OR u1.apply_dept_id = '간호학과(인문사회계열)' OR u1.apply_dept_id = '간호학부' )