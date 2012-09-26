package marsrover

import (
    "fmt"
    "testing"
)

func TestCreatePlateau(t *testing.T) {
    width, height := 5, 5

    input := fmt.Sprintf("%d %d", width, height)

    p := CreatePlateauFromInput(input)
    if p == nil || p.width != width || p.height != height {
        t.Errorf("Can't create plateau with input %s", input)
    }
}

func TestControllRoverTurnLeft(t *testing.T) {
    rover := &Rover{x:3, y:3, facing:EAST}

    ControllRover("L", rover, plateau)
    if rover.facing != NORTH {
        t.Errorf("Wrong direction after turning left.")
    }

    rover.facing = EAST
    ControllRover("LL", rover, plateau)
    if rover.facing != WEST {
        t.Errorf("Wrong direction after turning left.")
    }
}

func TestControlRoverTurnRight(t *testing.T) {
    rover := &Rover{x:3, y:3, facing:EAST}

    ControllRover("R", rover, plateau)
    if rover.facing != SOUTH {
        t.Errorf("Wrong direction after turning left.")
    }

    rover.facing = EAST
    ControllRover("RR", rover, plateau)
    if rover.facing != WEST {
        t.Errorf("Wrong direction after turning left.")
    }
}

func TestControllRoverMove(t *testing.T) {
    x, y := 3, 3
    rover := &Rover{x:x, y:y, facing:EAST}

    ControllRover("M", rover, plateau)
    if rover.x != x + 1 && rover.y == y {
        t.Errorf("Wrong moving.")
    }

    rover = &Rover{x:x, y:y, facing:SOUTH}
    ControllRover("M", rover, plateau)
    if rover.y != y - 1 && rover.x == x {
        t.Errorf("Wrong moving.")
    }
}

func TestControllRoverTurnAndMove(t *testing.T) {
    x, y := 3, 3
    rover := &Rover{x:x, y:y, facing:EAST}

    ControllRover("LM", rover, plateau)
    if rover.facing != NORTH && rover.y != y + 1 && rover.x == x {
        t.Errorf("Wrong move and turn.")
    }
}

func TestStringToOrientation(t *testing.T) {
    testData := make(map[string]int)
    testData["E"] = EAST
    testData["S"] = SOUTH
    testData["W"] = WEST
    testData["N"] = NORTH

    for k, v := range testData {
        facing, _ := stringToOrientation(k)
        if facing != v {
            t.Errorf("Failed to convert string to orientation.")
        }
    }
}

func TestCreateRoverFromInput(t *testing.T) {
    rover := CreateRoverFromInput("3 3 E")

    if rover.x != 3 || rover.y != 3 || rover.facing != EAST {
        t.Errorf("Failed to create rover from input.")
    }
}
