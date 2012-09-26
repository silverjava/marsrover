package marsrover

import (
    "strings"
    "strconv"
    "fmt"
    "errors"
)

const (
    EAST = iota
    SOUTH
    WEST
    NORTH
)

func Run(lines []string) {
    // plateau := CreatePlateauFromInput(lines[0])

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

func stringToOrientation(s string) (int, error) {
    facing := -1
    switch s[0] {
    case 'E': facing = EAST
    case 'S': facing = SOUTH
    case 'W': facing = WEST
    case 'N': facing = NORTH
    }

    if facing == -1 {
        return -1, errors.New("Invalid given parameter")
    }
    return facing, nil
}

func CreateRoverFromInput(line string) *Rover {
    strs := strings.Split(line, " ")

    x, _ := strconv.Atoi(strs[0])
    y, _ := strconv.Atoi(strs[1])
    facing, _ := stringToOrientation(strs[2])

    return &Rover{x:x, y:y, facing:facing}
}
