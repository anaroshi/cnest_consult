package app

import (
	"cnest_consult/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type univ struct {
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

type Data [4]float32

type klineData struct {
	UnivName string  `json:"univ_name"`
	Data
}

// chart 생성
// html 내부에 포함 시키는 것으로
// head foot 태그 내용 제외시킴
// /home/sundor/workspace/go/src/github.com/go-echarts/go-echarts/templates/base.go
// /home/sundor/workspace/go/src/github.com/go-echarts/go-echarts/templates/page.go
// 수정함
func indexPuniv(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v", err)
	}
	db = connDB()
	defer db.Close()

	univs := []univ{}
	applyDeptLists := []applyDeptList{}	

	r.ParseForm() // Required if you don't call r.FormValue()
	fmt.Println("r.Form : ", r.Form)
	
	dept := r.Form["selectedItem[]"]
	//fmt.Println(dept)
	
	if dept == nil || dept[0] == "" {
		fmt.Println("*")
		js := `<script type="text/javascript" charset="utf-8">
		alert("모집단위를--선택하세요!");
		</script>`

		w.Write([]byte(js))
		tpl.ExecuteTemplate(w, "index.gohtml", univs)
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
	var myValueInt float64
	if myValue != nil {		
		myValueInt, _ = strconv.ParseFloat(myValue[0], 64)
		qryStr = getMyOpts(myValueInt, qryStr)
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
	qry := `SELECT u1.id, u2.univName, u1.apply_formName, u1.apply_dept_id, u1.apply_line_id, u1.recruit_volume, 
	u3.recruitName, u3.recruit1st, u3.recruit1stVol, u3.recruitfinal, u4.sulimitName 
	FROM univ_susi_info_fst u1 
	INNER JOIN univ_info u2 ON u1.univ_id = u2.univId 
	INNER JOIN recruit_method u3 ON u1.recruit_met_id=u3.recruitId 
	INNER JOIN recruit_sununglimit u4 ON u1.recruit_su_lim_id = u4.sulimitId
	INNER JOIN univ_susi_info_2021 u5 ON u1.suin_cd=u5.suin_cd
	WHERE ` + qryStr
	//fmt.Println("qry :",qry )
	rows, err := db.Query(qry)
	utils.HandleError(w, err, "dbselect")
	defer rows.Close()

	var s univ
	
	for rows.Next() {
			
		err = rows.Scan(
			&s.Id, &s.UnivName, &s.Apply_formName, &s.Apply_dept_id, &s.Apply_line_id, &s.Recruit_volume, &s.RecruitName, 
			&s.Recruit1st, &s.Recruit1stVol, &s.RecruitFinal, &s.SulimitName,
		)
		utils.HandleError(w, err, "loadValue")
		univs = append(univs, s)	
	}	
	data := applyData1 { UnivInfo: univs, ApplyDept: applyDeptLists, }	
	//fmt.Println(data)
	

	// var kd = []klineData{
	// 	{UnivName: "명지대", Data: [4]float32{4.1, 2.78, 2.78, 4.1}},
	// 	{UnivName: "시립대", Data: [4]float32{3.48, 1.8, 1.8, 3.48}},
	// 	{UnivName: "산업대", Data: [4]float32{3.7, 3.3, 3.3, 3.7}},
	// 	{UnivName: "경기대", Data: [4]float32{4.22, 3.98, 3.98, 4.22}},
	// 	{UnivName: "중앙대", Data: [4]float32{5.75, 4.48, 4.48, 5.75}},
	// 	{UnivName: "상명대", Data: [4]float32{4.1, 2.78, 2.78, 4.1}},
	// 	{UnivName: "숙대", Data: [4]float32{3.48, 1.8, 1.8, 3.48}},
	// 	{UnivName: "성대", Data: [4]float32{3.7, 3.3, 3.3, 3.7}},
	// 	{UnivName: "단대", Data: [4]float32{4.22, 3.98, 3.98, 4.22}},
	// 	{UnivName: "카톨릭대", Data: [4]float32{5.75, 4.48, 4.48, 5.75}},
	// }

	// var kd = []klineData{
	// 	{UnivName: "명지대", Data: [4]float32{4.4, 0, 4.4, 4.0}},
	// 	{UnivName: "시립대", Data: [4]float32{3.91, 2.48, 3.91, 2.48,}},
	// 	{UnivName: "산업대", Data: [4]float32{3.9, 0, 3.9, 0}},
	// 	{UnivName: "경기대", Data: [4]float32{4.72, 3.42, 4.72, 3.42}},
	// 	{UnivName: "중앙대", Data: [4]float32{2.4, 1.9, 2.4, 1.9}},
	// 	{UnivName: "상명대", Data: [4]float32{3.19, 0, 3.19, 0}},
	// 	{UnivName: "숙대", Data: [4]float32{2.93, 0, 2.93, 0}},
	// 	{UnivName: "성대", Data: [4]float32{2.8, 0, 2.8, 0}},
	// 	{UnivName: "단대", Data: [4]float32{0, 0, 3.98, 4.22}},
	// 	{UnivName: "카톨릭대", Data: [4]float32{3.61, 1.67, 3.61, 1.67}},
	// }

	chartData := []klineData{}
	var kd klineData

	// Chart Data	
	qry = `SELECT u2.univName, max(u5.nasin_mean_2021), min(u5.nasin_mean_2021), max(u5.nasin_mean_2021), min(u5.nasin_mean_2021)
	FROM univ_susi_info_fst u1 
	INNER JOIN univ_info u2 ON u1.univ_id = u2.univId 
	INNER JOIN univ_susi_info_2021 u5 ON u1.suin_cd=u5.suin_cd
	WHERE ` + qryStr + ` GROUP BY u1.univ_id`
	fmt.Println("qry :",qry )

	rows, err = db.Query(qry)
	utils.HandleError(w, err, "dbselect for chart")
	defer rows.Close()

	for rows.Next() {
			
		err = rows.Scan(
			&kd.UnivName, &kd.Data[0], &kd.Data[1], &kd.Data[2], &kd.Data[3],
		)
		utils.HandleError(w, err, "load ChartValue")
		chartData = append(chartData, kd)	
	}	

	fmt.Println(chartData)
	//chart(w, chartData)

	//chart(w, kd)

	tpl.ExecuteTemplate(w, "indexPuniv.gohtml", data)
}

func getMyOpts(myValueInt float64, qryStr string) string {
	fMyValue := myValueInt - 0.3
	lMyValue := myValueInt + 0.3
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