package marsrover

import "testing"

func TestRoverOutofPlateau(t *testing.T) {
    plateau := CreatePlateauFromInput("5 5")
    roverArr := [...]*Rover{
        &Rover{x:6, y:3, facing:EAST},
        &Rover{x:3, y:6, facing:EAST},
        &Rover{x:-1, y:3, facing:EAST},
        &Rover{x:3, y:-1, facing:EAST},
    }

    for _, r := range roverArr {
        if !plateau.IsRoverOutofArea(r) {
            t.Errorf("Rover(%v) is still in the area of plateau (%v)?", r, plateau)
        }
    }
}
