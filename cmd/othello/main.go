// Command othello is a console-based program to play the othello game. 
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/z-rui/game/board"
	"github.com/z-rui/game/othello"
	"os"
	"runtime/pprof"
)

var (
	cpuLevel      = flag.Uint("L", 5, "CPU Level: 1(weakest)...9(strongest)")
	demoMode      = flag.Bool("a", false, "Two Cpus play with each other")
	verboseSearch = flag.Bool("v", false, "Show Cpu decision details")
	boxChars      = flag.Bool("U", false, "Use box-drawing characters")
	cpuProfile    = flag.String("p", "", "Write cpu profile to file")
)

var (
	stdin = bufio.NewReader(os.Stdin)
)

type Player interface {
	Next(s *othello.State) *othello.State
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
		if askPlaying() == othello.X {
			p[0], p[1] = p[1], p[0]
		}
	}

	s := othello.NewState()
	i := 0
	for {
		if *boxChars {
			board.Print(os.Stdout, s, board.UnicodeBox)
		} else {
			board.Print(os.Stdout, s, board.AsciiBox)
		}
		o, x := s.Count()
		fmt.Printf("O: %d, X: %d\n", o, x)
		t := p[i].Next(s)
		if t == nil {
			break
		}
		s = t
		who := p[i].Name()
		if !s.LastMove.Valid() {
			fmt.Println(who, "passes")
		} else {
			fmt.Println(who, "went", s.LastMove)
		}
		i ^= 1
	}

	fmt.Print("Game over.  ")
	o, x := s.Count()
	switch {
	case o > x:
		fmt.Println(p[0].Name(), "won")
	case o < x:
		fmt.Println(p[1].Name(), "won")
	default:
		fmt.Println("It was a draw")
	}
}
