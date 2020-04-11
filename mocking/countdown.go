package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	SleepTime int
}

type DefaultSleeper struct{}

type CheckSequenceSleeper struct {
	Seq []string
}

type ConfigurableSleeper struct {
	sleepTime time.Duration
	sleepFunc func(time.Duration)
}

func (m *SpySleeper) Sleep() {
	(*m).SleepTime++
}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (c *CheckSequenceSleeper) Sleep() {
	c.Seq = append(c.Seq, sleep)
}

func (c *CheckSequenceSleeper) Write(w []byte) (written int, err error) {
	c.Seq = append(c.Seq, write)
	return
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleepFunc(c.sleepTime)
}

func CountDown(w io.Writer, t Sleeper) {
	for i := 3; i > 0; i-- {
		t.Sleep()
		fmt.Fprintf(w, "%d\n", i)
	}
	t.Sleep()
	fmt.Fprintf(w, "Go!")
}

var sleep string
var write string

func main() {
	c := &SpySleeper{}
	CountDown(os.Stdout, c)
}
