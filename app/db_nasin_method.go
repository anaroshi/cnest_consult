package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_nasin_method(db *sql.DB) {

	// Drop table
	statement, err := db.Prepare(`DROP TABLE IF EXISTS nasin_method;`)

	utils.HandleErr(err, "Drop nasin_method table")
	_, err = statement.Exec()
	utils.HandleErr(err, "Exec for Drop nasin_method table")


	// Create table
	statement, err = db.Prepare(
		`CREATE TABLE nasin_method (
			id SERIAL PRIMARY KEY,
			nasinmetId varchar(16) NOT NULL,
			nasinmetName varchar(500) NOT NULL,
			sort integer,
			flag integer,
			inUser varchar(24),
			inDate TIMESTAMP,
			reUser varchar(24),
			reDate TIMESTAMP
		);`)

	utils.HandleErr(err, "Create nasin_method table")
	_, err = statement.Exec()
	utils.HandleErr(err, "Exec for Create nasin_method table")


	// Insert Data
	stmt, err := db.Prepare(
		`INSERT INTO nasin_method (
			id, nasinmetId, nasinmetName, sort, flag) VALUES		
			(1, 'namt_1001', '교과 10, 가중치 없음, 국영수사/과 각 교과별 상위 3과목씩 총 12과목, 석차등급', 0, 0),
			(2, 'namt_1002', '교과 10, 국수영 + 택1(사회/과학) 교과별 상위 3과목, 학년별가중치 없음, 석차등급', 0, 0),
			(3, 'namt_1003', '교과 10, 국수영사과한국사 전과목, 석차등급 및 성취도', 0, 0),
			(4, 'namt_1004', '교과 100 , 가중치 없음, 국수영사과한국사, 석차등급 (단, 2022년 졸업예정자의 경우 진로선택과목 성취도 상위 3과목 반영)', 0, 0),
			(5, 'namt_1005', '교과 100, (공통 및 일반선택) 전과목, 성착등급, (진로선택 및 전문교과) 학업수월성(10점) + 학업충실성(10점), 정성평가', 0, 0),
			(6, 'namt_1006', '교과 100, (공통/일반선택)국(15)수(35)영(25)과(25) 전과목, 석차등급 (진로선택)성취도 등급', 0, 0),
			(7, 'namt_1007', '교과 100, (공통/일반선택)국(15)수(35)영(25)사과(25) 전과목, 석차등급 (진로선택)성취도 등급', 0, 0),
			(8, 'namt_1008', '교과 100, (공통/일반선택)국(15)수(35)영(35)사(15) 전과목, 석차등급 (진로선택)성취도 등급', 0, 0),
			(9, 'namt_1009', '교과 100, (공통/일반선택)국(35)수(15)영(35)사(15) 전과목, 석차등급 (진로선택)성취도 등급', 0, 0),
			(10, 'namt_1010', '교과 100, (반영교과 A)국수영사과, 공통(30) 일반선택(50)-석차등급(50%) 및 Z점수(50%), 진로선택(20)-성취도, (반영교과 B)반영교과 A를 제외한 기타 과목, 최대 5점 감점, ', 0, 0),
			(11, 'namt_1011', '교과 100, 1학년(20) 2~3학년(80), 국(20)수(30)영(30)과(20) 교과 전과목, 석차등급', 0, 0),
			(12, 'namt_1012', '교과 100, 1학년(20) 2~3학년(80), 국(30)수(20)영(30)과(20) 교과 전과목, 석차등급', 0, 0),
			(13, 'namt_1013', '교과 100, 1학년(30) 2·3학년(70), 국영수사/과 전 과목, 석차등급', 0, 0),
			(14, 'namt_1014', '교과 100, 가중치 없음, (고1)국수영사과, (고2·3)국영수과, 석차등급', 0, 0),
			(15, 'namt_1015', '교과 100, 가중치 없음, (고1)국수영사과, (고2·3)국영수사, 석차등급', 0, 0),
			(16, 'namt_1016', '교과 100, 가중치 없음, 국수영과 교과 전과목, 석차등급', 0, 0),
			(17, 'namt_1017', '교과 100, 가중치 없음, 국수영과 교과별 상위 4과목, 석차등급', 0, 0),
			(18, 'namt_1018', '교과 100, 가중치 없음, 국수영과 교과별 상위 5과목, 석차등급', 0, 0),
			(19, 'namt_1019', '교과 100, 가중치 없음, 국수영과 상위 3과목, 35:25:25:15, 석차등급 및 성취도 ※진로선택과목 성취도 A인 경우 2개 교과에 한하여 최하등급 1개 과목을 1개 등급 상향하여 적용', 0, 0),
			(20, 'namt_1020', '교과 100, 가중치 없음, 국수영과 상위 5과목, 35:25:25:15, 석차등급 및 성취도 ※진로선택과목 성취도 A인 경우 2개 교과에 한하여 최하등급 1개 과목을 1개 등급 상향하여 적용', 0, 0),
			(21, 'namt_1021', '교과 100, 가중치 없음, 국수영과 전과목, 35:25:25:15, 석차등급', 0, 0),
			(22, 'namt_1022', '교과 100, 가중치 없음, 국수영과 전과목, 석차등급', 0, 0),
			(23, 'namt_1023', '교과 100, 가중치 없음, 국수영사 교과 전과목, 석차등급', 0, 0),
			(24, 'namt_1024', '교과 100, 가중치 없음, 국수영사 교과별 상위 4과목, 석차등급', 0, 0),
			(25, 'namt_1025', '교과 100, 가중치 없음, 국수영사 상위 3과목, 35:25:25:15, 석차등급 및 성취도 ※진로선택과목 성취도 A인 경우 2개 교과에 한하여 최하등급 1개 과목을 1개 등급 상향하여 적용', 0, 0),
			(26, 'namt_1026', '교과 100, 가중치 없음, 국수영사 상위 5과목, 35:25:25:15, 석차등급 및 성취도 ※진로선택과목 성취도 A인 경우 2개 교과에 한하여 최하등급 1개 과목을 1개 등급 상향하여 적용', 0, 0),
			(27, 'namt_1027', '교과 100, 가중치 없음, 국수영사 전과목, 석차등급', 0, 0),
			(28, 'namt_1028', '교과 100, 가중치 없음, 국수영사/과 교과별 상위 3과목씩 총 12과목, 석차등급', 0, 0),
			(29, 'namt_1029', '교과 100, 가중치 없음, 국수영사/과 교과별 상위 5과목, 석차등급', 0, 0),
			(30, 'namt_1030', '교과 100, 가중치 없음, 국수영사/과 전과목, 석차등급', 0, 0),
			(31, 'namt_1031', '교과 100, 가중치 없음, 국수영사과 교과 구분없이 석차등급 상위 3과목, 총 3과목 반영, 석차등급', 0, 0),
			(32, 'namt_1032', '교과 100, 가중치 없음, 국수영사과 전과목, 석차등급', 0, 0),
			(33, 'namt_1033', '교과 100, 가중치 없음, 국수영사과, 석차등급', 0, 0),
			(34, 'namt_1034', '교과 100, 가중치 없음, 국수영사과한국사 교과 전과목, 석차등급', 0, 0),
			(35, 'namt_1035', '교과 100, 가중치 없음, 국수영사과한국사 교과 중 상위 12과목, 과학 2과목 필수, 석차등급', 0, 0),
			(36, 'namt_1036', '교과 100, 가중치 없음, 국수영사과한국사 전과목, (공통 및 일반선택) 석차등급, (진로선택) 상위 3과목, 성취도', 0, 0),
			(37, 'namt_1037', '교과 100, 가중치 없음, 국수영사과한국사 전과목, 석차등급', 0, 0),
			(38, 'namt_1038', '교과 100, 가중치 없음, 국수영사과한국사, 석차등급', 0, 0),
			(39, 'namt_1039', '교과 100, 가중치 없음, 국수영사과한국사, 진로선택과목 가산점, 석차등급 및 성취도', 0, 0),
			(40, 'namt_1040', '교과 100, 가중치 없음, 국수영탐 상위 3과목, 35:25:25:15, 석차등급 및 성취도 ※진로선택과목 성취도 A인 경우 2개 교과에 한하여 최하등급 1개 과목을 1개 등급 상향하여 적용', 0, 0),
			(41, 'namt_1041', '교과 100, 가중치 없음, 국수영탐 상위 5과목, 35:25:25:15, 석차등급 및 성취도 ※진로선택과목 성취도 A인 경우 2개 교과에 한하여 최하등급 1개 과목을 1개 등급 상향하여 적용', 0, 0),
			(42, 'namt_1042', '교과 100, 가중치 없음, 국어/수학 3과목 + 영어 3과목 + 사회/과학 3과목, 총 9과목, 석차등급', 0, 0),
			(43, 'namt_1043', '교과 100, 가중치 없음, 국영 + 수사과 중 선택 1교과 전과목, 석차등급', 0, 0),
			(44, 'namt_1044', '교과 100, 가중치 없음, 국영 교과별 상위 4과목, 석차등급', 0, 0),
			(45, 'namt_1045', '교과 100, 가중치 없음, 국영수 6개 + 사과 2개, 총 8개 과목, 진로선택과목은 최대 2개, 석차등급', 0, 0),
			(46, 'namt_1046', '교과 100, 가중치 없음, 국영수과 교과 전과목, 석차등급', 0, 0),
			(47, 'namt_1047', '교과 100, 가중치 없음, 국영수과 전과목, 석차등급', 0, 0),
			(48, 'namt_1048', '교과 100, 가중치 없음, 국영수사 교과 전과목, 석차등급', 0, 0),
			(49, 'namt_1049', '교과 100, 가중치 없음, 국영수사/과 교과 전과목, 석차등급', 0, 0),
			(50, 'namt_1050', '교과 100, 가중치 없음, 국영수사/과 교과 중 상위 3개 교과의 각 상위 4개 과목(단, 총 12단위 이상 이수), 석차등급', 0, 0),
			(51, 'namt_1051', '교과 100, 가중치 없음, 국영수사/과 중 2개 교과, 석차등급', 0, 0),
			(52, 'namt_1052', '교과 100, 가중치 없음, 국영수사/과 중 3개 교과 전과목, 석차등급', 0, 0),
			(53, 'namt_1053', '교과 100, 가중치 없음, 국영수사과 교과 중 학년별 4과목, 석차등급', 0, 0),
			(54, 'namt_1054', '교과 100, 가중치 없음, 국영수사과 전 교과목, 석차등급', 0, 0),
			(55, 'namt_1055', '교과 100, 가중치 없음, 국영수사과한국사 교과 중 15개 과목, 석차등급', 0, 0),
			(56, 'namt_1056', '교과 100, 가중치 없음, 국영수사과한국사 전과목, 석차등급', 0, 0),
			(57, 'namt_1057', '교과 100, 가중치 없음, 국영수사과한국사 전과목, 석차등급 및 성취도 ※진로 선택 과목은 최대 3과목 반영', 0, 0),
			(58, 'namt_1058', '교과 100, 가중치 없음, 국영수사과한국사 전과목, 석차등급, 3-1학기까지 반영', 0, 0),
			(59, 'namt_1059', '교과 100, 가중치 없음, 국영수사과한국사 전과목, 석차등급, 3-1학기까지 반영, 최고 30 ~ 최저 19.2', 0, 0),
			(60, 'namt_1060', '교과 100, 가중치 없음, 국영수사교과 전과목, 석차등급', 0, 0),
			(61, 'namt_1061', '교과 100, 가중치 없음, 석차등급 평가된 전 교과목 + 성취등급 평가된 최대 3과목', 0, 0),
			(62, 'namt_1062', '교과 100, 가중치 없음, 수학+과학 교과별 상위 5과목, 석차등급', 0, 0),
			(63, 'namt_1063', '교과 100, 가중치 없음, 수학+영어 교과별 상위 5과목, 석차등급', 0, 0),
			(64, 'namt_1064', '교과 100, 가중치 없음, 전교과 전과목, 석차등급', 0, 0),
			(65, 'namt_1065', '교과 100, 가중치 없음, 전학년 전과목, 석차등급', 0, 0),
			(66, 'namt_1066', '교과 100, 가중치 없음, 전학년 전과목, 석차등급 및 성취도', 0, 0),
			(67, 'namt_1067', '교과 100, 국(20) 수(30) 영(30) 과(20), 석차등급', 0, 0),
			(68, 'namt_1068', '교과 100, 국(20)수(30)영(30)과(20) 교과 전과목, 석차등급', 0, 0),
			(69, 'namt_1069', '교과 100, 국(20)수(30)영(30)과(20) 교과 전과목, 학년별 가중치 없음, 석차등급', 0, 0),
			(70, 'namt_1070', '교과 100, 국(20)수(30)영(30)사(20) 교과 전과목, 석차등급', 0, 0),
			(71, 'namt_1071', '교과 100, 국(30) 수(30) 영(30) 사/과(10), 석차등급', 0, 0),
			(72, 'namt_1072', '교과 100, 국(30)수(20)영(30)사(20) 교과 전과목, 석차등급', 0, 0),
			(73, 'namt_1073', '교과 100, 국(30)수(20)영(30)사(20) 교과 전과목, 학년별 가중치 없음, 석차등급', 0, 0),
			(74, 'namt_1074', '교과 100, 국(30)수(30)영(30)사(10) 교과 전과목, 석차등급', 0, 0),
			(75, 'namt_1075', '교과 100, 국수영과 각 교과별 상위 5과목, 점수가 높은 순으로 30:30:25:15, 석차등급', 0, 0),
			(76, 'namt_1076', '교과 100, 국수영사 각 교과별 상위 5과목, 점수가 높은 순으로 30:30:25:15, 석차등급', 0, 0),
			(77, 'namt_1077', '교과 100, 국어교과 전과목, 석차등급', 0, 0),
			(78, 'namt_1078', '교과 100, 국영수과, (공통 및 일반선택)(90) 석차등급 + (진로선택)(10) 성취도', 0, 0),
			(79, 'namt_1079', '교과 100, 국영수과, 공통(40) + 일반선택(60) + 진로선택 2과목 가산점(3), 석차등급 및 성취도', 0, 0),
			(80, 'namt_1080', '교과 100, 국영수사, (공통 및 일반선택)(90) 석차등급 + (진로선택)(10) 성취도', 0, 0),
			(81, 'namt_1081', '교과 100, 수학교과 전과목, 석차등급', 0, 0),
			(82, 'namt_1082', '교과 100, 학년별 30:40:30, 국영수과 전 과목, 석차등급', 0, 0),
			(83, 'namt_1083', '교과 100, 학년별 가중치 20:40:40, 국수영과사 교과 중 1학년 8개, 2학년 8개, 3학년 8개 과목, 석차등급', 0, 0),
			(84, 'namt_1084', '교과 100, 학년별 가중치 30:30:40, 전교과 전과목, 석차등급', 0, 0),
			(85, 'namt_1085', '교과 100, 학년별 국영수사/과(한국사포함) 교과별 1개 과목, 총 12개 과목, 석차등급)', 0, 0),
			(86, 'namt_1086', '교과 100, 학년별가중치 30:35:35, 국수영사과 교과 중 3개 교과, 석차등급', 0, 0),
			(87, 'namt_1087', '교과 20, 가중치 없음, 국수영과 교과 상위 10개 과목, 석차등급', 0, 0),
			(88, 'namt_1088', '교과 20, 가중치 없음, 국수영사 교과 상위 10개 과목, 석차등급', 0, 0),
			(89, 'namt_1089', '교과 20, 가중치 없음, 국수영사과한국사 중 상위 10과목, 이수단위 미적용, 석차등급', 0, 0),
			(90, 'namt_1090', '교과 20, 가중치 없음, 국영수사/과 교과 중 상위 3개 교과의 각 상위 4개 과목(단, 총 12단위 이상 이수), 석차등급', 0, 0),
			(91, 'namt_1091', '교과 21, 가중치 없음, (공통·일반선택)국영, 석차등급 (진로선택)국영예술체육 상위 2개 과목, 성취도', 0, 0),
			(92, 'namt_1092', '교과 21, 가중치 없음, (공통·일반선택)국영수과, 석차등급 (진로선택)국영수과 상위 4개 과목, 성취도', 0, 0),
			(93, 'namt_1093', '교과 21, 가중치 없음, (공통·일반선택)국영수사, 석차등급 (진로선택)국영수사 상위 2개 과목, 성취도', 0, 0),
			(94, 'namt_1094', '교과 27, 국(30)수(20)영(30)사(20) 전과목, 석차등급', 0, 0),
			(95, 'namt_1095', '교과 27, 학년별 가중치 30:40:30, 국수영과 전과목, 학년별가중치 20:40:40, 석차등급', 0, 0),
			(96, 'namt_1096', '교과 27, 학년별 가중치 30:40:30, 국수영사 전과목, 학년별가중치 20:40:40, 석차등급', 0, 0),
			(97, 'namt_1097', '교과 30 , 가중치 없음, 국수영사과한국사, 석차등급 (단, 2022년 졸업예정자의 경우 진로선택과목 성취도 상위 3과목 반영)', 0, 0),
			(98, 'namt_1098', '교과 30, 가중치 없음, 국수영과 교과 전과목, 석차등급', 0, 0),
			(99, 'namt_1099', '교과 30, 가중치 없음, 국수영과 전과목, 석차등급', 0, 0),
			(100, 'namt_1100', '교과 30, 가중치 없음, 국수영사 교과 전과목, 석차등급', 0, 0),
			(101, 'namt_1101', '교과 30, 가중치 없음, 국수영사 전과목, 석차등급', 0, 0),
			(102, 'namt_1102', '교과 30, 가중치 없음, 국수영사/과 교과별 상위 3과목씩 총 12과목, 석차등급', 0, 0),
			(103, 'namt_1103', '교과 30, 가중치 없음, 국수영사과 상위 30단위, (공통/일반선택) 석차등급 + (진로선택/전문교과Ⅰ) 최대 3단위, 성취도', 0, 0),
			(104, 'namt_1104', '교과 30, 가중치 없음, 국수영사과한국사 교과 전과목, 석차등급', 0, 0),
			(105, 'namt_1105', '교과 30, 가중치 없음, 국수영사과한국사, 석차등급', 0, 0),
			(106, 'namt_1106', '교과 30, 가중치 없음, 국영사/수영과 중 학기당 교과별 1과목, 총 15과목, 석차등급', 0, 0),
			(107, 'namt_1107', '교과 30, 가중치 없음, 국영수과 교과 전과목, 석차등급', 0, 0),
			(108, 'namt_1108', '교과 30, 가중치 없음, 국영수사 교과 전과목, 석차등급', 0, 0),
			(109, 'namt_1109', '교과 30, 가중치 없음, 전교과 전과목, 석차등급', 0, 0),
			(110, 'namt_1110', '교과 30, 국(20)수(30)영(30)사(20) 교과 전과목, 석차등급', 0, 0),
			(111, 'namt_1111', '교과 30, 국(30)수(20)영(30)사(20) 교과 전과목, 석차등급', 0, 0),
			(112, 'namt_1112', '교과 30, 국(30)수(30)영(30)사(10) 교과 전과목, 석차등급', 0, 0),
			(113, 'namt_1113', '교과 40, 가중치 없음, 국수영사 전과목,  석차등급', 0, 0),
			(114, 'namt_1114', '교과 40, 가중치 없음, 국수영통사통과사회, 석차등급', 0, 0),
			(115, 'namt_1115', '교과 40, 국수영과 각 교과별 상위 5과목, 점수가 높은 순으로 30:30:25:15, 석차등급', 0, 0),
			(116, 'namt_1116', '교과 40, 국수영사 각 교과별 상위 5과목, 점수가 높은 순으로 30:30:25:15, 석차등급', 0, 0),
			(117, 'namt_1117', '교과 60, 가중치 없음, 국수영사과 교과 중 상위 12과목, 석차등급', 0, 0),
			(118, 'namt_1118', '교과 60, 국수영과 각 교과별 상위 5과목, 점수가 높은 순으로 30:30:25:15, 석차등급', 0, 0),
			(119, 'namt_1119', '교과 60, 국수영과한국사 상위 10과목, 이수단위 미적용, 석차등급', 0, 0),
			(120, 'namt_1120', '교과 60, 국수영사 각 교과별 상위 5과목, 점수가 높은 순으로 30:30:25:15, 석차등급', 0, 0),
			(121, 'namt_1121', '교과 60, 국수영사한국사 상위 10과목, 이수단위 미적용, 석차등급', 0, 0),
			(122, 'namt_1122', '교과 63, 가중치 없음, 국어/수학 3과목 + 영어 3과목 + 사회/과학 3과목, 총 9과목, 석차등급', 0, 0),
			(123, 'namt_1123', '교과 70, 가중치 없음, (공통/일반선택)국수영 중 2개 교과 + 사과 중 1개 교과, 합 3개교과, 석차등급, (진로선택)3개 과목, 성취도', 0, 0),
			(124, 'namt_1124', '교과 70, 가중치 없음, 국수영과 전과목, (공통/일반선택)(90), 석차등급 + (진로선택)(10), 성취도', 0, 0),
			(125, 'namt_1125', '교과 70, 가중치 없음, 국수영과한국사 전 과목 + 통합사회 + 진로선택(2과목), 석차등급 및 성취도', 0, 0),
			(126, 'namt_1126', '교과 70, 가중치 없음, 국수영사 전과목, (공통/일반선택)(90), 석차등급 + (진로선택)(10), 성취도', 0, 0),
			(127, 'namt_1127', '교과 70, 가중치 없음, 국수영사과한국사 교과 전과목, 석차등급', 0, 0),
			(128, 'namt_1128', '교과 70, 학년별 가중치 30:30:40, 국수영사과 교과별 학기당 3과목, 총 15과목, 석차등급', 0, 0),
			(129, 'namt_1129', '교과 70, 학년별가중치 30:35:35, 국수영사과 교과 중 3개 교과, 석차등급', 0, 0),
			(130, 'namt_1130', '교과 80, 가중치 없음, (공통+일반선택)국수영 중 10개과목 + 사과 중 2과목, 총 12과목, 석차등급, (진로선택)국수영사과 중 2개 과목, 성취도', 0, 0),
			(131, 'namt_1131', '교과 80, 가중치 없음, (공통·일반선택)국영, 석차등급 (진로선택)국영예술체육 상위 2개 과목, 성취도', 0, 0),
			(132, 'namt_1132', '교과 80, 가중치 없음, (공통·일반선택)국영수과, 석차등급 (진로선택)국영수과 상위 4개 과목, 성취도', 0, 0),
			(133, 'namt_1133', '교과 80, 가중치 없음, (공통·일반선택)국영수사, 석차등급 (진로선택)국영수사 상위 2개 과목, 성취도', 0, 0),
			(134, 'namt_1134', '교과 80, 가중치 없음, 국수영과 교과별 상위 5과목, 석차등급', 0, 0),
			(135, 'namt_1135', '교과 80, 가중치 없음, 국수영사 상위 5과목, 35:25:25:15, 석차등급 및 성취도 ※진로선택과목 성취도 A인 경우 2개 교과에 한하여 최하등급 1개 과목을 1개 등급 상향하여 적용', 0, 0),
			(136, 'namt_1136', '교과 80, 가중치 없음, 국수영사/과 교과별 상위 5과목, 석차등급', 0, 0),
			(137, 'namt_1137', '교과 80, 가중치 없음, 국수영사과 전과목, (공통/일반선택)(90) 석차등급 + (진로선택/전문교과Ⅰ)(10) 성취도', 0, 0),
			(138, 'namt_1138', '교과 80, 가중치 없음, 국수영사과한국사 교과 중 상위 12과목, 과학 2과목 필수, 석차등급', 0, 0),
			(139, 'namt_1139', '교과 80, 가중치 없음, 전교과, 석차등급 및 성취도', 0, 0),
			(140, 'namt_1140', '교과 80, 국영수사/과 중 3교과, 반영교과 점수 높은 순 50%:35%:15%, 석차등급', 0, 0),
			(141, 'namt_1141', '교과 80, 학년별 가중치 30:30:40, 국수영사/과 교과 전과목, 석차등급', 0, 0),
			(142, 'namt_1142', '교과 80, 학년별 국영수사/과(한국사포함) 교과별 1개 과목, 총 12개 과목, 석차등급)', 0, 0),
			(143, 'namt_1143', '교과 90, 가중치 없음, 국수영통사통과과학, 석차등급', 0, 0),
			(144, 'namt_1144', '교과 90, 가중치 없음, 국수영통사통과사회, 석차등급', 0, 0),
			(145, 'namt_1145', '교과 90, 가중치 없음, 전교과 전과목, 석차등급', 0, 0),
			(146, 'namt_1146', '교과 90, 가중치없음, 국수영사과 교과 중 15개 과목, 석차등급', 0, 0),
			(147, 'namt_1147', '교과 90, 국(20)수(30)영(30)과(20) 전과목, 석차등급', 0, 0),
			(148, 'namt_1148', '교과 90, 국(30)수(20)영(30)사(20) 전과목, 석차등급', 0, 0),
			(149, 'namt_1149', '교과 90, 국수영사과한국사 전과목, 석차등급 및 성취도', 0, 0),
			(150, 'namt_1150', '교과 90, 학년별 가중치 30:30:40, 국수영사과 교과별 학기당 3과목, 총 15과목, 석차등급', 0, 0),
			(151, 'namt_1151', '교과 90, 학년별 가중치 30:40:30, 국수영과 전과목, 학년별가중치 20:40:40, 석차등급', 0, 0),
			(152, 'namt_1152', '교과 90, 학년별 가중치 30:40:30, 국수영사 전과목, 학년별가중치 20:40:40, 석차등급', 0, 0),
			(153, 'namt_1153', '교과 95, 가중치 없음, 국영사/수영과 중 학기당 교과별 1과목, 총 15과목, 석차등급', 0, 0),
			(154, 'namt_1154', '정성평가', 0, 0),
			(155, 'namt_1155', '해당없음', 0, 0),
			(156, 'namt_9999', '', 0, 0)
		;`)
	utils.HandleErr(err, "Insert Data into nasin_method table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Insert Data into nasin_method table")
}