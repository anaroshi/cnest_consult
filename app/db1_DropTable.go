package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db1_DropTable(db *sql.DB) {
			
	statement, err := db.Prepare(`DROP TABLE IF EXISTS apply_subject;`)

	utils.HandleErr(err, "drop table1")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table11 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS apply_line;`)
	utils.HandleErr(err, "drop table2")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table12 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS apply_form;`)
	utils.HandleErr(err, "drop table3")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table13 exec")

	statement, err = db.Prepare(`DROP TABLE IF EXISTS apply_dept;`)
	utils.HandleErr(err, "drop table4")
	_, err = statement.Exec()
	utils.HandleErr(err, "drop table14 exec")
}
