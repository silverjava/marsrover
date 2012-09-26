package marsrover

type Rover struct {
    x, y, facing int
}

func (r *Rover) MoveForward() (oldX, oldY int) {
    oldX = r.x
    oldY = r.y

    switch r.facing {
        case EAST: r.x++
        case SOUTH: r.y--
        case WEST: r.x--
        case NORTH: r.y++
    }
    return
}

func (r *Rover) TurnLeft() {
    r.facing = (r.facing + 3) % 4
}

func (r *Rover) TurnRight() {
    r.facing = (r.facing + 1) % 4
}
