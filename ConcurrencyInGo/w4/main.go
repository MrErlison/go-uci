package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	number int
	leftCS,
	rightCS *ChopS
}

func (p Philo) eat(ch chan int) {
	const eatTimes int = 3

	for i := 0; i < eatTimes; i++ {
		<-ch

		fmt.Printf("starting to eat %d\n", p.number)

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("finishing eating %d\n", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		ch <- i
	}
	wg.Done()
}

func host(ch chan int) {
	ch <- 1
	ch <- 2
	<-ch
}

func main() {
	ch := make(chan int, 2)

	const philosophers int = 5
	CSticks := make([]*ChopS, philosophers)
	for i := 0; i < philosophers; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, philosophers)
	for i := 0; i < philosophers; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%philosophers]}
	}

	go host(ch)

	wg.Add(philosophers)
	for i := 0; i < philosophers; i++ {
		go philos[i].eat(ch)
	}
	wg.Wait()
}
