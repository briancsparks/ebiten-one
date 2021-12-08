package sprites

import (
  "bytes"
  "fmt"
  "github.com/hajimehoshi/ebiten/v2"
  "image"
  _ "image/png"
  "log"
)

// =================================================================================================================

type Grid struct {
  // Pixels
  CellWidth  int
  CellHeight int
}

func (g *Grid) GridPoint(tx, ty int) image.Point  {
  return image.Point{
    X: tx*g.CellWidth,
    Y: ty*g.CellHeight,
  }
}

// =================================================================================================================

type Spritesheet struct {
  tiles         *ebiten.Image
  grid          *Grid

  // Pixels
  width       int
  height      int

  // Pixels
  tileWidth    int
  tileHeight   int

  // Width, height in tiles
  tileXNum int
  tileYNum int
}

func xy(p image.Point) (int, int) {
  return p.X, p.Y
}

func NewSpritesheet(buf []byte, twidth, theight, tilexnum, tileynum int, grid *Grid) *Spritesheet {
  tilesIm, _, err := image.Decode(bytes.NewReader(buf))
  check(err)

  width, height := xy(tilesIm.Bounds().Size())

  ss := Spritesheet{
    tiles:      ebiten.NewImageFromImage(tilesIm),
    grid:       grid,
    width:      width,
    height:     height,
    tileWidth:  twidth,
    tileHeight: theight,
    tileXNum:   tilexnum,       /*width  / twidth*/
    tileYNum:   tileynum,       /*height / theight*/
  }

  fmt.Printf("image: (%3d,%3d), user-tile: (%2d,%2d), user: %2d, calc: %2d / user: %2d, calc: %2d\n", width, height, ss.tileWidth, ss.tileHeight, ss.tileXNum, width / twidth, ss.tileYNum, height / theight)

  return &ss





  //ss := Spritesheet{
  //  tiles:      ebiten.NewImageFromImage(tilesIm),
  //  grid:       grid,
  //  width:      width,
  //  height:     height,
  //  tileWidth:  twidth,
  //  tileHeight: theight,
  //  tileXNum:   tilexnum,       /*width / twidth*/
  //  //tileYNum:   height / theight,
  //}
  //
  //fmt.Printf("image: (%3d,%3d), user-tile: (%2d,%2d), user: %2d, calc: %2d\n", width, height, ss.tileWidth, ss.tileHeight, ss.tileXNum, width / twidth)
  //
  //return &ss
}


func NewSpritesheet0(buf []byte, twidth, theight, tilexnum int, grid *Grid) *Spritesheet {
  tilesIm, _, err := image.Decode(bytes.NewReader(buf))
  check(err)

  width, height := xy(tilesIm.Bounds().Size())

  ss := Spritesheet{
    tiles:      ebiten.NewImageFromImage(tilesIm),
    grid:       grid,
    width:      width,
    height:     height,
    tileWidth:  twidth,
    tileHeight: theight,
    tileXNum:   tilexnum,       /*width / twidth*/
    //tileYNum:   height / theight,
  }

  fmt.Printf("image: (%3d,%3d), user-tile: (%2d,%2d), user: %2d, calc: %2d\n", width, height, ss.tileWidth, ss.tileHeight, ss.tileXNum, width / twidth)

  return &ss
}

func (ss *Spritesheet) SpriteBounds(id int) image.Rectangle {
  tx := (id % ss.tileXNum) * ss.tileWidth
  ty := (id / ss.tileXNum) * ss.tileHeight

  return image.Rect(tx, ty, tx+ss.tileWidth, ty+ss.tileHeight)
}

func (ss *Spritesheet) GridPoint(tx, ty int) image.Point {
  return image.Point{
    X: tx * ss.grid.CellWidth,
    Y: ty * ss.grid.CellHeight,
  }
  //return image.Point{
  //  X: tx*ss.tileWidth,
  //  Y: ty*ss.tileHeight,
  //}
}

func (ss *Spritesheet) NewTile(id int) *Tile {
  r := ss.SpriteBounds(id)

  t := Tile{
    ss: ss,
    subImage: ss.tiles.SubImage(r).(*ebiten.Image),
    id: id,
  }

  return &t
}


func (ss *Spritesheet) NewSprite(id int) *Sprite {
  r := ss.SpriteBounds(id)

  t := Tile{
    ss: ss,
    subImage: ss.tiles.SubImage(r).(*ebiten.Image),
    id: id,
  }

  s := Sprite{
    ss: ss,
    tiles: make([]*Tile, 0),
  }
  s.tiles = append(s.tiles, &t)

  return &s
}

// =================================================================================================================


type Tile struct {
  ss        *Spritesheet
  subImage  *ebiten.Image
  id        int
}

func (t *Tile) GridDraw(screen *ebiten.Image, tx,ty int)  {
  x,y := xy(t.ss.GridPoint(tx, ty))
  op := &ebiten.DrawImageOptions{}
  op.GeoM.Translate(float64(x), float64(y))

  screen.DrawImage(t.subImage, op)
}


// =================================================================================================================


type Sprite struct {
  ss        *Spritesheet
  tiles   []*Tile
}

func (s *Sprite) GridDraw(g *Game, screen *ebiten.Image, tx,ty int)  {
  //x,y := xy(s.ss.GridPoint(tx, ty))
  //op := &ebiten.DrawImageOptions{}
  //op.GeoM.Translate(float64(x), float64(y))

  for _, tile := range s.tiles {
    tile.GridDraw(screen, tx,ty)
  }
}





// =================================================================================================================


func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

