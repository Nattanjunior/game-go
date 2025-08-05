package game

import (
	"game-go/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image    *ebiten.Image
	position Vector
}

func NewPlayer() *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	half := float64(bounds.Dx() / 2)

	position := Vector{
		X: (screenWidth / 2) - half,
		Y: 500,
	}

	return &Player{image: image, position: position}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// posicao do player
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.image, op)
}
