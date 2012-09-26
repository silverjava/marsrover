package marsrover

type Plateau struct {
    width int
    height int
}

func (p *Plateau) IsRoverOutofArea(r *Rover) bool {
    return r.x > p.width || r.y > p.height || r.x < 0 || r.y < 0
}

