package marsrover

import (
    "strings"
    "strconv"
    "fmt"
)

const (
    EAST = iota
    SOUTH
    WEST
    NORTH
)

func Run(lines []string) {
    for i := 1; i < len(lines); {
        rover := CreateRoverFromInput(lines[i])
        ControllRover(lines[i + 1], rover)

        fmt.Printf("%v\n", rover)
        i = i + 2
    }
}

func CreatePlateauFromInput(line string) *Plateau {
    strs := strings.Split(line, " ")

    width, _ := strconv.Atoi(strs[0])
    height, _ := strconv.Atoi(strs[1])
    return &Plateau{width:width, height:height}
}

func ControllRover(cmd string, r *Rover) {
    for _, c := range cmd {
        switch c {
        case 'L': r.TurnLeft()
        case 'R': r.TurnRight()
        case 'M': r.MoveForward()
        }
    }
}

func stringToOrientation(s string) (facing int, err error) {
    switch s[0] {
    case 'E': facing = EAST
    case 'S': facing = SOUTH
    case 'W': facing = WEST
    case 'N': facing = NORTH
    default:
        facing = -1
        err = fmt.Errorf("Invalid given parameter.")
    }
    return
}

func CreateRoverFromInput(line string) *Rover {
    strs := strings.Split(line, " ")

    x, _ := strconv.Atoi(strs[0])
    y, _ := strconv.Atoi(strs[1])
    facing, _ := stringToOrientation(strs[2])

    return &Rover{x:x, y:y, facing:facing}
}
