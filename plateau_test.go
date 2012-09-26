package marsrover

import "testing"

func TestRoverOutofPlateau(t *testing.T) {
    plateau := CreatePlateauFromInput("5 5")
    rover := &Rover{x:6, y:3, facing:EAST}

    if !plateau.IsRoverOutofArea(rover) {
        t.Errorf("Rover(%v) is still in the area of plateau (%v)?", rover, plateau)
    }

    rover = &Rover{x:3, y:6, facing:EAST}
    if !plateau.IsRoverOutofArea(rover) {
        t.Errorf("Rover(%v) is still in the area of plateau (%v)?", rover, plateau)
    }

    rover = &Rover{x:-1, y:3, facing:EAST}
    if !plateau.IsRoverOutofArea(rover) {
        t.Errorf("Rover(%v) is still in the area of plateau (%v)?", rover, plateau)
    }

    rover = &Rover{x:3, y:-1, facing:EAST}
    if !plateau.IsRoverOutofArea(rover) {
        t.Errorf("Rover(%v) is still in the area of plateau (%v)?", rover, plateau)
    }
}
