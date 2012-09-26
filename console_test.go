package marsrover

import (
    "fmt"
    "testing"
)

func TestCreatePlateau(t *testing.T) {
    width, height := 5, 5

    input := fmt.Sprintf("%d %d", width, height)

    p := CreatePlateauFromInput(input)
    if p == nil {
        t.Errorf("Can't create plateau with input %s", input)
    }

    if p.width != width {
        t.Errorf("The width of plateau should be %d", width)
    }

    if p.height != height {
        t.Errorf("The height of plateau should be %d", height)
    }
}

func TestControllRoverTurnLeft(t *testing.T) {
    rover := &Rover{x:3, y:3, facing:EAST}

    ControllRover("L", rover)
    if rover.facing != NORTH {
        t.Errorf("Wrong direction after turning left.")
    }

    rover.facing = EAST
    ControllRover("LL", rover)
    if rover.facing != WEST {
        t.Errorf("Wrong direction after turning left.")
    }
}

func TestControlRoverTurnRight(t *testing.T) {
    rover := &Rover{x:3, y:3, facing:EAST}

    ControllRover("R", rover)
    if rover.facing != SOUTH {
        t.Errorf("Wrong direction after turning left.")
    }

    rover.facing = EAST
    ControllRover("RR", rover)
    if rover.facing != WEST {
        t.Errorf("Wrong direction after turning left.")
    }
}

func TestControllRoverMove(t *testing.T) {
    x, y := 3, 3
    rover := &Rover{x:x, y:y, facing:EAST}

    ControllRover("M", rover)
    if rover.x != x + 1 && rover.y == y {
        t.Errorf("Wrong moving.")
    }

    rover = &Rover{x:x, y:y, facing:SOUTH}
    ControllRover("M", rover)
    if rover.y != y - 1 && rover.x == x {
        t.Errorf("Wrong moving.")
    }
}

func TestControllRoverTurnAndMove(t *testing.T) {
    x, y := 3, 3
    rover := &Rover{x:x, y:y, facing:EAST}

    ControllRover("LM", rover)
    if rover.facing != NORTH && rover.y != y + 1 && rover.x == x {
        t.Errorf("Wrong move and turn.")
    }
}

func TestStringToOrientation(t *testing.T) {
    if facing, _ := stringToOrientation("E"); facing != EAST {
        t.Errorf("Failed to convert string to orientation.")
    }

    if facing, _ := stringToOrientation("S"); facing != SOUTH {
        t.Errorf("Failed to convert string to orientation.")
    }

    if facing, _ := stringToOrientation("W"); facing != WEST {
        t.Errorf("Failed to convert string to orientation.")
    }

    if facing, _ := stringToOrientation("N"); facing != NORTH {
        t.Errorf("Failed to convert string to orientation.")
    }

    if _, err := stringToOrientation("A"); err == nil {
        t.Errorf("Failed to convert string to orientation.")
    }
}

func TestCreateRoverFromInput(t *testing.T) {
    rover := CreateRoverFromInput("3 3 E")

    if rover.x != 3 || rover.y != 3 || rover.facing != EAST {
        t.Errorf("Failed to create rover from input.")
    }
}
