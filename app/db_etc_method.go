package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_etc_method(db *sql.DB) {

	// Drop table
	statement, err := db.Prepare(`DROP TABLE IF EXISTS etc_method;`)

	utils.HandleErr(err, "Drop etc_method table")
	_, err = statement.Exec()
	utils.HandleErr(err, "Exec for Drop etc_method table")


	// Create table
	statement, err = db.Prepare(
		`CREATE TABLE etc_method (
			id SERIAL PRIMARY KEY,
			etcmetId varchar(16) NOT NULL,
			etcmetName varchar(500) NOT NULL,
			sort integer,
			flag integer,
			inUser varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP
		);`)

	utils.HandleErr(err, "Create etc_method table")
	_, err = statement.Exec()
	utils.HandleErr(err, "Exec for Create etc_method table")


	// Insert Data
	stmt, err := db.Prepare(
		`INSERT INTO etc_method (
			id, etcmetId, etcmetName, sort, flag) VALUES		
			(1, 'etmt_1001', '(기타) 교육과정 관련 활동 증빙, A4 5매 이내 PDF파일로 제출', 0, 0),
			(2, 'etmt_1002', '(기타) 교육과정 관련 활동 증빙, A4 6매 이내 PDF파일로 제출', 0, 0),
			(3, 'etmt_1003', '(면접기초자료) A4 1장(2~3문항) 간략한 자기소개 등, 20분 작성', 0, 0),
			(4, 'etmt_1004', '(봉사)181시간 이상 만점 (출결)각 해당 항목별 감점', 0, 0),
			(5, 'etmt_1005', '(실기)앉아 윗몸 앞으로 굽히기, 제자리 멀리뛰기, 팔굽혀 매달리기', 0, 0),
			(6, 'etmt_1006', '(실기)작문(운문 또는 산문), 60분', 0, 0),
			(7, 'etmt_1007', '(자소서)공통 2번과 3번 문항', 0, 0),
			(8, 'etmt_1008', '(출결)49시간 이상 만점', 0, 0),
			(9, 'etmt_1009', '(출결)결석 1일당 1점 감점, 지각·조퇴·결과 3회당 1점 감점', 0, 0),
			(10, 'etmt_1010', '(출결)결석 4일이하 만점', 0, 0),
			(11, 'etmt_1011', '(출결)결석일수 0일 이하 만점', 0, 0),
			(12, 'etmt_1012', '(출결)무단/미인정 1일 이하, (봉사)전원 만점 처리', 0, 0),
			(13, 'etmt_1013', '(출결)미인정(무단) 결석 3일 이하 만점', 0, 0),
			(14, 'etmt_1014', '(출결)미인정(무단) 결석 3일이하 만점', 0, 0),
			(15, 'etmt_1015', '(출결)미인정결석 0일 이하 만점', 0, 0),
			(16, 'etmt_1016', '(출결)사고 결석 0일 이하 만점', 0, 0),
			(17, 'etmt_1017', '(출결)사고결석 3일 이하 만점 (봉사)코로나-19로 인한 만점', 0, 0),
			(18, 'etmt_1018', '(출석)결석 1일당 1점 감점', 0, 0),
			(19, 'etmt_1019', '(출석)무단 1일이하 만점', 0, 0),
			(20, 'etmt_1020', '(출석)미인정 결석 1일이하 만점', 0, 0),
			(21, 'etmt_1021', '(활동보고서)수학, 과학관련 수상, 활동 관련 기술 1,000자', 0, 0),
			(22, 'etmt_1022', '(활동보고서)어학관련 수상, 인증, 활동 관련 기술 1,000자', 0, 0),
			(23, 'etmt_1023', '(활동보고서)영어관련 학업능력, 인증, 수사 및 활동 관련 영어 Writing 600자 내외', 0, 0),
			(24, 'etmt_1024', 'COVID-19로 인하여 전원 만점 처리, (출결)무단 또는 미인정 3일이하, (봉사)25시간이상 만점', 0, 0),
			(25, 'etmt_1025', 'COVID19상황으로 모두 만점처리((출결봉사)무단 3일이하, 28시간이상 만점)', 0, 0),
			(26, 'etmt_1026', 'MMPI-Ⅱ 인성검사(합/불 판정)', 0, 0),
			(27, 'etmt_1027', '미인정 결석 3일이하 만점', 0, 0),
			(28, 'etmt_1028', '미인정 결석 3일이하 만점 (실기)운문(글자 수 제한 없음) 또는 산문(2,000자 이내), 표현력(40) 창의력(40) 정서법(20), 120분', 0, 0),
			(29, 'etmt_1029', '미인정(무단) 결석 2일 이하 및 15시간 이상 봉사시 만점', 0, 0),
			(30, 'etmt_1030', '미인정결석 3일이하 및 봉사활동 20시간 이상 만점', 0, 0),
			(31, 'etmt_1031', '신체검사, 체력검정, 신원조사(합·불 판정)', 0, 0),
			(32, 'etmt_1032', '해당없음', 0, 0),
			(33, 'etmt_1033', '활동자료실적물 A4 20매이내', 0, 0),
			(34, 'etmt_9999', '', 0, 0)
		;`)
	utils.HandleErr(err, "Insert Data into etc_method table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Insert Data into etc_method table")
}
