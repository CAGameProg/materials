# Asset Loader

You may also want to look at the [notes](./notes) directory.

## Set up

You will need to download a number of assets first. I recommend downloading the whole repository by going to the
[main github page](https://github.com/cagameprog/materials) and then pressing `Clone or download` and then `Download ZIP`.

In the zip file you'll find an `assets` folder. This is the folder which stores images, sounds, and fonts. You should load
from this folder in your program.

## Code

Write a program that will load all images, fonts, and sounds from a certain folder (in this case the `assets` folder you downloaded above)
into maps to be later used by your game.

Write a struct to store this called `Resources`. It should store maps for images (`.png`), fonts (`.ttf`),
and sounds (`.wav` or `.ogg`) where the key of each element is the string filename and the value is the
object. Here is the struct you should use:

```go
type Resources struct {
    images map[string]*sf.Texture
    fonts  map[string]*sf.Font
    sounds map[string]*sf.SoundBuffer
}
```

You should write 3 methods:

* `func (r *Resources) LoadImages(dir string)`: loads all images in `dir`
* `func (r *Resources) LoadFonts(dir string)`: loads all fonts in `dir`
* `func (r *Resources) LoadSounds(dir string)`: loads all sounds in `dir`

Here is how to load each filetype:

Image:
```go
img := sf.NewTexture(filename)
```

Font:
```go
font := sf.NewFont(filename)
```

Sound:
```go
sound := sf.NewSoundBuffer(filename)
```

Hint: in order to search recursively you might want to refresh your memory and
take a look at the [`gofind`](../unix-tools/solutions/gofind.go) program we made before break.

And then you should make a function:

* `func NewResources(dir string) *Resources`

The constructor should create a new resources struct, load all the images from 
`dir + "/images"`, load the sounds from `dir + "/sounds"` and load all the fonts from `dir + "/fonts"`
(you should use the methods that you wrote above to do this).

Please put all this code in an `assets.go` file.

Then from your `main.go` file, you can create a new resources object:

```go
func main() {
    res := NewResources("./assets") // You will write this function

    // Now you can use textures and assets to play a sound for example:
    sound := sf.NewSound(res.sounds["sfx_laser1.ogg"])
    sound.Play()

    // Or you could make a new sprite from an image:
    sprite := sf.NewSprite(res.images["playerShip1_blue.png"])
}
```

*Make sure that your program (the `main.go` and `assets.go` files) is in the same directory as the `assets` directory.
