package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, n)
	var wg sync.WaitGroup
	sem := make(chan struct{}, pool)
	wg.Add(int(n))
	var i int64
	for i = 0; i < n; i++ {
		sem <- struct{}{}
		go func(j int64) {
			res[j] = getOne(j)
			<-sem
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
