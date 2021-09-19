package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db4_apply_qualification(db *sql.DB) {

	// Drop table
	statement, err := db.Prepare(`DROP TABLE IF EXISTS apply_qualification;`)

	utils.HandleErr(err, "drop table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table11 exec")


	// Create table
	statement, err = db.Prepare(
		`CREATE TABLE apply_qualification (
				id int NOT NULL,
				qualiId varchar(16) NOT NULL,
				qualiName varchar(500) NOT NULL,
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
		`INSERT INTO apply_qualification (
		id, qualiId, qualiName, sort, flag) VALUES		
		(1, 'qual_1001', '2013년 2월 이후 국내 고교 졸업(예정)자로서 국내 고교 교육과정에서 3학기 이상의 성적을 취득한 자', 0, 0),
		(2, 'qual_1002', '2016년 이후 국내 고교 졸업(예정)자 / 검정고시 지원가 / 특성화고 지원불가', 0, 0),
		(3, 'qual_1003', '2017년 1월 이후 국내 고교 졸업(예정)자로서 3개 학기 이상 있는 자로서 대학교에서 인정하는 대회 운문, 산문 부문 상위 3위 이내 입상자', 0, 0),
		(4, 'qual_1004', '2017년 2월 이후 국내 고교 졸업(예정)자', 0, 0),
		(5, 'qual_1005', '2017년 2월 이후 국내 고교 졸업(예정)자 중 출신학교장의 추천을 받은자', 0, 0),
		(6, 'qual_1006', '2017년 2월 이후 국내 고교 졸업(예정)자로서 3개 학기 이상의 본교 지정교과목 석차 등급이 있는 자', 0, 0),
		(7, 'qual_1007', '2017년 2월 이후 국내 고교 졸업(예정)자로서 소속 고등학교장의 추천을 받은 자이며, 5개학기 교과성적 산출이 가능한 자, 고교별 추천인원 10명 이내', 0, 0),
		(8, 'qual_1008', '2017년 2월 이후 국내 고교 졸업(예정)자로서 소속 고등학교장의 추천을 받은 자이며, 군 인사법에 저촉되지 아니한 자이며, 5개학기 교과성적 산출이 가능한 자, 고교별 추천인원 10명 이내', 0, 0),
		(9, 'qual_1009', '2017년 2월 이후 국내 고교 졸업(예정)자로서 소속 학교장 추천을 받은 자, 고교별 3학년 입학정원의 10%이내', 0, 0),
		(10, 'qual_1010', '2017년 2월 이후 국내 고교 졸업(예정)자로서 학교생활기록부에 5개학기의 교육과정을 이수하고 교과성적 산출이 가능한 자', 0, 0),
		(11, 'qual_1011', '2017년 2월 이후 국내 고교 졸업(예정)자인 여자로서 3개 학기 이상 성적 취득자', 0, 0),
		(12, 'qual_1012', '2017년 2월 이후 국내 고교 졸업(예정)자인 여자로서 5개 학기 이상 성적 취득자 중 학교장의 추천을 받은 자, 고교별 추천 인원 제한 없음', 0, 0),
		(13, 'qual_1013', '2017년 이후 국내 고교 졸업(예정)자 또는 검정고시 합격자', 0, 0),
		(14, 'qual_1014', '2017년 이후 국내 고교 졸업(예정)자 중 국내 고교에서 3학기 이상 이수하고 소속 고등학교장의 추천을 받은 자, 고교별 추천 인원 제한 없음', 0, 0),
		(15, 'qual_1015', '2018년 2월 이후 고교 졸업(예정)자', 0, 0),
		(16, 'qual_1016', '2019년 2월 이후 국내 정규 고교 졸업(예정)자로서 통산 3개 학기 이상 국내 고교 성적 취득자이며 군인사법 제10조(결격사유 등) 제2항에 저촉되지 아니한 자로서 임관일 기준 만 20세 이상 27세 이하인 자', 0, 0),
		(17, 'qual_1017', '2020년 1월 이후 국내 고교 졸업(예정)자로서 3학년 1학기까지 전 학기를 모두 이수한 자', 0, 0),
		(18, 'qual_1018', '2020년 1월 이후 국내 고교 졸업(예정)자로서 소속 고등학교장의 추천을 받은 자', 0, 0),
		(19, 'qual_1019', '2020년 2월 이후 고교 졸업(예정)자로서 SW분야 우수성과 성장잠재력을 보여줄 수 있는 자', 0, 0),
		(20, 'qual_1020', '2020년 2월 이후 국내 고교 졸업(예정)자', 0, 0),
		(21, 'qual_1021', '2021년 2월 이후 고교 졸업(예정)자로, 국내 고교에서 3학기 이상 이수한 자, 고교별 5명 이내', 0, 0),
		(22, 'qual_1022', '2021년 2월 이후 국내 고교 졸업(예정)자, 고교별 재적 여학생 수의 5%이내(최대 10명)', 0, 0),
		(23, 'qual_1023', '2021년 2월 이후 국내 고교 졸업(예정)자로서 5개 학기 이상 국내 고교 성적 취득 및 출신 고등학교장의 추천을 받은 자', 0, 0),
		(24, 'qual_1024', '2021년 2월 이후 국내 정규 고교 졸업(예정)자로서 통산 3개학기 이상 국내 고교 성적 취득자', 0, 0),
		(25, 'qual_1025', '2021년 이후 국내 고교 졸업(예정)자로서 소속 고등학교 추천을 받은 자, 고교별 20명 이내', 0, 0),
		(26, 'qual_1026', '2022년 2월 고교 졸업예정인 국내 일반고, 특성화고, 자율고 재학생으로 추천기준에 따라 소속 학교장이 추천한 자, 고교별 최대 2명', 0, 0),
		(27, 'qual_1027', '2022년 2월 국내 고교 졸업예정자, 학교별 추천 2명 이내', 0, 0),
		(28, 'qual_1028', '2022년 2월 졸업예정자, 3학기 이상 이수하고 학생부 교과성적이 산출 가능한 출신 고등학교의 추천을 받은 자', 0, 0),
		(29, 'qual_1029', '2022년 국내 고교 졸업예정자 중 5학기 이상 교과성적이 기재되어 있는 자로 출신 고등학교장의 추천을 받은 자, 3학년 재적 학생수의 4%까지 추천할 수 있음', 0, 0),
		(30, 'qual_1030', '2022학년도 수능 응시자', 0, 0),
		(31, 'qual_1031', '강원도 소재 고교에 입학하여 전 교육과정을 이수한 졸업(예정)자', 0, 0),
		(32, 'qual_1032', '고교 졸업(예정)자 또는 동등학력 인정된 자', 0, 0),
		(33, 'qual_1033', '고교 졸업(예정)자', 0, 0),
		(34, 'qual_1034', '고교 졸업(예정)자 *단, 국내 학교생활기록부가 없는 자는 지원 불가', 0, 0),
		(35, 'qual_1035', '고교 졸업(예정)자 / 동등학력 이상', 0, 0),
		(36, 'qual_1036', '고교 졸업(예정)자 / 동등학력 이상 / 국내 학생부 3학기 이상', 0, 0),
		(37, 'qual_1037', '고교 졸업(예정)자 또는 검정고시 합격자', 0, 0),
		(38, 'qual_1038', '고교 졸업(예정)자 또는 고교 졸업이상의 학력 인정되는 자', 0, 0),
		(39, 'qual_1039', '고교 졸업(예정)자 또는 동등이상의 학력 인정되는 자', 0, 0),
		(40, 'qual_1040', '고교 졸업(예정)자 또는 동등이상의 학력소지자', 0, 0),
		(41, 'qual_1041', '고교 졸업(예정)자 또는 동등이상의 학력소지자로서 남을 도울 수 있는 인성을 갖춘자', 0, 0),
		(42, 'qual_1042', '고교 졸업(예정)자 또는 동등학력 이상 인정', 0, 0),
		(43, 'qual_1043', '고교 졸업(예정)자 또는 동등학력 이상 인정된 여자', 0, 0),
		(44, 'qual_1044', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자', 0, 0),
		(45, 'qual_1045', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 3개 학기 이상 있는 자', 0, 0),
		(46, 'qual_1046', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 과학 또는 수학분야에 우수한 역량 및 활동 실적을 갖춘 자', 0, 0),
		(47, 'qual_1047', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 국제학분야에 성장잠재력을 갖추고 영어강의 수강이 가능한 자', 0, 0),
		(48, 'qual_1048', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 군 인사법 제10조(결격사유 등) 2항에 저촉되지 아니한 자로서 임관일 기준 만 20세 이상 27세 이하인 자', 0, 0),
		(49, 'qual_1049', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 다산형 인재 핵심역량(융복합 사고 역량, 실천적 창의 역량, 의사소통 역량, 글로벌 역량)에 강점이 있는 자', 0, 0),
		(50, 'qual_1050', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 수학 및 과학을 바탕으로 SW분야의 역량과 잠재력을 갖춘 자', 0, 0),
		(51, 'qual_1051', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 어학분야에 우수한 역량 및 활동 실적을 갖춘 자', 0, 0),
		(52, 'qual_1052', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 해당 고등학교장의 추천을 받은 사람', 0, 0),
		(53, 'qual_1053', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서, 소프트웨어 분야 재능이 있는 자', 0, 0),
		(54, 'qual_1054', '고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서, 소프트웨어 분야에 대한 재능과 열정을 가진 자', 0, 0),
		(55, 'qual_1055', '고교 졸업(예정)자 또는 동등학력 인정된 자', 0, 0),
		(56, 'qual_1056', '고교 졸업(예정)자 또는 동등학력 인정된 자로서 국내 고교 3학기 이상의 성적을 취득한 자이며 학교장 또는 교사 등의 추천을 받은 자', 0, 0),
		(57, 'qual_1057', '고교 졸업(예정)자 또는 동등학력 인정된 자로서 잠재역량을 적극적으로 계발한 자', 0, 0),
		(58, 'qual_1058', '고교 졸업(예정)자 또는 동등학력 인정자', 0, 0),
		(59, 'qual_1059', '고교 졸업(예정)자 또는 동등학력 인정자로서 4차 산업혁명을 이끌어갈 우수 인재로서 교과, 연구, 기타 분야에서 탁월한 역량을 갖춘 자(예시: 국제올림피아드 수상, R&E연구 실적, SW / 로봇 / 산업디자인 및 설계 등)', 0, 0),
		(60, 'qual_1060', '고교 졸업(예정)자 또는 동등학력 인정자로서 소프트웨어 등 특정분야에 영재성을 가진 자', 0, 0),
		(61, 'qual_1061', '고교 졸업(예정)자 또는 동등학력으로 인정받을 수 있는 자로서 수학 및 과학 분야의 영재성을 가진 자 또는 과학기술 분야의 성장 잠재력이 높은 자', 0, 0),
		(62, 'qual_1062', '고교 졸업(예정)자 또는 졸업 이상의 학력이 인정되는 자', 0, 0),
		(63, 'qual_1063', '고교 졸업(예정)자 및 동등학력 인정', 0, 0),
		(64, 'qual_1064', '고교 졸업(예정)자 및 동등학력으로 인정받을 수 있는 자', 0, 0),
		(65, 'qual_1065', '고교 졸업(예정)자 및 동등학력으로 인정받을 수 있는 자로서 수학-과학 관련 학업역량이 우수하고 과학기술 연구 또는 창업을 통한 글로벌 리더로서의 성장 비전을 가진 자 또는 특정분야에 우수한 결과물을 산출한 이력이 있는 자', 0, 0),
		(66, 'qual_1066', '고교 졸업(예정)자 및 동등학력으로 인정받을 수 있는 자로서 울산광역시 소재 고교에서 교육 전 과정을 이수하고 고교 재학 전 기간 동안 주민등록상 울산광역시에 거주한 자', 0, 0),
		(67, 'qual_1067', '고교 졸업(예정)자 중 국내 고교 전 교육과정을 이수한 자', 0, 0),
		(68, 'qual_1068', '고교 졸업(예정)자 중 국내 고교 전 교육과정을 이수한 자로서 학교장 추천을 받은 자, 고교별 추천인원 제한 없음', 0, 0),
		(69, 'qual_1069', '고교 졸업(예정)자로서 대전-세종-충남-충북지역 소재 정규 고교에서 입학부터 졸업까지 3학년 전 과정을 이수한 자', 0, 0),
		(70, 'qual_1070', '고교 졸업(예정)자로서 입학일부터 졸업(예정)일까지 강원지원 고교에서 전 교육과정을 이수한 자', 0, 0),
		(71, 'qual_1071', '고교 졸업(예정)자로서 충남-충북-대전-세종 소재 고교에서 입학부터 졸업까지 3년간의 교육과정을 이수한 사람', 0, 0),
		(72, 'qual_1072', '고교 졸업(예정)자로서 학생이 입학일부터 졸업(예정)일까지 강원지역 고교에서 전 교육과정을 이수한 자', 0, 0),
		(73, 'qual_1073', '국내 고교 2014년 2월 이후 졸업(예정)자 또는 검정고시 합격자', 0, 0),
		(74, 'qual_1074', '국내 고교 졸업(예정)자', 0, 0),
		(75, 'qual_1075', '국내 고교 졸업(예정)자 또는 검정고시 합격자', 0, 0),
		(76, 'qual_1076', '국내 고교 졸업(예정)자 또는 고교 졸업 이상의 학력이 인정되는 자', 0, 0),
		(77, 'qual_1077', '국내 고교 졸업(예정)자 또는 동등이상 학력 인정된 자, ※국내 고등학교 학교생활기록부가 없는 자는 지원 불가능', 0, 0),
		(78, 'qual_1078', '국내 고교 졸업(예정)자 또는 동등학력 이상 인정된 자', 0, 0),
		(79, 'qual_1079', '국내 고교 졸업(예정)자 또는 동등학력 이상 인정된 자, 본당 주임신부의 추천을 받은 가톨릭신자, 교구 성직지망자는 지원할 수 없음', 0, 0),
		(80, 'qual_1080', '국내 고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서, TOEIC 900점 또는 IBT 95점 또는 TEPS 452점 이상인 자', 0, 0),
		(81, 'qual_1081', '국내 고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서, 新 JLPT N1급인 자', 0, 0),
		(82, 'qual_1082', '국내 고교 졸업(예정)자 또는 동등학력 이상 학력 소지자', 0, 0),
		(83, 'qual_1083', '국내 고교 졸업(예정)자 또는 동등학력 인정된 자', 0, 0),
		(84, 'qual_1084', '국내 고교 졸업(예정)자 또는 재외 한국학교 졸업(예정)자, 학교생활기록부가 없거나 반영교과 점수를 산출할 수 없는 자는 지원 불가', 0, 0),
		(85, 'qual_1085', '국내 고교 졸업(예정)자 및 5개 학기 이상 이수한 자', 0, 0),
		(86, 'qual_1086', '국내 고교 졸업(예정)자 및 동등학력 소지자', 0, 0),
		(87, 'qual_1087', '국내 고교 졸업(예정)자 중 국내 고교에서 3학기 이상 교육과정을 이수한 자', 0, 0),
		(88, 'qual_1088', '국내 고교 졸업(예정)자 중 소속 고등학교장의 추천을 받은 자, 학교별 추천인원 4명 이내', 0, 0),
		(89, 'qual_1089', '국내 고교 졸업(예정)자 중 충청권소재 고교 전교육과정 이수자', 0, 0),
		(90, 'qual_1090', '국내 고교 졸업(예정)자로 국제인재 성장잠재력을 보여 줄 수 있는 자', 0, 0),
		(91, 'qual_1091', '국내 고교 졸업(예정)자로 소속 고등학교장의 추천을 받은 자, 고교별 추천 인원은 10명 이내', 0, 0),
		(92, 'qual_1092', '국내 고교 졸업(예정)자로서 20187.10.1. 이후 4년제 대학 주최 컴퓨터프로그래밍 경시대회 및 정보올림피아드 입상한 자, 보안, 네트워크 분야 등은 제외', 0, 0),
		(93, 'qual_1093', '국내 고교 졸업(예정)자로서 3개 학기 이상 성적을 취득하고 소속 학교장의 추천을 받은 자, 학교별 추천인원 제한 없음', 0, 0),
		(94, 'qual_1094', '국내 고교 졸업(예정)자로서 3개 학기 이상 있고 소속(졸업) 고등학교장의 추천을 받은 자', 0, 0),
		(95, 'qual_1095', '국내 고교 졸업(예정)자로서 3학기 이상 성적을 취득한 자이며 소속 고등학교장의 추천을 받은 자', 0, 0),
		(96, 'qual_1096', '국내 고교 졸업(예정)자로서 4학기 이상 이수(예정)자로서 전국규모 대회 3위 이내 입상 실적이 있는 자', 0, 0),
		(97, 'qual_1097', '국내 고교 졸업(예정)자로서 국내 고교 3학기 이상 학생부(교과) 성적이 있는 자', 0, 0),
		(98, 'qual_1098', '국내 고교 졸업(예정)자로서 국내 고교 3학기 이상의 학교생활기록부가 있는 자 ※지원제한(불가)고교 확인 필', 0, 0),
		(99, 'qual_1099', '국내 고교 졸업(예정)자로서 국내 고교 3학기 이상의 학교생활기록부가 있는 자 또는 검정고시 합격자', 0, 0),
		(100, 'qual_1100', '국내 고교 졸업(예정)자로서 소속 고등학교장의 추천을 받은 자, 고교별 추천인원 7명 이내(인문계열 4명 이내, 자연계열 4명 이내)', 0, 0),
		(101, 'qual_1101', '국내 고교 졸업(예정)자로서 소속(출신) 고등학교장의 추천을 받은 자', 0, 0),
		(102, 'qual_1102', '국내 고교 졸업(예정)자로서 소속(출신) 고등학교장의 추천을 받은 자, 고교별 추천인원 1명', 0, 0),
		(103, 'qual_1103', '국내 고교 졸업(예정)자로서 졸업(예정) 고등학교의 학교장 추천을 받은 자, 학교별 추천인원 제한 없음', 0, 0),
		(104, 'qual_1104', '국내 고교 졸업(예정)자로서 출신 고등학교 추천을 받은 자, 고교별 재적 학생 수의 4%까지 추천, 계열별 지원인원 제한 없음, 특성화고 및 일반고 등 전문계 과정은 지원 불가', 0, 0),
		(105, 'qual_1105', '국내 고교 졸업(예정)자로서 학교생활기록부 교과 성적이 4개 학기 이상 있으며 학교생활을 충실히 한 자', 0, 0),
		(106, 'qual_1106', '국내 고교 졸업(예정)자로서 학교생활기록부 성적이 3개 학기 이상 있는 자', 0, 0),
		(107, 'qual_1107', '국내 고교 졸업(예정)자로서 학교장의 추천을 받는 자, 고교별 추천인원 제한 없음', 0, 0),
		(108, 'qual_1108', '국내 고교 졸업예정자 중 4학기 이상 성적 취득한 자로서 출신 고등학교장의 추천을 받은 자, 고교별 추천 인원 10명 이내', 0, 0),
		(109, 'qual_1109', '국내 고교 졸업예정자로서 학교장 추천한 자, 문화인재/ 글로벌인재/ 리더십인재/ 과학인재, 고교별 인문 최대 2명, 자연 최대 3명, 예술-체육 최대 1명 추천 가능', 0, 0),
		(110, 'qual_1110', '국내 고교 졸업예정자이며 소속 고등학교장의 추천을 받은 자로서, 고교 전 교육과정을 국내 고교에서 이수하고 추천인원은 3학년 재학인원의 5%이내', 0, 0),
		(111, 'qual_1111', '국내 고교에서 3학기 이상 이수한 졸업(예정)자로서 소속(졸업) 고등학교장의 추천을 받은 자, 고교별 7명 이내', 0, 0),
		(112, 'qual_1112', '국내 정규 고교 졸업(예정)자', 0, 0),
		(113, 'qual_1113', '국내 정규 고교 졸업(예정)자 / 검정고시 지원불가 / 군인사법 저촉없음', 0, 0),
		(114, 'qual_1114', '국내 정규 고교 졸업(예정)자 또는 동등학력 소지자', 0, 0),
		(115, 'qual_1115', '국내 정규 고교 졸업(예정)자 또는 동등학력 인정된 자', 0, 0),
		(116, 'qual_1116', '국내 정규 고교 졸업(예정)자 또는 동등학력 인정된 자로 첨단 분야에 관심과 역량을 갖춘 자', 0, 0),
		(117, 'qual_1117', '국내 정규 고교 졸업(예정)자, 수영과 교과 70단위 이상(전문교과 포함) 이수한 자만 지원가능', 0, 0),
		(118, 'qual_1118', '국내 정규 고교 졸업(예정)자로 소속 고등학교장의 추천을 받은 자, 고교별 추천인원은 최대 10명(계열구분 없음)', 0, 0),
		(119, 'qual_1119', '국내 정규 고교 졸업(예정)자로서 3개 학기 이상 성적을 취득하고, 소프트웨어 및 정보보안 분야에 관심과 활동이 있는 자', 0, 0),
		(120, 'qual_1120', '국내 정규 고교 졸업(예정)자로서 3개 학기 이상 성적을 취득한 자', 0, 0),
		(121, 'qual_1121', '국내 정규 고교 졸업(예정)자로서 3학기 이상 학생부 성적이 기재된 자', 0, 0),
		(122, 'qual_1122', '국내 정규 고교 졸업(예정)자로서 4개 학기 이상 학교생활기록부 성적이 있는 자', 0, 0),
		(123, 'qual_1123', '국내 정규 고교 졸업(예정)자로서 4개 학기 이상 학교생활기록부 성적이 있는 자 중 학교장 추천을 받은 자', 0, 0),
		(124, 'qual_1124', '국내 정규 고교 졸업(예정)자로서 5학기 이상 학생부 성적이 기재된 자로서 출신 고등학교장의 추천을 받은 자', 0, 0),
		(125, 'qual_1125', '국내 정규 고교 졸업(예정)자로서 군인사법에 저촉되지 아니한 자로서 임관일 기준 만 20세 이상 27세  이하인자', 0, 0),
		(126, 'qual_1126', '국내 정규 고교 졸업(예정)자로서 소속 고등학교장의 추천을 받은 자, 고교별 추천인원은 8명 제한', 0, 0),
		(127, 'qual_1127', '국내 정규 고교 졸업(예정)자로서 창의적 사고와 능동적 실천능력을 갖추고 전공 분야에 대한 열정과 우수한 재능을 가진 자', 0, 0),
		(128, 'qual_1128', '국내외 고교 졸업(예정) 또는 동등학력 이상 인정된 자', 0, 0),
		(129, 'qual_1129', '국내외 고교 졸업(예정) 또는 동등학력 이상 인정된 자로서 대학 인정 대회 3위 이내 입상자', 0, 0),
		(130, 'qual_1130', '국내외 고교 졸업(예정)자', 0, 0),
		(131, 'qual_1131', '국내외 고교 졸업(예정)자 또는 검정고시', 0, 0),
		(132, 'qual_1132', '국내외 고교 졸업(예정)자 또는 동등학력 이상 인정된 자', 0, 0),
		(133, 'qual_1133', '국내외 고교 졸업(예정)자 또는 동등학력 이상 인정된 자로서 해당 고교에서 추천받는 자', 0, 0),
		(134, 'qual_1134', '소속 고등학교장의 추천을 받은 자로서, 국내 고교 2017년 1월 이후 졸업(예정)자로서 3개 학기 이상 학교생활기록부 성적이 있는 자, 고교별 추천인원 20명 이내', 0, 0),
		(135, 'qual_1135', '인문계 고교 졸업(예정)자, 특목고 졸업(예정)자(예-체능특목고 제외), 인문계 고교 교육과정 이수(예정)자', 0, 0),
		(136, 'qual_1136', '입학 시부터 졸업 시까지 강원지역 고교 졸업(예정)자로서 학교생활기록부가 있는 자', 0, 0),
		(137, 'qual_1137', '입학에서 졸업(예정)까지 고교 전 교육과정을 대구-경북지역에서 고교에서 이수한 자', 0, 0),
		(138, 'qual_1138', '입학에서 졸업까지 고등학교 전 과정을 대구-경북지역 고교에서 이수한 자', 0, 0),
		(139, 'qual_1139', '충북-충남-대전-세종 소재 고교 입학부터 졸업(예정)한 자', 0, 0),
		(140, 'qual_1140', '충북-충남-대전-세종 소재 고교 입학하여 전 교육과정을 이수한 졸업(예정)자', 0, 0),
		(141, 'qual_1141', '충북-충남-대전-세종 소재 고교 전 교육과정을 이수한 졸업(예정)자', 0, 0)
		;`)
	utils.HandleErr(err, "create table1")
	_, err = stmt.Exec()
	utils.HandleErr(err, "create table11 exec")
}