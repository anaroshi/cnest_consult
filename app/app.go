package app

import (
	"cnest_consult/utils"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// type rawTime []byte

// func (t rawTime) Time() (time.Time, error) {
//     return time.Parse("15:04:05", string(t))
// }

const PartsPath string = "templates/"

var (
	tpl *template.Template
	db *sql.DB	
)

func init() {	

	db = connDB()
	defer db.Close()

	ApplyFormOptionList(db)
	ApplyLineOptionList(db)
	ApplySubjectOptionList(db)
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type applyDeptList struct {
	Apply_dept_id string
}

type option struct {
	OptForm string
	OptLine string
	OptSubject string
}

type searchApplyOption struct {
	ApplyOptions1 string
	ApplyOptions2 string
}

var aaaa searchApplyOption

func indexG(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func indexP(w http.ResponseWriter, r *http.Request) {
	
	db = connDB()
	defer db.Close()

	applyDeptLists := []applyDeptList{}
if r.Method == http.MethodPost {
	r.ParseForm() // Required if you don't call r.FormValue()
	fmt.Println("r.Form : ", r.Form)
	form := r.Form["getApplyForm[]"]		
	line := r.Form["getApplyLine"]
	subject := r.Form["getApplySubject[]"]

	fmt.Println(form,"/",line,"/",subject)
	
	formQry := []string{}
	formOpt := []string{}	
	// 전형구분 (쿼리조건1)
	for _, v := range form {
		str := fmt.Sprintf(` apply_form = '%v' `,v)
		formQry = append(formQry, str)
		formOpt = append(formOpt, v)
	}
	optform := strings.Join(formOpt,",")
	formQry1 := strings.Join(formQry,"OR")
	formQry1 = " (" + formQry1 + ") "

	// 계열구분 (쿼리조건2)
	formQry2 := fmt.Sprintf(` apply_line_id = '%v' `, line[0])		

	// 전공구분 (쿼리조건3)
	formQry = []string{}
	formOpt = []string{}
	for _, v := range subject {
		str := fmt.Sprintf(` apply_sub_id = '%v' `,v)
		formQry = append(formQry, str)
		formOpt = append(formOpt, v)
	}
	optsub := strings.Join(formOpt,",")
	formQry3 := strings.Join(formQry,"OR")
	formQry3 = " (" + formQry3 + ") "

	options := option{ optform, line[0], optsub}

	// 쿼리 조건 완성
	formQrys := []string{}
	formQrys = append(formQrys, formQry1)
	formQrys = append(formQrys, formQry2)
	formQrys = append(formQrys, formQry3)
	formQryList := strings.Join(formQrys,"AND")
	qryStr := aaaa.getSrhOpt1(formQryList)
	//fmt.Println(formQryList)

	qry := "SELECT apply_dept_id FROM univ_susi_info_fst WHERE " + qryStr + " GROUP BY apply_dept_id ORDER BY apply_dept_id"
	fmt.Println(qry)
	rows, err := db.Query(qry)
	utils.HandleError(w, err, "applyDept Select")
	defer rows.Close()

	var af applyDeptList			
	
	for rows.Next() {
		err := rows.Scan( &af.Apply_dept_id )
		utils.HandleError(w, err, "applyFormList Load")
		applyDeptLists = append(applyDeptLists, af)
	}
	fmt.Println(applyDeptLists)
	ad := applyData2{ ApplyOpt :options, ApplyDept :applyDeptLists,}
	fmt.Println(ad)
	fmt.Println(options)
}
	tpl.ExecuteTemplate(w, "applyDept.gohtml", applyDeptLists)
}

//DB connection
func connDB() *sql.DB {
	connStr := "sundor:anaroshi@tcp(localhost:3306)/cst_univ"
	db, err := sql.Open("mysql", connStr)
	utils.HandleErr(err, "dbconn")
	err = db.Ping()
	utils.HandleErr(err, "dbping")
	fmt.Println("DB connected")
	return db
}

func getChart(w http.ResponseWriter, r *http.Request) {	
	tpl.ExecuteTemplate(w,"chart.gohtml", "")
}

func MakeHandler() http.Handler {
	
	// db = connDB()
	//defer db.Close()
	
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexG).Methods("GET")
	mux.HandleFunc("/indexP", indexP).Methods("GET")
	mux.HandleFunc("/indexP", indexP).Methods("POST")
	mux.HandleFunc("/indexPuniv", indexPuniv).Methods("POST")
	mux.HandleFunc("/univDetail", univDetail)	
	mux.HandleFunc("/chart", getChart)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./public/"))))
	return mux
}

// 전형구분, 계열구분, 전공구분 쿼리 조건
func (sao *searchApplyOption) getSrhOpt1(ao string) string {
	sao.ApplyOptions1 = ao
	return sao.ApplyOptions1
}

// 전형, 계열, 전공 구분 + 모집단위 쿼리 조건
func (sao *searchApplyOption) getSrhOpt2(ao string) string {
	sao.ApplyOptions2 = sao.ApplyOptions1 + "AND (" + ao + ") "
	return sao.ApplyOptions2
}

func selectbox(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "selectbox.gohtml", nil)
}