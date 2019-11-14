package demo

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Demo struct {
	steps []step
}

type step struct {
	text, command []string
}

func New(description ...string) *Demo {
	for _, d := range description {
		color.Cyan.Printf("# %s\n", d)
	}
	return &Demo{}
}

func S(s ...string) []string {
	return s
}

func (d *Demo) Step(text []string, command []string) {
	d.steps = append(d.steps, step{text, command})
}

func (d *Demo) Run(ctx *cli.Context) {
	run := func() {
		for i, step := range d.steps {
			step.run(i+1, len(d.steps), ctx.GlobalBool("auto"))
		}
	}
	if ctx.GlobalBool("continuously") {
		for {
			run()
		}
	} else {
		run()
	}
}

func Ensure(commands ...[]string) {
	for _, c := range commands {
		command := strings.Join(c, " ")
		cmd := exec.Command("bash", "-c", command)
		cmd.Stderr = nil
		cmd.Stdout = nil
		_ = cmd.Run()
	}
}

func (s *step) run(current, max int, auto bool) {
	waitOrSleep(auto)
	s.echo(current, max)
	s.execute(auto)
}

func (s *step) echo(current, max int) {
	prepared := []string{""}
	for i, x := range s.text {
		if i == len(s.text)-1 {
			prepared = append(
				prepared,
				color.White.Darken().Sprintf(
					"# %s [%d/%d]:\n",
					x, current, max,
				),
			)
		} else {
			m := color.White.Darken().Sprintf("# %s", x)
			prepared = append(prepared, m)
		}
	}
	print(prepared...)
}

func (s *step) execute(auto bool) {
	joinedCommand := strings.Join(s.command, " ")
	cmd := exec.Command("bash", "-c", joinedCommand)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmdString := color.Green.Sprintf("> %s", strings.Join(s.command, " \\\n    "))
	print(cmdString)
	waitOrSleep(auto)
	_ = cmd.Run()
}

func print(msg ...string) {
	for _, m := range msg {
		for _, c := range m {
			time.Sleep(time.Duration(rand.Intn(40)) * time.Millisecond)
			fmt.Printf("%c", c)
		}
		println()
	}
}

func waitOrSleep(auto bool) {
	if auto {
		time.Sleep(3 * time.Second)
	} else {
		if _, err := bufio.NewReader(os.Stdin).ReadBytes('\n'); err != nil {
			logrus.Fatalf("Unable to read input: %v", err)
		}
	}
}
