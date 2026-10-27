//ff:func feature=source type=loader control=sequence
//ff:what entityHead — whether the head query finds the entity, with its datatype (properties only) and its revision ID and time
package source

import "database/sql"

// entityHead runs query (count, datatype, lastrevid and modified as POSIX
// seconds for $1 = id) and reports whether the entity was loaded (a deleted
// or never created ID has no row), returning it with Datatype (the Wikibase
// datatype of a property — url, external-id …; empty for an item),
// LastRevID and Modified set (0 when the dump has none).
func entityHead(db *sql.DB, query, id string) (Entity, bool, error) {
	var n int
	e := Entity{ID: id}
	if err := db.QueryRow(query, id).Scan(&n, &e.Datatype, &e.LastRevID, &e.Modified); err != nil {
		return Entity{}, false, err
	}
	return e, n > 0, nil
}
