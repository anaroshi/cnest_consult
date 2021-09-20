package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_univ_info(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS univ_info;`)
	utils.HandleErr(err, "Drop univ_info table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop univ_info table")

	// Create Table
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS univ_info (
			id SERIAL PRIMARY KEY,
			univId varchar(16) NOT NULL,
			univName varchar(32) NOT NULL,
			univFullName varchar(64),
			location varchar(32),
			address varchar(128),
			gender varchar(14),
			phone varchar(16),
			fax varchar(16),
			email varchar(32),
			website varchar(32),
			remark varchar(128),
			inUser varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP );
		`)

	utils.HandleErr(err, "Create univ_info table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create univ_info table")

	//  Insert Data
	stmt, err = db.Prepare(`
		INSERT INTO univ_info (
			id, univId, univName, univFullName, location, address, gender, phone, fax, email, website, remark) VALUES
			(1, 'univ_101', 'DGIST', '', '', '', '', '', '', '', '', ''),
			(2, 'univ_102', 'GIST', '', '', '', '', '', '', '', '', ''),
			(3, 'univ_103', 'KAIST', '', '', '', '', '', '', '', '', ''),
			(4, 'univ_104', 'UNIST', '', '', '', '', '', '', '', '', ''),
			(5, 'univ_105', '가야대', '', '', '', '', '', '', '', '', ''),
			(6, 'univ_106', '가천대', '', '', '', '', '', '', '', '', ''),
			(7, 'univ_107', '가톨릭관동대', '', '', '', '', '', '', '', '', ''),
			(8, 'univ_108', '가톨릭꽃동네대', '', '', '', '', '', '', '', '', ''),
			(9, 'univ_109', '가톨릭대', '', '', '', '', '', '', '', '', ''),
			(10, 'univ_110', '강남대', '', '', '', '', '', '', '', '', ''),
			(11, 'univ_111', '강원대(도계캠퍼스)', '', '', '', '', '', '', '', '', ''),
			(12, 'univ_112', '강원대(춘천)', '', '', '', '', '', '', '', '', ''),
			(13, 'univ_113', '건국대', '', '', '', '', '', '', '', '', ''),
			(14, 'univ_114', '건국대(글로컬)', '', '', '', '', '', '', '', '', ''),
			(15, 'univ_115', '건양대(메디컬캠퍼스)', '', '', '', '', '', '', '', '', ''),
			(16, 'univ_116', '경기대', '', '', '', '', '', '', '', '', ''),
			(17, 'univ_117', '경북대', '', '', '', '', '', '', '', '', ''),
			(18, 'univ_118', '경희대', '', '', '', '', '', '', '', '', ''),
			(19, 'univ_119', '계명대', '', '', '', '', '', '', '', '', ''),
			(20, 'univ_120', '고려대', '', '', '', '', '', '', '', '', ''),
			(21, 'univ_121', '공주대', '', '', '', '', '', '', '', '', ''),
			(22, 'univ_122', '광운대', '', '', '', '', '', '', '', '', ''),
			(23, 'univ_123', '국립강릉원주대(강릉캠퍼스)', '', '', '', '', '', '', '', '', ''),
			(24, 'univ_124', '국민대', '', '', '', '', '', '', '', '', ''),
			(25, 'univ_125', '나사렛대', '', '', '', '', '', '', '', '', ''),
			(26, 'univ_126', '남서울대', '', '', '', '', '', '', '', '', ''),
			(27, 'univ_127', '단국대(죽전)', '', '', '', '', '', '', '', '', ''),
			(28, 'univ_128', '단국대(천안)', '', '', '', '', '', '', '', '', ''),
			(29, 'univ_129', '대구가톨릭대', '', '', '', '', '', '', '', '', ''),
			(30, 'univ_130', '대구대', '', '', '', '', '', '', '', '', ''),
			(31, 'univ_131', '대구한의대', '', '', '', '', '', '', '', '', ''),
			(32, 'univ_132', '덕성여대', '', '', '', '', '', '', '', '', ''),
			(33, 'univ_133', '동국대', '', '', '', '', '', '', '', '', ''),
			(34, 'univ_134', '동덕여대', '', '', '', '', '', '', '', '', ''),
			(35, 'univ_135', '명지대', '', '', '', '', '', '', '', '', ''),
			(36, 'univ_136', '백석대', '', '', '', '', '', '', '', '', ''),
			(37, 'univ_137', '삼육대', '', '', '', '', '', '', '', '', ''),
			(38, 'univ_138', '상명대', '', '', '', '', '', '', '', '', ''),
			(39, 'univ_139', '상명대(천안)', '', '', '', '', '', '', '', '', ''),
			(40, 'univ_140', '상지대', '', '', '', '', '', '', '', '', ''),
			(41, 'univ_141', '서강대', '', '', '', '', '', '', '', '', ''),
			(42, 'univ_142', '서울과기대', '', '', '', '', '', '', '', '', ''),
			(43, 'univ_143', '서울대', '', '', '', '', '', '', '', '', ''),
			(44, 'univ_144', '서울시립대', '', '', '', '', '', '', '', '', ''),
			(45, 'univ_145', '서울여대', '', '', '', '', '', '', '', '', ''),
			(46, 'univ_146', '선문대', '', '', '', '', '', '', '', '', ''),
			(47, 'univ_147', '성균관대', '', '', '', '', '', '', '', '', ''),
			(48, 'univ_148', '성신여대', '', '', '', '', '', '', '', '', ''),
			(49, 'univ_149', '세종대', '', '', '', '', '', '', '', '', ''),
			(50, 'univ_150', '수원대', '', '', '', '', '', '', '', '', ''),
			(51, 'univ_151', '숙명여대', '', '', '', '', '', '', '', '', ''),
			(52, 'univ_152', '순천향대', '', '', '', '', '', '', '', '', ''),
			(53, 'univ_153', '숭실대', '', '', '', '', '', '', '', '', ''),
			(54, 'univ_154', '아주대', '', '', '', '', '', '', '', '', ''),
			(55, 'univ_155', '연세대', '', '', '', '', '', '', '', '', ''),
			(56, 'univ_156', '용인대', '', '', '', '', '', '', '', '', ''),
			(57, 'univ_157', '이화여대', '', '', '', '', '', '', '', '', ''),
			(58, 'univ_158', '인천가톨릭대', '', '', '', '', '', '', '', '', ''),
			(59, 'univ_159', '인천대', '', '', '', '', '', '', '', '', ''),
			(60, 'univ_160', '인하대', '', '', '', '', '', '', '', '', ''),
			(61, 'univ_161', '중앙대', '', '', '', '', '', '', '', '', ''),
			(62, 'univ_162', '차의과학대', '', '', '', '', '', '', '', '', ''),
			(63, 'univ_163', '충남대', '', '', '', '', '', '', '', '', ''),
			(64, 'univ_164', '충북대', '', '', '', '', '', '', '', '', ''),
			(65, 'univ_165', '케이씨대', '', '', '', '', '', '', '', '', ''),
			(66, 'univ_166', '평택대', '', '', '', '', '', '', '', '', ''),
			(67, 'univ_167', '한국교통대', '', '', '', '', '', '', '', '', ''),
			(68, 'univ_168', '한국기술교대', '', '', '', '', '', '', '', '', ''),
			(69, 'univ_169', '한국산업기술대', '', '', '', '', '', '', '', '', ''),
			(70, 'univ_170', '한국외대', '', '', '', '', '', '', '', '', ''),
			(71, 'univ_171', '한국외대(글로벌)', '', '', '', '', '', '', '', '', ''),
			(72, 'univ_172', '한성대', '', '', '', '', '', '', '', '', ''),
			(73, 'univ_173', '한신대', '', '', '', '', '', '', '', '', ''),
			(74, 'univ_174', '한양대', '', '', '', '', '', '', '', '', ''),
			(75, 'univ_175', '한양대ERICA', '', '', '', '', '', '', '', '', ''),
			(76, 'univ_176', '호서대', '', '', '', '', '', '', '', '', ''),
			(77, 'univ_177', '홍익대', '', '', '', '', '', '', '', '', ''),
			(78, 'univ_999', '', '', '', '', '', '', '', '', '', '');
		`)
	
	utils.HandleErr(err, "Insert Data into univ_info table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Insert Data into univ_info table")
}