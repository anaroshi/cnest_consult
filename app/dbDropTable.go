package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func dbDropTable(db *sql.DB) {
			
	statement, err := db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_fst;`)

	utils.HandleErr(err, "drop table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table11 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_snd;`)
	utils.HandleErr(err, "drop table2")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table12 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_2021;`)
	utils.HandleErr(err, "drop table3")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table13 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_2020;`)
	utils.HandleErr(err, "drop table4")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table14 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS univ_susi_info_2019;`)
	utils.HandleErr(err, "drop table5")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table15 exec")

}