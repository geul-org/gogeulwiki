//ff:func feature=port type=helper control=sequence
//ff:what q42Refs — test fixture: references of Q42 P19 as the 20261019 dump has them (Britannica, Guardian, Telegraph, GND) plus URL-only, imported-from, P248 somevalue and the Guardian again in another snaks-order
package port

import "github.com/geul-org/wikigeul/internal/source"

// q42Refs returns references shaped like Q42's (wikidata.reference_snak,
// dump 1, 2026-10-27): ① Britannica (P248 + P1417), ② Guardian (URL,
// publication, retrieved, published in, author, title, language), ③ Telegraph
// (two P2093 authors), ④ GND (P248 + P227 + P407 + P813), ⑤ URL only, ⑥
// imported from (P143) with its import URL, ⑦ P248 somevalue with a page and
// ⑧ the Guardian description of ② in another snaks-order with another
// retrieval date.
func q42Refs() []source.Reference {
	item := func(p, q string) source.Qualifier {
		return source.Qualifier{Property: p, Type: "wikibase-entityid", Datatype: "wikibase-item", Value: q, SnakType: "value"}
	}
	str := func(p, dt, v string) source.Qualifier {
		return source.Qualifier{Property: p, Type: "string", Datatype: dt, Value: v, SnakType: "value"}
	}
	day := func(p, d string) source.Qualifier {
		return source.Qualifier{Property: p, Type: "time", Datatype: "time", Value: `{"time": "+` + d + `T00:00:00Z", "after": 0, "before": 0, "timezone": 0, "precision": 11, "calendarmodel": "http://www.wikidata.org/entity/Q1985727"}`, SnakType: "value"}
	}
	text := func(p, s string) source.Qualifier {
		return source.Qualifier{Property: p, Type: "monolingualtext", Datatype: "monolingualtext", Value: `{"text": "` + s + `", "language": "en"}`, SnakType: "value"}
	}
	guardian := "http://www.theguardian.com/news/2001/may/15/guardianobituaries.books"
	return []source.Reference{
		{Hash: "1", Snaks: []source.Qualifier{item("P248", "Q5375741"), str("P1417", "external-id", "biography/Douglas-Adams")}},
		{Hash: "2", Snaks: []source.Qualifier{str("P854", "url", guardian), day("P577", "2001-05-15"), day("P813", "2013-12-07"), item("P1433", "Q5614018"), item("P50", "Q18145749"), text("P1476", "Obituary: Douglas Adams"), item("P407", "Q1860")}},
		{Hash: "3", Snaks: []source.Qualifier{text("P1476", "Hitch Hiker's Guide author Douglas Adams dies aged 49"), day("P577", "2001-05-13"), item("P407", "Q1860"), str("P854", "url", "http://www.telegraph.co.uk/news/uknews/1330072/"), day("P813", "2015-01-03"), str("P2093", "string", "Andrew Alderson"), str("P2093", "string", "Daniel Foggo"), item("P1433", "Q192621")}},
		{Hash: "4", Snaks: []source.Qualifier{item("P248", "Q36578"), str("P227", "external-id", "119033364"), item("P407", "Q188"), day("P813", "2022-10-09")}},
		{Hash: "5", Snaks: []source.Qualifier{str("P854", "url", "http://www.eskimo.com/~rkj/weekly/aa051701a.htm")}},
		{Hash: "6", Snaks: []source.Qualifier{item("P143", "Q328"), str("P4656", "url", "https://en.wikipedia.org/w/index.php?oldid=1")}},
		{Hash: "7", Snaks: []source.Qualifier{{Property: "P248", Datatype: "wikibase-item", SnakType: "somevalue"}, str("P304", "string", "12")}},
		{Hash: "8", Snaks: []source.Qualifier{item("P407", "Q1860"), text("P1476", "Obituary: Douglas Adams"), item("P50", "Q18145749"), item("P1433", "Q5614018"), day("P813", "2014-01-03"), day("P577", "2001-05-15"), str("P854", "url", guardian)}},
	}
}
