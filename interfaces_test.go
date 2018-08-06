package sqlx_test

import (
	"database/sql"
	"github.com/srikrsna/sqlx"
	"testing"
)

func TestCompatability(t *testing.T) {
	testDB(&sql.DB{})
	testScanner(&sql.Rows{})
	testScanner(&sql.Row{})
}

func testDB(_ sqlx.DB) {}

func testScanner(_ sqlx.Scanner) {}
