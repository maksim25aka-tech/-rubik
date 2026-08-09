// rubik.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"
)

type Rubik4x4 struct {
	faces [6][4][4]string
}

func NewRubik4x4() *Rubik4x4 {
	r := &Rubik4x4{}
	colors := [6]string{"W", "Y", "R", "O", "B", "G"}
	for f := 0; f < 6; f++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				r.faces[f][i][j] = colors[f]
			}
		}
	}
	return r
}

func (r *Rubik4x4) rotateFaceCW(face *[4][4]string) {
	// Поворот 4x4 по часовой
	newFace := [4][4]string{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newFace[i][j] = face[3-j][i]
		}
	}
	*face = newFace
}

func (r *Rubik4x4) rotateFaceCCW(face *[4][4]string) {
	newFace := [4][4]string{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newFace[i][j] = face[j][3-i]
		}
	}
	*face = newFace
}

func (r *Rubik4x4) applyMove(move string) {
	// Упрощённая реализация: только U, R, F
	switch move {
	case "U":
		r.rotateFaceCW(&r.faces[0])
	case "U'":
		r.rotateFaceCCW(&r.faces[0])
	case "R":
		r.rotateFaceCW(&r.faces[2])
	case "R'":
		r.rotateFaceCCW(&r.faces[2])
	case "F":
		r.rotateFaceCW(&r.faces[4])
	case "F'":
		r.rotateFaceCCW(&r.faces[4])
	default:
		fmt.Println("Неизвестный ход")
	}
}

func (r *Rubik4x4) scramble(moves int) []string {
	possible := []string{"U", "U'", "R", "R'", "F", "F'"}
	seq := make([]string, moves)
	for i := 0; i < moves; i++ {
		seq[i] = possible[rand.Intn(len(possible))]
		r.applyMove(seq[i])
	}
	return seq
}

func (r *Rubik4x4) display() {
	fmt.Println("\n--- Кубик Рубика 4x4 ---")
	names := []string{"U", "D", "R", "L", "F", "B"}
	for f := 0; f < 6; f++ {
		fmt.Printf("%s:\n", names[f])
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				fmt.Print(r.faces[f][i][j], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cube := NewRubik4x4()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Добро пожаловать в симулятор кубика Рубика 4x4!")
	fmt.Println("Команды: U, U', R, R', F, F', scramble, reset, quit, help")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		cmd := strings.TrimSpace(scanner.Text())
		switch cmd {
		case "quit":
			return
		case "help":
			fmt.Println("U, U', R, R', F, F', D, D', L, L', B, B' - ходы")
			fmt.Println("scramble - перемешать")
			fmt.Println("reset - собрать")
			fmt.Println("quit - выход")
		case "reset":
			cube = NewRubik4x4()
			fmt.Println("Куб собран.")
		case "scramble":
			seq := cube.scramble(50)
			fmt.Println("Перемешивание:", seq)
		default:
			if strings.HasPrefix(cmd, "U") || strings.HasPrefix(cmd, "R") || strings.HasPrefix(cmd, "F") {
				cube.applyMove(cmd)
			} else {
				fmt.Println("Неизвестная команда.")
			}
		}
		cube.display()
	}
}
