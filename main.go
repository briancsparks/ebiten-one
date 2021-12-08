package main

import (
  "bytes"
  _ "embed"
  "github.com/briancsparks/ebiten-one/sprites"
  "github.com/hajimehoshi/ebiten/v2"
  "image"
  _ "image/png"
  "log"
)

const (
  width  = 640
  height = 480

  ss1TileWidth  = 18
  ss1TileHeight = 18
  ss1TileXNum   = 20
  ss1TileYNum   = 9

  ss2TileWidth  = 24
  ss2TileHeight = 24
  ss2TileXNum   = 9
  ss2TileYNum   = 3
)

//go:embed assets/tiles_packed.png
var tilesBytes []byte

//go:embed assets/characters_packed.png
var charactersBytes []byte

var (
  tilesImageIm      *image.Image
  charactersImageIm *image.Image
  tilesImage        *ebiten.Image
  charactersImage   *ebiten.Image
)

type GameX struct {
  Grid         *sprites.Grid
  Spritesheets []*sprites.Spritesheet
  Sprites      []*sprites.Sprite
}

func init() {
  tilesImageIm, _, err := image.Decode(bytes.NewReader(tilesBytes))
  check(err)
  charactersImageIm, _, err := image.Decode(bytes.NewReader(charactersBytes))
  check(err)

  tilesImage = ebiten.NewImageFromImage(tilesImageIm)
  charactersImage = ebiten.NewImageFromImage(charactersImageIm)
}

func (g *GameX) Update() error {
  return nil
}

func (g *GameX) Layout(outsideWidth, outsideHeight int) (int, int) {
  return width, height
}

func (g *GameX) Draw(screen *ebiten.Image) {
  // Draw each tile with each DrawImage call.
  // As the source images of all DrawImage calls are always same,
  // this rendering is done very efficiently.
  // For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage

  //for i, sprite := range g.Sprites {
  //  sprite.GridDraw(g, screen, i, i)
  //}

  cx, cy := 0, 3
  t := 28

  op := &ebiten.DrawImageOptions{}
  op.GeoM.Translate(float64(cx*ss1TileWidth), float64(cy*ss1TileHeight))

  tx := (t % ss1TileXNum) * ss1TileWidth
  ty := (t / ss1TileXNum) * ss1TileHeight

  screen.DrawImage(tilesImage.SubImage(image.Rect(tx, ty, tx+ss1TileWidth, ty+ss1TileHeight)).(*ebiten.Image), op)
  //ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func main() {

  g := sprites.NewGame(16, 16)

  ss1     := g.NewSpritesheet(tilesBytes, ss1TileWidth, ss1TileHeight, ss1TileXNum, ss1TileYNum)
  ss1s29  := g.NewSprite(ss1, 29)

  ss2     := g.NewSpritesheet(charactersBytes, ss2TileWidth, ss2TileHeight, ss2TileXNum, ss2TileYNum)
  ss2s25  := g.NewSprite(ss2, 25)
  ss2s22  := g.NewSprite(ss2, 22)

  _,_,_    = ss1s29,ss2s25,ss2s22


  grid := &sprites.Grid{CellWidth: 16, CellHeight: 16}

  gx := &GameX{
    Grid:         grid,
    Spritesheets: make([]*sprites.Spritesheet, 0),
    Sprites:      make([]*sprites.Sprite, 0),
  }

  // -------------
  ssx1 := sprites.NewSpritesheet0(tilesBytes, ss1TileWidth, ss1TileHeight, ss1TileXNum, grid)
  gx.Spritesheets = append(gx.Spritesheets, ssx1)

  ssx1s29 := ssx1.NewSprite(29)
  gx.Sprites = append(gx.Sprites, ssx1s29)

  // -------------
  ssx2 := sprites.NewSpritesheet0(charactersBytes, ss2TileWidth, ss2TileHeight, ss2TileXNum, grid)
  gx.Spritesheets = append(gx.Spritesheets, ssx2)

  ssx2s25 := ssx2.NewSprite(25)
  gx.Sprites = append(gx.Sprites, ssx2s25)

  ssx2s22 := ssx2.NewSprite(22)
  gx.Sprites = append(gx.Sprites, ssx2s22)



  ebiten.SetWindowSize(width*2, height*2)
  ebiten.SetWindowTitle("Tiles")
  if err := ebiten.RunGame(g); err != nil {
    log.Fatal(err)
  }
}



func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
