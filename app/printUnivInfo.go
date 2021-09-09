package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/xuri/excelize"
)

func printUnivInfo(w http.ResponseWriter, r *http.Request) {
	// Required if you don't call r.FormValue()
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "printUnivInfo err %v", err)
	}	
	fmt.Println(r.Form)
	id := r.Form["chk_val[]"]

	if id == nil || id[0] == "" {
		js := `<script type="text/javascript" charset="utf-8">
			alert("지원대학을 선택하세요!");
			</script>
		`
		w.Write([]byte(js))
		return
	}

	
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	
	// Set Title of a cell.
	f.SetCellValue("Sheet2", "A1", "지원 가능 대학")
	f.SetCellValue("Sheet2", "A2", "Id")
	f.SetCellValue("Sheet2", "B2", "대학교")
	f.SetCellValue("Sheet2", "C2", "모집시기")
	f.SetCellValue("Sheet2", "D2", "전형구분")
	f.SetCellValue("Sheet2", "E2", "전형명칭")
	f.SetCellValue("Sheet2", "F2", "모집단위")
	f.SetCellValue("Sheet2", "G2", "계열구분")
	f.SetCellValue("Sheet2", "H2", "전공구분")
	f.SetCellValue("Sheet2", "I2", "인원")
	f.SetCellValue("Sheet2", "J2", "지원자격")
	f.SetCellValue("Sheet2", "K2", "전형방법")
	f.SetCellValue("Sheet2", "L2", "1단계")
	f.SetCellValue("Sheet2", "M2", "2단계")
	f.SetCellValue("Sheet2", "N2", "최종")
	f.SetCellValue("Sheet2", "O2", "수능최저학력기준")
	f.SetCellValue("Sheet2", "P2", "학생부(교과)")
	f.SetCellValue("Sheet2", "Q2", "제출서류")
	f.SetCellValue("Sheet2", "R2", "면접")
	f.SetCellValue("Sheet2", "S2", "논술")
	f.SetCellValue("Sheet2", "T2", "기타")
	f.SetCellValue("Sheet2", "U2", "서류목록")
	f.SetCellValue("Sheet2", "V2", "자소서3번문항")
	
	f.SetCellValue("Sheet2", "W2", "비고")
	f.SetCellValue("Sheet2", "X2", "원서접수")
	f.SetCellValue("Sheet2", "Y2", "서류접수")
	f.SetCellValue("Sheet2", "Z2", "전형일")
	f.SetCellValue("Sheet2", "AA2", "1단계")
	f.SetCellValue("Sheet2", "AB2", "최종")
	f.SetCellValue("Sheet2", "AC2", "결과비고")

	f.SetCellValue("Sheet2", "AD2", "2021모집인원")
	f.SetCellValue("Sheet2", "AE2", "2021경쟁룰")
	f.SetCellValue("Sheet2", "AF2", "2021충원인원")
	f.SetCellValue("Sheet2", "AG2", "2021내신최고")
	f.SetCellValue("Sheet2", "AH2", "2021내신평균")
	f.SetCellValue("Sheet2", "AI2", "2021내신최저")
	f.SetCellValue("Sheet2", "AJ2", "2021전형평균")
	
	f.SetCellValue("Sheet2", "AK2", "2020모집인원")
	f.SetCellValue("Sheet2", "AL2", "2020경쟁룰")
	f.SetCellValue("Sheet2", "AM2", "2020충원인원")
	f.SetCellValue("Sheet2", "AN2", "2020내신최고")
	f.SetCellValue("Sheet2", "AO2", "2020내신평균")
	f.SetCellValue("Sheet2", "AP2", "2020내신최저")
	f.SetCellValue("Sheet2", "AQ2", "2020전형평균")

	f.SetCellValue("Sheet2", "AR2", "2019모집인원")
	f.SetCellValue("Sheet2", "AS2", "2019경쟁룰")
	f.SetCellValue("Sheet2", "AT2", "2019충원인원")
	f.SetCellValue("Sheet2", "AU2", "2019내신최고")
	f.SetCellValue("Sheet2", "AV2", "2019내신평균")
	f.SetCellValue("Sheet2", "AW2", "2019내신최저")
	f.SetCellValue("Sheet2", "AX2", "2019전형평균")


	db = connDB()
	defer db.Close()

	var unviAllInfo items
	unviAllInfos := []items{}

	for i, eachId := range id {		
		unviAllInfo = GetDetailUniv(db, eachId)
		unviAllInfos = append(unviAllInfos, unviAllInfo)

		// Set value of a cell.
		f.SetCellValue("Sheet2", "A"+strconv.Itoa(i+3), i+1)	
		f.SetCellValue("Sheet2", "B"+strconv.Itoa(i+3), unviAllInfo.Info_1.UnivName)
		f.SetCellValue("Sheet2", "C"+strconv.Itoa(i+3), unviAllInfo.Info_1.Apply_term_id)
		f.SetCellValue("Sheet2", "D"+strconv.Itoa(i+3), unviAllInfo.Info_1.Apply_form)
		f.SetCellValue("Sheet2", "E"+strconv.Itoa(i+3), unviAllInfo.Info_1.Apply_formName)
		f.SetCellValue("Sheet2", "F"+strconv.Itoa(i+3), unviAllInfo.Info_1.Apply_dept_id)
		f.SetCellValue("Sheet2", "G"+strconv.Itoa(i+3), unviAllInfo.Info_1.Apply_line_id)
		f.SetCellValue("Sheet2", "H"+strconv.Itoa(i+3), unviAllInfo.Info_1.ApSubjectName)
		f.SetCellValue("Sheet2", "I"+strconv.Itoa(i+3), unviAllInfo.Info_1.Recruit_volume)
		f.SetCellValue("Sheet2", "J"+strconv.Itoa(i+3), unviAllInfo.Info_1.QualiName)
		f.SetCellValue("Sheet2", "K"+strconv.Itoa(i+3), unviAllInfo.Info_1.RecruitName)
		f.SetCellValue("Sheet2", "L"+strconv.Itoa(i+3), unviAllInfo.Info_1.Recruit1st)
		f.SetCellValue("Sheet2", "M"+strconv.Itoa(i+3), unviAllInfo.Info_1.Recruit1stVol)
		f.SetCellValue("Sheet2", "N"+strconv.Itoa(i+3), unviAllInfo.Info_1.RecruitFinal)
		f.SetCellValue("Sheet2", "O"+strconv.Itoa(i+3), unviAllInfo.Info_1.SulimitName)
		f.SetCellValue("Sheet2", "P"+strconv.Itoa(i+3), unviAllInfo.Info_1.NasinmetName)
		f.SetCellValue("Sheet2", "Q"+strconv.Itoa(i+3), unviAllInfo.Info_1.DocmetName)
		f.SetCellValue("Sheet2", "R"+strconv.Itoa(i+3), unviAllInfo.Info_1.IntmetName)
		f.SetCellValue("Sheet2", "S"+strconv.Itoa(i+3), unviAllInfo.Info_1.NasinmetName)
		f.SetCellValue("Sheet2", "T"+strconv.Itoa(i+3), unviAllInfo.Info_1.EtcmetName)
		f.SetCellValue("Sheet2", "U"+strconv.Itoa(i+3), unviAllInfo.Info_1.DoclstName)
		f.SetCellValue("Sheet2", "V"+strconv.Itoa(i+3), unviAllInfo.Info_1.AtwrtName)

		f.SetCellValue("Sheet2", "W"+strconv.Itoa(i+3), unviAllInfo.Info_2.Remarks)
		f.SetCellValue("Sheet2", "X"+strconv.Itoa(i+3), unviAllInfo.Info_2.ApplyDatetime)
		f.SetCellValue("Sheet2", "Y"+strconv.Itoa(i+3), unviAllInfo.Info_2.ApplyDocuDatetime)
		f.SetCellValue("Sheet2", "Z"+strconv.Itoa(i+3), unviAllInfo.Info_2.RecruitSelectDate)
		f.SetCellValue("Sheet2", "AA"+strconv.Itoa(i+3), unviAllInfo.Info_2.RecruitTestDate)
		f.SetCellValue("Sheet2", "AB"+strconv.Itoa(i+3), unviAllInfo.Info_2.RecruitPassDate)
		f.SetCellValue("Sheet2", "AC"+strconv.Itoa(i+3), unviAllInfo.Info_2.ResultRemarks)

		f.SetCellValue("Sheet2", "AD"+strconv.Itoa(i+3), unviAllInfo.N_21.Recruit_volume_2021)
		f.SetCellValue("Sheet2", "AE"+strconv.Itoa(i+3), unviAllInfo.N_21.Compete_ratio_2021)
		f.SetCellValue("Sheet2", "AF"+strconv.Itoa(i+3), unviAllInfo.N_21.Recruit_replace_2021)
		f.SetCellValue("Sheet2", "AG"+strconv.Itoa(i+3), unviAllInfo.N_21.Nasin_best_2021)
		f.SetCellValue("Sheet2", "AH"+strconv.Itoa(i+3), unviAllInfo.N_21.Nasin_mean_2021)
		f.SetCellValue("Sheet2", "AI"+strconv.Itoa(i+3), unviAllInfo.N_21.Nasin_limit_2021)
		f.SetCellValue("Sheet2", "AJ"+strconv.Itoa(i+3), unviAllInfo.N_21.Recruit_factor_mean_2021)

		f.SetCellValue("Sheet2", "AK"+strconv.Itoa(i+3), unviAllInfo.N_20.Recruit_volume_2020)
		f.SetCellValue("Sheet2", "AL"+strconv.Itoa(i+3), unviAllInfo.N_20.Compete_ratio_2020)
		f.SetCellValue("Sheet2", "AM"+strconv.Itoa(i+3), unviAllInfo.N_20.Recruit_replace_2020)
		f.SetCellValue("Sheet2", "AN"+strconv.Itoa(i+3), unviAllInfo.N_20.Nasin_best_2020)
		f.SetCellValue("Sheet2", "AO"+strconv.Itoa(i+3), unviAllInfo.N_20.Nasin_mean_2020)
		f.SetCellValue("Sheet2", "AP"+strconv.Itoa(i+3), unviAllInfo.N_20.Nasin_limit_2020)
		f.SetCellValue("Sheet2", "AQ"+strconv.Itoa(i+3), unviAllInfo.N_20.Recruit_factor_mean_2020)

		f.SetCellValue("Sheet2", "AR"+strconv.Itoa(i+3), unviAllInfo.N_19.Recruit_volume_2019)
		f.SetCellValue("Sheet2", "AS"+strconv.Itoa(i+3), unviAllInfo.N_19.Compete_ratio_2019)
		f.SetCellValue("Sheet2", "AT"+strconv.Itoa(i+3), unviAllInfo.N_19.Recruit_replace_2019)
		f.SetCellValue("Sheet2", "AU"+strconv.Itoa(i+3), unviAllInfo.N_19.Nasin_best_2019)
		f.SetCellValue("Sheet2", "AV"+strconv.Itoa(i+3), unviAllInfo.N_19.Nasin_mean_2019)
		f.SetCellValue("Sheet2", "AW"+strconv.Itoa(i+3), unviAllInfo.N_19.Nasin_limit_2019)
		f.SetCellValue("Sheet2", "AX"+strconv.Itoa(i+3), unviAllInfo.N_19.Recruit_factor_mean_2019)

	}

	fmt.Println(unviAllInfos)

	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("data-selectedUnvi.xlsx"); err != nil {
			fmt.Println(err)
	}
}

