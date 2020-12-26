package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)



type SpySleeper struct {
	Call int
}

func (s *SpySleeper) Sleep(){
	s.Call++
}

type CountDownOperationSpy struct {
	Call []string
}
const sleep = "sleep"
const write = "write"

func (c *CountDownOperationSpy) Sleep(){
	c.Call = append(c.Call, "sleep")
}

func (c *CountDownOperationSpy) Write(p []byte) (n int, err error){
	c.Call = append(c.Call, "write")
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s * SpyTime) Sleep(time time.Duration){
	s.durationSlept = time
}


func TestCountDown(t *testing.T) {
	t.Run("test logic function", func(t *testing.T){
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		CountDown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got ,want)
		}
		if spySleeper.Call != 4{
			t.Errorf("sleep shoule be called 4 times")
		}
	})

	t.Run("test sequence write and sleep", func(t *testing.T){
		spySleepPrinter := &CountDownOperationSpy{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Call){
			t.Errorf("wanted %q got %q", want,spySleepPrinter.Call)
		}
	})

}

func TestConfigurableSleeper(t *testing.T){
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime{
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}