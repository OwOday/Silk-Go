/*
* Copyright 2022, The Cozo Project Authors.
*
* This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
* If a copy of the MPL was not distributed with this file,
* You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package silk

import (
	"testing"

	"github.com/cozodb/cozo-lib-go"
)

func TestDb(t *testing.T) {

	db, err := cozo.New("mem", "", nil)
	if err != nil {
		t.Error(err)
	}
	_ = db.Restore("../db/test.db")
	_, _ = db.Run("?[a,b,c] <- [[1,2,3]] :create s{a, b, c}", nil)
	{
		res, _ := db.Run("?[a,b,c] := *s[a,b,c]", nil)
		if len(res.Rows) != 1 || len(res.Rows[0]) != 3 {
			t.Error("Bad number of rows")
		}
	}
	_ = db.Backup("../db/test.db")
	db.Close()
	// {
	// 	_, err := db.Run("?[x] <- [[1,2,3]]", nil)
	// 	if err == nil {
	// 		t.Error("expect an error")
	// 	}
	// }

	// {
	// 	_ = db.Backup("test.db")

	// 	db2, _ := cozo.New("mem", "", nil)
	// 	_ = db2.Restore("test.db")

	// 	res, err := db2.Run("?[a,b,c] := *s[a,b,c]", nil)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	if len(res.Rows) != 1 || len(res.Rows[0]) != 3 {
	// 		t.Error("Bad number of rows")
	// 	}
	// }

	// {
	// 	data, err := db.ExportRelations([]string{"s"})
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	db3, _ := cozo.New("mem", "", nil)
	// 	_, _ = db3.Run(":create s {a, b, c}", nil)

	// 	res, err := db3.Run("?[a,b,c] := *s[a,b,c]", nil)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	if len(res.Rows) != 0 {
	// 		t.Error("Bad number of rows")
	// 	}
	// 	_ = db3.ImportRelations(data)

	// 	res, err = db3.Run("?[a,b,c] := *s[a,b,c]", nil)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}

	// 	if len(res.Rows) != 1 || len(res.Rows[0]) != 3 {
	// 		t.Error("Bad number of rows")
	// 	}

	// 	db4, _ := cozo.New("mem", "", nil)
	// 	_, _ = db4.Run(":create s {a, b, c}", nil)

	// 	res, err = db4.Run("?[a,b,c] := *s[a,b,c]", nil)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	if len(res.Rows) != 0 {
	// 		t.Error("Bad number of rows")
	// 	}
	// 	_ = db4.ImportRelationsFromBackup("test.db", []string{"s"})

	// 	res, err = db4.Run("?[a,b,c] := *s[a,b,c]", nil)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}

	// 	if len(res.Rows) != 1 || len(res.Rows[0]) != 3 {
	// 		t.Error("Bad number of rows")
	// 	}

	// }

	// db.Close()
}
