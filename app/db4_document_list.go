package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db4_document_list(db *sql.DB) {

	// Drop table
	statement, err := db.Prepare(`DROP TABLE IF EXISTS document_list;`)

	utils.HandleErr(err, "drop table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table11 exec")


	// Create table
	statement, err = db.Prepare(
		`CREATE TABLE document_list (
			id SERIAL PRIMARY KEY,
			doclstId varchar(16) NOT NULL,
			doclstName varchar(500) NOT NULL,
			sort integer,
			flag integer,
			inUser varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP
		);`)

	utils.HandleErr(err, "create table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "create table11 exec")


	// Insert Data
	stmt, err := db.Prepare(
		`INSERT INTO document_list (
		id, doclstId, doclstName, sort, flag) VALUES		
		(1, 'dcli_1001', '(1)학교생활기록부 (2)자기소개서, 신원진술서 등등', 0, 0),
		(2, 'dcli_1002', '학교생활기록부', 0, 0),
		(3, 'dcli_1003', '학교생활기록부, 실적확인서 또는 문예지등단 확인서', 0, 0),
		(4, 'dcli_1004', '학교생활기록부, 어학성적표', 0, 0),
		(5, 'dcli_1005', '학교생활기록부, 입상실적 증명서', 0, 0),
		(6, 'dcli_1006', '학교생활기록부, 입상실적확인서, *단, 포트폴리오(작품당 A4 단면 2장, 총 작품 5개 이내)는 면접시 제출', 0, 0),
		(7, 'dcli_1007', '학교생활기록부, 자기소개서', 0, 0),
		(8, 'dcli_1008', '학교생활기록부, 자기소개서(공통 1번 2번 문항)', 0, 0),
		(9, 'dcli_1009', '학교생활기록부, 자기소개서(공통 1번 2번 문항), 추천인확인서 또는 세례성사증명서', 0, 0),
		(10, 'dcli_1010', '학교생활기록부, 자기소개서, 개인활동자료 및 실적물(A4 단면 최대 20장까지)', 0, 0),
		(11, 'dcli_1011', '학교생활기록부, 자기소개서, 교사추천서', 0, 0),
		(12, 'dcli_1012', '학교생활기록부, 자기소개서, 교사추천서, 특기증빙자료 및 목록 (5건 이내)', 0, 0),
		(13, 'dcli_1013', '학교생활기록부, 자기소개서, 기타입증자료', 0, 0),
		(14, 'dcli_1014', '학교생활기록부, 자기소개서, 우수성 입증자료', 0, 0),
		(15, 'dcli_1015', '학교생활기록부, 자기소개서, 활동보고서', 0, 0),
		(16, 'dcli_1016', '학교생활기록부, 학교장 추천 명단', 0, 0),
		(17, 'dcli_1017', '학교생활기록부, 활동보고서 및 증빙서류', 0, 0)
		;`)
	utils.HandleErr(err, "create table1")
	_, err = stmt.Exec()
	utils.HandleErr(err, "create table11 exec")
}
