//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what pageKind — send every entity of one kind the dump holds after a number and up to a bound, page by page
package main

import (
	"database/sql"

	"github.com/geul-org/wikigeul/internal/source"
)

// pageKind walks the dump's entities of kind in (after, to] in pages of
// 10,000 (source.EntityNumbers) and calls send for each, in order.
func pageKind(db *sql.DB, c config, kind byte, after, to uint64, send func(packJob)) error {
	for {
		nums, err := source.EntityNumbers(db, source.Store{Schema: c.Schema, Dump: c.Dump}, kind, after, to, 10000)
		if err != nil || len(nums) == 0 {
			return err
		}
		for _, n := range nums {
			send(packJob{Kind: kind, Num: n})
		}
		after = nums[len(nums)-1]
	}
}
