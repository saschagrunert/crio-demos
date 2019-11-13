package demo

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
)

type Demo struct {
	steps []step
}

type step struct {
	text, command []string
}

func New(title string) *Demo {
	color.Cyan.Printf("# %s\n", title)
	return &Demo{}
}

func X(x ...string) []string {
	return x
}

func (d *Demo) Step(text []string, command []string) {
	d.steps = append(d.steps, step{text, command})
}

func (d *Demo) Run() {
	for i, step := range d.steps {
		step.run(i+1, len(d.steps))
	}
}

func (s *step) run(current, max int) {
	s.echo(current, max)
	s.execute()
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

func (s *step) execute() {
	joinedCommand := strings.Join(s.command, " ")
	args := strings.Fields(joinedCommand)
	cmd := exec.Command(s.command[0])
	if len(args) > 1 {
		cmd = exec.Command(args[0], args[1:]...)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	color.Green.Printf("> %s", joinedCommand)
	waitEnter()

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
