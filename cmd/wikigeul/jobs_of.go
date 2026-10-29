//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what jobsOf — jobs for entity numbers of one kind
package main

// jobsOf makes a job per number of kind.
func jobsOf(kind byte, nums []uint64) []packJob {
	out := make([]packJob, len(nums))
	for i, n := range nums {
		out[i] = packJob{Kind: kind, Num: n}
	}
	return out
}
