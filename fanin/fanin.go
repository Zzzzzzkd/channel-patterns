package fallin

import "sync"

func FallIn(chs ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chs))
		for _, ch := range chs {
			go func(ch <-chan interface{}) {
				for v := range ch {
					out <- v
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		close(out)
	}()
	return out
}
