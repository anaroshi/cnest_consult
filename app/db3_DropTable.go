package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db3_DropTable(db *sql.DB) {
			
	statement, err := db.Prepare(`DROP TABLE IF EXISTS univ_info;`)

	utils.HandleErr(err, "drop table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table11 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS recruit_method;`)
	utils.HandleErr(err, "drop table2")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table12 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS recruit_sununglimit;`)
	utils.HandleErr(err, "drop table3")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table13 exec")
}
