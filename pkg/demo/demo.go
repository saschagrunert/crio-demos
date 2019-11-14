package demo

import (
	"bufio"
	"fmt"
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
	for i, step := range d.steps {
		step.run(i+1, len(d.steps), ctx.GlobalBool("auto"))
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

	color.Green.Printf("> %s", strings.Join(s.command, " \\\n    "))
	if auto {
		time.Sleep(3 * time.Second)
	} else {
		waitEnter()
	}

	if err := cmd.Run(); err != nil {
		logrus.Fatalf("Command execution failed: %v", err)
	}
}

func print(msg ...string) {
	fmt.Print(strings.Join(msg, "\n"))
}

func waitEnter() {
	if _, err := bufio.NewReader(os.Stdin).ReadBytes('\n'); err != nil {
		logrus.Fatalf("Unable to read input: %v", err)
	}
}
