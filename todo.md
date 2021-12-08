# TODO

* Add a 'tile-stride' to `Spritesheet`. Not all sheets will be well-compacted.
* Animate sprites.
  * Have `tickNum` in `game`
  * And multiple `Tiles` in `Sprite`
* Give classes pointer to Game, don't pass it in during Draw()
* Sprites shouldn't be told to draw on the grid, they should just be
  told to draw, and they decide to be on the grid or not.
* `Camera` or some other abstraction to adjust coordinates, so the
  game can use normal coordinate system, like lower-left is origin,
  or center is origin.

### Animation

Simple things like the flag just constantly animate between the two
states, depending on the `tickNum`.

Generally, though:

* Have a normal or default state (tile) for when they are standing around
  or doing nothing.
* Then, when they start to move, you record the `tickNum` for when the
  movement started, and base the animation on relative to this
  *'tick-duration'*.
