package app

import (
	"cnest_consult/utils"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type ApplyOptionList struct {
	OptionId string
	OptionName string
}

// 전형구분
func ApplyFormOptionList(db *sql.DB) {

	optList := ApplyOptionList{}
	optLists := []string{}
	optStr := `{{ define "applyFormOptionList" }}`
	optLists = append(optLists, optStr)
//	optLists = append(optLists, "<option selected>*</option>")
	
	qry := "SELECT formId, formName FROM apply_form"
	rows, err := db.Query(qry)
	utils.HandleErr(err,"applyFormOptionQuery")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&optList.OptionId, &optList.OptionName)
		utils.HandleErr(err,"applyFormOptionRows")
		//optStr = fmt.Sprint(`<option value="`+optList.OptionId+`">`+optList.OptionName+`</option>`)
		optStr = fmt.Sprint(`<option value="`+optList.OptionName+`">`+optList.OptionName+`</option>`)
		optLists = append(optLists, optStr)
	}
	optLists = append(optLists, "{{ end }}")
	optStr = strings.Join(optLists,"\n")
	
	nf, err := os.Create(PartsPath+"include_applyform_option.gohtml")
	utils.HandleErr(err,"applyFormOptionFileCreate")
	defer nf.Close()
	io.Copy(nf, strings.NewReader(optStr))
}

// 계열구분
func ApplyLineOptionList(db *sql.DB) {

	optList := ApplyOptionList{}
	optLists := []string{}
	optStr := `{{ define "applyLineOptionList" }}`
	optLists = append(optLists, optStr)
//	optLists = append(optLists, "<option selected>*</option>")
	
	qry := "SELECT lineId, lineName FROM apply_line"
	rows, err := db.Query(qry)
	utils.HandleErr(err,"applyLineOptionQuery")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&optList.OptionId, &optList.OptionName)
		utils.HandleErr(err,"applyLineOptionRows")
		//optStr = fmt.Sprint(`<option value="`+optList.OptionId+`">`+optList.OptionName+`</option>`)
		optStr = fmt.Sprint(`<option value="`+optList.OptionName+`">`+optList.OptionName+`</option>`)
		optLists = append(optLists, optStr)
	}
	optLists = append(optLists, "{{ end }}")
	optStr = strings.Join(optLists,"\n")
	
	nf, err := os.Create(PartsPath+"include_applyline_option.gohtml")
	utils.HandleErr(err,"applyLineOptionFileCreate")
	defer nf.Close()
	io.Copy(nf, strings.NewReader(optStr))
}

// 전공구분
func ApplySubjectOptionList(db *sql.DB) {

	optList := ApplyOptionList{}
	optLists := []string{}
	optStr := `{{ define "applySubjectOptionList" }}`
	optLists = append(optLists, optStr)
//	optLists = append(optLists, "<option selected>*</option>")
	
	qry := "SELECT apSubjectId, apSubjectName FROM apply_subject"
	rows, err := db.Query(qry)
	utils.HandleErr(err,"applySubjectOptionQuery")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&optList.OptionId, &optList.OptionName)
		utils.HandleErr(err,"applySubjectOptionRows")
		optStr = fmt.Sprint(`<option value="`+optList.OptionId+`">`+optList.OptionName+`</option>`)
		optLists = append(optLists, optStr)
	}
	optLists = append(optLists, "{{ end }}")
	optStr = strings.Join(optLists,"\n")
	
	nf, err := os.Create(PartsPath+"include_applysubject_option.gohtml")
	utils.HandleErr(err,"applySubjectOptionFileCreate")
	defer nf.Close()
	io.Copy(nf, strings.NewReader(optStr))
}

//모집단위
// func applyDeptOptionList(db *sql.DB) {

// 	optList := ApplyOptionList{}
// 	optLists := []string{}
// 	optStr := `{{ define "applyDeptOptionList" }}`
// 	optLists = append(optLists, optStr)
// 	optLists = append(optLists, "<option selected>*</option>")
	
// 	qry := "SELECT deptId, deptName FROM apply_dept"
// 	rows, err := db.Query(qry)
// 	utils.HandleErr(err,"applyDeptOptionQuery")
// 	defer rows.Close()

// 	for rows.Next() {
// 		err := rows.Scan(&optList.OptionId, &optList.OptionName)
// 		utils.HandleErr(err,"applyDeptOptionRows")
// 		optStr = fmt.Sprint(`<option value="`+optList.OptionId+`">`+optList.OptionName+`</option>`)		
// 		optLists = append(optLists, optStr)
// 	}
// 	optLists = append(optLists, "{{ end }}")
// 	optStr = strings.Join(optLists,"\n")
	
// 	//nf, err := os.Create("/home/sundor/workspace/go/src/cnest/db/../templates/parts/applydept_option.gohtml")
// 	nf, err := os.Create(PartsPath+"include_applydept_option.gohtml")
// 	utils.HandleErr(err,"applyDeptOptionFileCreate")
// 	defer nf.Close()
// 	io.Copy(nf, strings.NewReader(optStr))
// }


// func main() {

// 	db, err := connDB()
// 	utils.HandleErr(err, "callCommDB")
// 	defer db.Close()
	
// 	applyFormOptionList(db)
// 	applyLineOptionList(db)
// 	applySubjectOptionList(db)
// 	applyDeptOptionList(db)

// dir, _ := os.Getwd()
// fmt.Println(strings.Replace(dir, " ", "\\ ", -1))


// }

