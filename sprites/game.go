package sprites

import (
  "github.com/hajimehoshi/ebiten/v2"
  "time"
)

const (
  width  = 640
  height = 480

  ss1TileWidth  = 18
  ss1TileHeight = 18
  ss1TileXNum   = 20

  ss2TileWidth  = 24
  ss2TileHeight = 24
  ss2TileXNum   = 9
)


type Game struct {
  Grid            Grid
  Spritesheets []*Spritesheet
  Sprites      []*Sprite

  start         time.Time
  tickNum       int64               /* number of 10ths of a second since start */
}

func NewGame(gridX, gridY int) *Game {
  g := Game{
    Grid: Grid{
      CellWidth:  gridX,
      CellHeight: gridY,
    },
    Spritesheets: make([]*Spritesheet, 0),
    Sprites:      make([]*Sprite, 0),

    start:        time.Now(),
    tickNum:      0,
  }

  return &g
}

func (g *Game) NewSpritesheet(buf []byte, twidth, theight, tilexnum,tileynum int) *Spritesheet {
  ss := NewSpritesheet(buf, twidth, theight, tilexnum, tileynum, &g.Grid)

  g.Spritesheets = append(g.Spritesheets, ss)

  return ss
}

func (g *Game) NewSprite(ss *Spritesheet, id int) *Sprite {
  s := ss.NewSprite(id)

  g.Sprites = append(g.Sprites, s)

  return s
}



func (g *Game) Update() error {
  return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
  return width, height
}

func (g *Game) Draw(screen *ebiten.Image) {
  // Draw each tile with each DrawImage call.
  // As the source images of all DrawImage calls are always same,
  // this rendering is done very efficiently.
  // For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage

  elapsed := time.Now().Sub(g.start)
  g.tickNum = elapsed.Milliseconds() / 100

  for i, sprite := range g.Sprites {
    sprite.GridDraw(g, screen, i, i)
  }

  //cx, cy := 0, 3
  //t := 28
  //
  //op := &ebiten.DrawImageOptions{}
  //op.GeoM.Translate(float64(cx*ss1TileWidth), float64(cy*ss1TileHeight))
  //
  //tx := (t % ss1TileXNum) * ss1TileWidth
  //ty := (t / ss1TileXNum) * ss1TileHeight
  //
  //screen.DrawImage(tilesImage.SubImage(image.Rect(tx, ty, tx+ss1TileWidth, ty+ss1TileHeight)).(*ebiten.Image), op)
  ////ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

