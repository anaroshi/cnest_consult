package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_autowriting(db *sql.DB) {

	// Drop table
	statement, err := db.Prepare(`DROP TABLE IF EXISTS autowriting;`)

	utils.HandleErr(err, "Drop autowriting table")
	_, err = statement.Exec()
	utils.HandleErr(err, "Exec for Drop autowriting table")


	// Create table
	statement, err = db.Prepare(
		`CREATE TABLE autowriting (
			id SERIAL PRIMARY KEY,
			atwrtId varchar(16) NOT NULL,
			atwrtName varchar(120) NOT NULL,
			sort integer,
			flag integer,
			inUser varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP
		);`)

	utils.HandleErr(err, "Create atwrtName table")
	_, err = statement.Exec()
	utils.HandleErr(err, "Exec for Create atwrtName table")


	// Insert Data
	stmt, err := db.Prepare(
		`INSERT INTO autowriting (
			id, atwrtId, atwrtName, sort, flag) VALUES
			(1, 'atwr_1001', '없음', 0, 0),
			(2, 'atwr_1002', '지원동기와 노력과정 / 1,500자', 0, 0),
			(3, 'atwr_1003', '지원동기와 준비과정 800자', 0, 0),
			(4, 'atwr_1004', '입학 후 학업 및 진로계획 800자', 0, 0),
			(5, 'atwr_1005', '지원동기 및 노력 800자', 0, 0),
			(6, 'atwr_1006', '지원동기 및 노력과정 800자', 0, 0),
			(7, 'atwr_1007', '지원동기 및 노력과정 800자 이내', 0, 0),
			(8, 'atwr_1008', '지원동기 및 진로계획 800자', 0, 0),
			(9, 'atwr_1009', '지원동기 및 학업 또는 진로계획, 800자 이내', 0, 0),
			(10, 'atwr_1010', '지원동기 및 학업계획, 800자 이내', 0, 0),
			(11, 'atwr_1011', '지원동기와 노력과정 800자', 0, 0),
			(12, 'atwr_1012', '지원동기와 노력과정 800자 이내', 0, 0),
			(13, 'atwr_1013', '지원동기와 지원분야 도전경험 800자 이내', 0, 0),
			(14, 'atwr_1014', '지원동기와 진로계획을 위해 수행한 노력과 준비 기술(800자 이내)', 0, 0),
			(15, 'atwr_1015', '지원동기와 학업계획 800자', 0, 0),
			(16, 'atwr_1016', '지원동기와 학업계획 및 진로계획 800자', 0, 0),
			(17, 'atwr_1017', '지원동기와 학업계획 및 진로계획, 800자', 0, 0),
			(18, 'atwr_1018', '최근 3년 가장 큰 영향 준 책 2권, 각 도서별 400자', 0, 0),
			(19, 'atwr_1019', '특이 항목', 0, 0)
		;`)
	utils.HandleErr(err, "Insert Data into autowriting table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Insert Data into autowriting table")
}
