package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_apply_subject(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS apply_subject;`)
	utils.HandleErr(err, "Drop apply_subject table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop apply_subject table")

	// Create Table
	stmt, err = db.Prepare(
		`CREATE TABLE IF NOT EXISTS apply_subject (
			id SERIAL PRIMARY KEY,
			apSubjectId VARCHAR(16) NOT NULL ,
			apSubjectName VARCHAR(32) NOT NULL,
			sort integer,
			flag integer,
			inUser VARCHAR(24),
			inDate TIMESTAMP,
			reUser VARCHAR(24),
			reDate TIMESTAMP
		);`)

	utils.HandleErr(err, "Create apply_subject table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create apply_subject table")

	//  Insert Data
	stmt, err = db.Prepare(`
		INSERT INTO apply_subject (id, apSubjectId, apSubjectName, sort, flag) VALUES
		(1, 'subj_1001', '간호', 0, 0),
		(2, 'subj_1002', '건축공', 0, 0),
		(3, 'subj_1003', '게임', 0, 0),
		(4, 'subj_1004', '공학기타', 0, 0),
		(5, 'subj_1005', '관광', 0, 0),
		(6, 'subj_1006', '국방', 0, 0),
		(7, 'subj_1007', '기계공', 0, 0),
		(8, 'subj_1008', '농학', 0, 0),
		(9, 'subj_1009', '디자인', 0, 0),
		(10, 'subj_1010', '미디어', 0, 0),
		(11, 'subj_1011', '미술', 0, 0),
		(12, 'subj_1012', '법행정', 0, 0),
		(13, 'subj_1013', '보건', 0, 0),
		(14, 'subj_1014', '사범(인문)', 0, 0),
		(15, 'subj_1015', '사범(자연)', 0, 0),
		(16, 'subj_1016', '사회', 0, 0),
		(17, 'subj_1017', '상경', 0, 0),
		(18, 'subj_1018', '생활', 0, 0),
		(19, 'subj_1019', '수의예', 0, 0),
		(20, 'subj_1020', '약학', 0, 0),
		(21, 'subj_1021', '어문', 0, 0),
		(22, 'subj_1022', '영화연극', 0, 0),
		(23, 'subj_1023', '예체능', 0, 0),
		(24, 'subj_1024', '의치한', 0, 0),
		(25, 'subj_1025', '이학', 0, 0),
		(26, 'subj_1026', '인문', 0, 0),
		(27, 'subj_1027', '자율전공', 0, 0),
		(28, 'subj_1028', '전자공', 0, 0),
		(29, 'subj_1029', '정보통신', 0, 0),
		(30, 'subj_1030', '체육', 0, 0),
		(31, 'subj_1031', '화생공', 0, 0),
		(32, 'subj_1032', '사범(초등)', 0, 0);
	`)
	
	utils.HandleErr(err, "Insert Data into apply_subject table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Insert Data into apply_subject table")
}
