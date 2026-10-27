//ff:func feature=source type=loader control=iteration dimension=1
//ff:what loadClaims — the statements a query returns as (statement ID, property, type, value, rank) rows
package source

import "database/sql"

// loadClaims runs query (which selects statement ID, property, value type, datatype, value,
// value_json and rank for $1 = id) and returns the rows in order, each value
// exact (exactValue).
func loadClaims(db *sql.DB, query, id string) ([]Claim, error) {
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Claim
	for rows.Next() {
		var c Claim
		var exact string
		if err := rows.Scan(&c.ID, &c.Property, &c.Type, &c.Datatype, &c.Value, &exact, &c.Rank); err != nil {
			return nil, err
		}
		c.Value = exactValue(c.Value, exact)
		out = append(out, c)
	}
	return out, rows.Err()
}
