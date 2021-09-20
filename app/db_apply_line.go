package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_apply_line(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS apply_line;`)
	utils.HandleErr(err, "Drop apply_line table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop apply_line table")

	// Create Table
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS apply_line (
			id SERIAL PRIMARY KEY,
			lineId VARCHAR(16) NOT NULL ,
			lineName VARCHAR(32) NOT NULL,
			sort integer,
			flag integer,
			inUser VARCHAR(24),
			inDate TIMESTAMP,
			reUser VARCHAR(24),
			reDate TIMESTAMP );
		`)

	utils.HandleErr(err, "Create apply_line table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create apply_line table")

	//  Insert Data
	stmt, err = db.Prepare(`
		INSERT INTO apply_line (id, lineId, lineName, sort, flag) VALUES
		(1, 'line_101', '인문', 0, 0),
		(2, 'line_102', '자연', 0, 0),
		(3, 'line_103', '공통', 0, 0),
		(4, 'line_104', '예체능', 0, 0);
	`)
	
	utils.HandleErr(err, "Insert Data into apply_line table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Insert Data into apply_line table")
}
