// Command tictactoe is a console-based program to play the Tic-Tac-Toe game.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/z-rui/game"
	"github.com/z-rui/game/board"
	"github.com/z-rui/game/tictactoe"
	"os"
	"runtime/pprof"
)

var (
	cpuLevel      = flag.Uint("L", 9, "CPU Level: 1(weakest)...9(strongest)")
	demoMode      = flag.Bool("a", false, "Two Cpus play with each other")
	verboseSearch = flag.Bool("v", false, "Show Cpu decision details")
	boxChars      = flag.Bool("U", false, "Use box-drawing characters")
	cpuProfile    = flag.String("p", "", "Write cpu profile to file")
)

var (
	stdin = bufio.NewReader(os.Stdin)
)

type Player interface {
	Next(s *tictactoe.State) *tictactoe.State
	Name() string
}

func main() {
	flag.Parse()
	if *cpuLevel < 1 {
		*cpuLevel = 1
	}
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var p [2]Player
	if *demoMode {
		p[0] = &CpuPlayer{"CPU 1", *cpuLevel}
		p[1] = &CpuPlayer{"CPU 2", *cpuLevel}
	} else {
		p[0] = &HumanPlayer{"You"}
		p[1] = &CpuPlayer{"CPU", *cpuLevel}
		if askPlaying() == tictactoe.X {
			p[0], p[1] = p[1], p[0]
		}
	}

	s := tictactoe.NewState()
	i := 0
	for {
		if *boxChars {
			board.Print(os.Stdout, s, board.UnicodeBox)
		} else {
			board.Print(os.Stdout, s, board.AsciiBox)
		}
		if t := p[i].Next(s); t == nil {
			break
		} else {
			s = t
		}
		who := p[i].Name()
		fmt.Println(who, "went", s.LastMove)
		i ^= 1
	}

	fmt.Print("Game over.  ")
	switch s.Eval() {
	case game.Won:
		fmt.Println(p[0].Name(), "won")
	case game.Lost:
		fmt.Println(p[1].Name(), "won")
	default:
		fmt.Println("It was a draw")
	}
}
