//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what addCounts — add one document's reason counts to a running total
package main

// addCounts adds each count of from to sum.
func addCounts(sum, from map[string]int) {
	for k, n := range from {
		sum[k] += n
	}
}
