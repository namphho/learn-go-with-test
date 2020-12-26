package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countDownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
}

func main() {
	//sleeper := &DefaultSleeper{}
	//CountDown(os.Stdout, sleeper)
	configurableSleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, configurableSleeper)
}

func CountDown(w io.Writer, sleep Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleep.Sleep()
		fmt.Fprintln(w, i)
	}
	sleep.Sleep()
	_, _ = fmt.Fprint(w, finalWord)
}
