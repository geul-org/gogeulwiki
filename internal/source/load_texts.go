//ff:func feature=source type=loader control=iteration dimension=1
//ff:what loadTexts — the language-tagged strings a query returns as (language, text) rows
package source

import "database/sql"

// loadTexts runs query (which selects language, text and value_json for $1 =
// id) and returns the rows in order, each text exact
// (exactValue).
func loadTexts(db *sql.DB, query, id string) ([]Text, error) {
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Text
	for rows.Next() {
		var t Text
		var exact string
		if err := rows.Scan(&t.Lang, &t.Value, &exact); err != nil {
			return nil, err
		}
		t.Value = exactValue(t.Value, exact)
		out = append(out, t)
	}
	return out, rows.Err()
}
