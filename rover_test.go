package marsrover

import (
    "testing"
)

var x, y int = 3, 3

func TestRoverMovingWithFacingEast(t *testing.T) {
    facing := EAST
    rover := newRover(facing)

    rover.MoveForward()
    if rover.x != x + 1 {
        t.Fatalf("Moving is failed")
    }

    if rover.y != y {
        t.Fatalf("Moving is failed")
    }

    if rover.facing != rover.facing {
        t.Fatalf("Moving is failed")
    }
}

func TestRoverMovingWithFacingSouth(t *testing.T) {
    facing := SOUTH
    rover := newRover(facing)

    rover.MoveForward()
    if rover.y != y - 1 {
        t.Fatalf("Moving is failed")
    }

    if rover.x != x {
        t.Fatalf("Moving is failed")
    }

    if rover.facing != rover.facing {
        t.Fatalf("Moving is failed")
    }
}

func TestRoverMovingWithFacingWest(t *testing.T) {
    facing := WEST
    rover := newRover(facing)

    rover.MoveForward()
    if rover.x != x - 1 {
        t.Fatalf("Moving is failed")
    }

    if rover.y != y {
        t.Fatalf("Moving is failed")
    }

    if rover.facing != rover.facing {
        t.Fatalf("Moving is failed")
    }
}

func TestRoverMovingWithFacingNorth(t *testing.T) {
    facing := NORTH
    rover := newRover(facing)

    rover.MoveForward()
    if rover.y != y + 1 {
        t.Fatalf("Moving is failed")
    }

    if rover.x != x {
        t.Fatalf("Moving is failed")
    }

    if rover.facing != rover.facing {
        t.Fatalf("Moving is failed")
    }
}

func TestRoverTurnLeft(t *testing.T) {
    facing := EAST
    rover := newRover(facing)

    rover.TurnLeft()
    if rover.facing != NORTH {
        t.Fatalf("Turn left is failed")
    }

    rover.TurnLeft()
    if rover.facing != WEST {
        t.Fatalf("Turn left is failed")
    }

    rover.TurnLeft()
    if rover.facing != SOUTH {
        t.Fatalf("Turn left is failed")
    }

    rover.TurnLeft()
    if rover.facing != EAST {
        t.Fatalf("Turn left is failed")
    }
}

func TestRoverTurnRight(t *testing.T) {
    facing := EAST
    rover := newRover(facing)

    rover.TurnRight()
    if rover.facing != SOUTH {
        t.Fatalf("Turn right is failed")
    }

    rover.TurnRight()
    if rover.facing != WEST {
        t.Fatalf("Turn right is failed")
    }

    rover.TurnRight()
    if rover.facing != NORTH {
        t.Fatalf("Turn right is failed")
    }

    rover.TurnRight()
    if rover.facing != EAST {
        t.Fatalf("Turn right is failed")
    }
}

func newRover(facing int) *Rover {
    return &Rover{x:x, y:y, facing:facing}
}
