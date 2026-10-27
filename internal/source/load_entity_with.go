//ff:func feature=source type=loader control=sequence
//ff:what loadEntityWith — an entity's labels, descriptions, aliases, claims and their qualifiers by the given queries; false when the ID was never loaded
package source

import "database/sql"

// loadEntityWith reads the entity id by q: existence and datatype, then labels, descriptions and
// aliases in every language, then claims with their qualifiers (attachQualifiers), then sitelinks. It is false, with no error, when the entity does not
// exist.
func loadEntityWith(db *sql.DB, q entityQueries, id string) (Entity, bool, error) {
	datatype, ok, err := entityHead(db, q.Exists, id)
	if err != nil || !ok {
		return Entity{}, false, err
	}
	e := Entity{ID: id, Datatype: datatype}
	if e.Labels, err = loadTexts(db, q.Labels, id); err != nil {
		return Entity{}, false, err
	}
	if e.Descriptions, err = loadTexts(db, q.Descriptions, id); err != nil {
		return Entity{}, false, err
	}
	if e.Aliases, err = loadTexts(db, q.Aliases, id); err != nil {
		return Entity{}, false, err
	}
	if e.Claims, err = loadClaims(db, q.Claims, id); err != nil {
		return Entity{}, false, err
	}
	qs, err := loadQualifiers(db, q.Qualifiers, id)
	if err != nil {
		return Entity{}, false, err
	}
	attachQualifiers(e.Claims, qs)
	if e.Sitelinks, err = loadSitelinks(db, q.Sitelinks, id); err != nil {
		return Entity{}, false, err
	}
	return e, true, nil
}
