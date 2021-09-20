package app

import (
	"cnest_consult/univ/utils"
	"database/sql"
)

func db_apply_form(db *sql.DB) {
	
	// Drop Table
	stmt, err := db.Prepare(`DROP TABLE IF EXISTS apply_form;`)
	utils.HandleErr(err, "Drop apply_form table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Drop apply_form table")

	// Create Table
	stmt, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS apply_form (
			id SERIAL PRIMARY KEY,
			formId VARCHAR(16) NOT NULL ,
			formId VARCHAR(32) NOT NULL,
			sort integer,
			flag integer,
			inUser VARCHAR(24),
			inDate TIMESTAMP,
			reUser VARCHAR(24),
			reDate TIMESTAMP );
		`)

	utils.HandleErr(err, "Create apply_form table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Create apply_form table")

	//  Insert Data
	stmt, err = db.Prepare(`
		INSERT INTO apply_form (id, formId, formId, sort, flag) VALUES
		(1, 'form_101', '학생부교과', 0, 0),
		(2, 'form_102', '학생부종합', 0, 0),
		(3, 'form_103', '논술', 0, 0),
		(4, 'form_104', '실기', 0, 0);
	`)
	
	utils.HandleErr(err, "Insert Data into apply_form table")
	_, err = stmt.Exec()
	utils.HandleErr(err, "Exec for Insert Data into apply_form table")
}
