//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what sendAll — call send for each job in order
package main

// sendAll passes every job of js to send, in order.
func sendAll(js []packJob, send func(packJob)) {
	for _, j := range js {
		send(j)
	}
}
