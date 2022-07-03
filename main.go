package main

import (
    "bytes"
    "image"
    "image/color"
    _ "image/png"
    "log"
    "fmt"
    "math/rand"
    "os"
    "time"

    "cardmemorygame/assets"

    "golang.org/x/image/font"
    "golang.org/x/image/font/gofont/goregular"

    "github.com/golang/freetype/truetype"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text"
    "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"

    //"net/http"
    //_ "net/http/pprof"
)

var (
    err error
    ci *ebiten.Image
    cci *ebiten.Image
    playpauseImage *ebiten.Image
    cardImage *ebiten.Image
    quitImage *ebiten.Image
    bgimage image.Image
    startImage *ebiten.Image
    bgImage *ebiten.Image
    leftImage *ebiten.Image
    rightImage *ebiten.Image
    frontImage *ebiten.Image
    initial bool = true
    start bool = false
    flipped bool = false
    flippedcard [2]int
    flippedcards [][2]int
    newn int = 4
    clicked bool = false
    cardclicked [2]int
    pairs int = 0
    won bool = false
    sleep bool = false
    sleept int = 1
    t [2]time.Time
    dura time.Duration
    pdura time.Duration
    pdurat time.Duration
    mi, se int
    pause bool = false
    justpaused bool = false
    rights int = 0
    lefts int = 0
    fronts int = 0
    fon *truetype.Font
    fo2 font.Face
    fo4 font.Face
)

var cardz = make(map[[2]int]*ebiten.Image)
var cards = make(map[[2]int]*ebiten.Image)

type Game struct {
    n int
}

func (g *Game) N(n int) {
    g.n = n
}

func (g *Game) GetN() int {
    return g.n
}

func (g *Game) Restart() {
    flipped = false
    flippedcard = [2]int{}
    flippedcards = [][2]int{}
    clicked = false
    cardclicked = [2]int{}
    pairs = 0
    won = false
    sleep = false
    sleept = 1
    t = [2]time.Time{}
    dura, err = time.ParseDuration("0s")
    if err != nil {
        log.Fatal(err)
    }
    pdura, err = time.ParseDuration("0s")
    if err != nil {
        log.Fatal(err)
    }
    pdurat, err = time.ParseDuration("0s")
    if err != nil {
        log.Fatal(err)
    }
    rights = 0
    lefts = 0
    fronts = 0
    rand.Seed(time.Now().UnixNano())
    for a := 0; a < g.GetN(); a++ {
        for b := 0; b <= g.GetN() / 2; b++ {
            cards[[2]int{a, b}] = cardImage
            switch rand.Intn(3) {
            case 0:
                cardz[[2]int{a, b}] = rightImage
                rights++
            case 1:
                cardz[[2]int{a, b}] = leftImage
                lefts++
            case 2:
                cardz[[2]int{a, b}] = frontImage
                fronts++
            }
        }
    }
    for c := 0; c < g.GetN(); c++ {
        for d := g.GetN() / 2; d < g.GetN(); d++ {
            cards[[2]int{c, d}] = cardImage
            switch {
            case cardz[[2]int{c, d - (newn / 2)}] == leftImage:
                if rights > 0 {
                    cardz[[2]int{c, d}] = rightImage
                    rights--
                } else if lefts > 0 {
                    cardz[[2]int{c, d}] = leftImage
                    lefts--
                } else {
                    cardz[[2]int{c, d}] = frontImage
                    fronts--
                }
            case cardz[[2]int{c, d - (newn / 2)}] == rightImage:
                if fronts > 0 {
                    cardz[[2]int{c, d}] = frontImage
                    fronts--
                } else if rights > 0 {
                    cardz[[2]int{c, d}] = rightImage
                    rights--
                } else {
                    cardz[[2]int{c, d}] = leftImage
                    lefts--
                }
            case cardz[[2]int{c, d - (newn / 2)}] == frontImage:
                if lefts > 0 {
                    cardz[[2]int{c, d}] = leftImage
                    lefts--
                } else if fronts > 0 {
                    cardz[[2]int{c, d}] = frontImage
                    fronts--
                } else {
                    cardz[[2]int{c, d}] = rightImage
                    rights--
                }
            }
        }
    }
}

func (g *Game) Update() error {
    w, h := ebiten.WindowSize()
    if ebiten.IsFullscreen() {
        w, h = ebiten.ScreenSizeInFullscreen()
    }
    bw, bh := frontImage.Size()
    sw, sh := startImage.Size()
    qw, qh := quitImage.Size()
    if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
        ebiten.SetFullscreen(!ebiten.IsFullscreen())
    }
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        xx, yy := ebiten.CursorPosition()
        if xx >= w - (w / 18) && xx <= w - (w / 18) + (qw / 8) && yy >= h / 38 && yy <= (h / 38) + (qh / 8) {
            os.Exit(0)
        }
    }
    if won {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            x, y := ebiten.CursorPosition()
            if x >= w / 9 && x <= (w / 9) + 174 && y >= int(float64(h) / 1.375) && y <= int(float64(h) / 1.375) + 44 {
                g.Restart()
                initial = true
            }
        }
    }
    if !start {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            x, y := ebiten.CursorPosition()
            if x >= w - int(float64(w) / 8.714) && x <= w - int(float64(w) / 8.714) + (sw / 8) && y >= h - int(float64(h) / 3.835) && y <= h - int(float64(h) / 3.835) + (sh / 8) {
                g.Restart()
                start = true
            }
        }
    } else {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            x, y := ebiten.CursorPosition()
            if x >= 20 && x <= 70 && y >= 60 && y <= 110 {
                pause = !pause
                if pause {
                    justpaused = true
                } else {
                    pdurat += pdura
                }
            }
            switch g.GetN() {
            case 2:
                switch {
                case x >= (w / 2) - 10 - ((3 * bw) / 8) && x <= (w / 2) - 10:
                    switch {
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{0, 0}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{1, 0}
                    }
                case x >= (w / 2) + 10 && x <= (w / 2) + 10 + ((3 * bw) / 8):
                    switch {
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{0, 1}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{1, 1}
                    }
                }
            case 4:
                switch {
                case x >= (w / 2) - 30 - ((3 * bw) / 4) && x <= (w / 2) - 30 - ((3 * bw) / 8):
                    switch {
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{0, 0}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{1, 0}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{2, 0}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{3, 0}
                    }
                case x >= (w / 2) - 10 - ((3 * bw) / 8) && x <= (w / 2) - 10:
                    switch {
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{0, 1}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{1, 1}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{2, 1}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{3, 1}
                    }
                case x >= (w / 2) + 10 && x <= (w / 2) + 10 + ((3 * bw) / 8):
                    switch {
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{0, 2}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{1, 2}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{2, 2}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{3, 2}
                    }
                case x >= (w / 2) + 30 + ((3 * bw) / 8) && x <= (w / 2) + 30 + ((3 * bw) / 4):
                    switch {
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{0, 3}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{1, 3}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{2, 3}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{3, 3}
                    }
                }
            case 6:
                switch {
                case x >= (w / 2) - 50 - ((9 * bw) / 8) && x <= (w / 2) - 50 - ((3 * bw) / 4):
                    switch {
                    case y >= (h / 2) - 50 - ((9 * bh) / 8) && y <= (h / 2) - 50 - ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{0, 0}
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{1, 0}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{2, 0}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{3, 0}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{4, 0}
                    case y >= (h / 2) + 50 + ((3 * bh) / 4) && y <= (h / 2) + 50 + ((9 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{5, 0}
                    }
                case x >= (w / 2) - 30 - ((3 * bw) / 4) && x <= (w / 2) - 30 - ((3 * bw) / 8):
                    switch {
                    case y >= (h / 2) - 50 - ((9 * bh) / 8) && y <= (h / 2) - 50 - ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{0, 1}
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{1, 1}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{2, 1}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{3, 1}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{4, 1}
                    case y >= (h / 2) + 50 + ((3 * bh) / 4) && y <= (h / 2) + 50 + ((9 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{5, 1}
                    }
                case x >= (w / 2) - 10 - ((3 * bw) / 8) && x <= (w / 2) - 10:
                    switch {
                    case y >= (h / 2) - 50 - ((9 * bh) / 8) && y <= (h / 2) - 50 - ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{0, 2}
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{1, 2}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{2, 2}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{3, 2}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{4, 2}
                    case y >= (h / 2) + 50 + ((3 * bh) / 4) && y <= (h / 2) + 50 + ((9 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{5, 2}
                    }
                case x >= (w / 2) + 10 && x <= (w / 2) + 10 + ((3 * bw) / 8):
                    switch {
                    case y >= (h / 2) - 50 - ((9 * bh) / 8) && y <= (h / 2) - 50 - ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{0, 3}
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{1, 3}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{2, 3}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{3, 3}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{4, 3}
                    case y >= (h / 2) + 50 + ((3 * bh) / 4) && y <= (h / 2) + 50 + ((9 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{5, 3}
                    }
                case x >= (w / 2) + 30 + ((3 * bw) / 8) && x <= (w / 2) + 30 + ((3 * bw) / 4):
                    switch {
                    case y >= (h / 2) - 50 - ((9 * bh) / 8) && y <= (h / 2) - 50 - ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{0, 4}
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{1, 4}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{2, 4}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{3, 4}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{4, 4}
                    case y >= (h / 2) + 50 + ((3 * bh) / 4) && y <= (h / 2) + 50 + ((9 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{5, 4}
                    }
                case x >= (w / 2) + 50 + ((3 * bw) / 4) && x <= (w / 2) + 50 + ((9 * bw) / 8):
                    switch {
                    case y >= (h / 2) - 50 - ((9 * bh) / 8) && y <= (h / 2) - 50 - ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{0, 5}
                    case y >= (h / 2) - 30 - ((3 * bh) / 4) && y <= (h / 2) - 30 - ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{1, 5}
                    case y >= (h / 2) - 10 - ((3 * bh) / 8) && y <= (h / 2) - 10:
                        clicked = true
                        cardclicked = [2]int{2, 5}
                    case y >= (h / 2) + 10 && y <= (h / 2) + 10 + ((3 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{3, 5}
                    case y >= (h / 2) + 30 + ((3 * bh) / 8) && y <= (h / 2) + 30 + ((3 * bh) / 4):
                        clicked = true
                        cardclicked = [2]int{4, 5}
                    case y >= (h / 2) + 50 + ((3 * bh) / 4) && y <= (h / 2) + 50 + ((9 * bh) / 8):
                        clicked = true
                        cardclicked = [2]int{5, 5}
                    }
                }
            }
        }
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    w, h := ebiten.WindowSize()
    if ebiten.IsFullscreen() {
        w, h = ebiten.ScreenSizeInFullscreen()
    }
    dur, err := time.ParseDuration("1s")
    if err != nil {
        log.Fatal(err)
    }

    if initial {
        ebitenutil.DebugPrintAt(screen, "Number of rows and columns: " + fmt.Sprint(newn), w / 2, h / 2)
        ebitenutil.DebugPrintAt(screen, "Press either 2, 4, or 6 to update", w / 2, h / 2 + 20)
        if ebiten.IsKeyPressed(ebiten.KeyEnter) {
            g.N(newn)
            rand.Seed(time.Now().UnixNano())
            for a := 0; a < g.GetN(); a++ {
                for b := 0; b <= g.GetN() / 2; b++ {
                    cards[[2]int{a, b}] = cardImage
                    switch rand.Intn(3) {
                    case 0:
                        cardz[[2]int{a, b}] = rightImage
                        rights++
                    case 1:
                        cardz[[2]int{a, b}] = leftImage
                        lefts++
                    case 2:
                        cardz[[2]int{a, b}] = frontImage
                        fronts++
                    }
                }
            }
            for c := 0; c < g.GetN(); c++ {
                for d := g.GetN() / 2; d < g.GetN(); d++ {
                    cards[[2]int{c, d}] = cardImage
                    switch {
                    case cardz[[2]int{c, d - (newn / 2)}] == leftImage:
                        if rights > 0 {
                            cardz[[2]int{c, d}] = rightImage
                            rights--
                        } else if lefts > 0 {
                            cardz[[2]int{c, d}] = leftImage
                            lefts--
                        } else {
                            cardz[[2]int{c, d}] = frontImage
                            fronts--
                        }
                    case cardz[[2]int{c, d - (newn / 2)}] == rightImage:
                        if fronts > 0 {
                            cardz[[2]int{c, d}] = frontImage
                            fronts--
                        } else if rights > 0 {
                            cardz[[2]int{c, d}] = rightImage
                            rights--
                        } else {
                            cardz[[2]int{c, d}] = leftImage
                            lefts--
                        }
                    case cardz[[2]int{c, d - (newn / 2)}] == frontImage:
                        if lefts > 0 {
                            cardz[[2]int{c, d}] = leftImage
                            lefts--
                        } else if fronts > 0 {
                            cardz[[2]int{c, d}] = frontImage
                            fronts--
                        } else {
                            cardz[[2]int{c, d}] = rightImage
                            rights--
                        }
                    }
                }
            }
            initial = false
        }
        if inpututil.IsKeyJustPressed(ebiten.Key2) {
            newn = 2
        }
        if inpututil.IsKeyJustPressed(ebiten.Key4) {
            newn = 4
        }
        if inpututil.IsKeyJustPressed(ebiten.Key6) {
            newn = 6
        }
    } else {
        if !won && !pause && start {
            if t[0].IsZero() {
                dura, err = time.ParseDuration("0s")
                if err != nil {
                    log.Fatal(err)
                }
                t[0] = time.Now()
            } else {
                dura = time.Now().Sub(t[0]) - pdurat
            }
        }
        if !pause {
            mi = int(dura.Minutes()) % 60
            se = int(dura.Seconds()) % 60
        } else {
            if justpaused {
                pdura, err = time.ParseDuration("0s")
                if err != nil {
                    log.Fatal(err)
                }
                t[1] = time.Now()
                justpaused = false
            } else {
                pdura = time.Now().Sub(t[1])
            }
        }
        screen.DrawImage(bgImage, &ebiten.DrawImageOptions{})
        text.Draw(screen, fmt.Sprintf("%02d:%02d", mi, se), fo2, 20, 40, color.Black)
        pgm := ebiten.GeoM{}
        pgm.Scale(0.5, 0.5)
        pgm.Translate(float64(20), float64(60))
        pausesi := [4]int{152, 50, 252, 150}
        if pause {
            pausesi[0] = 29
            pausesi[2] = 129
        }
        screen.DrawImage(
            playpauseImage.SubImage(
                image.Rect(pausesi[0], pausesi[1], pausesi[2], pausesi[3])).(*ebiten.Image),
                &ebiten.DrawImageOptions{
                    GeoM: pgm})
        qgm := ebiten.GeoM{}
        qgm.Scale(0.125, 0.125)
        qgm.Translate(float64(w - (w / 18)), float64(h / 38))
        screen.DrawImage(
            quitImage, &ebiten.DrawImageOptions{
                GeoM: qgm})
        bw, bh := frontImage.Size()
        gm := ebiten.GeoM{}
        gm.Scale(0.375, 0.375)
        switch g.GetN() {
        case 2:
            gm.Translate(float64((w / 2) - 10 - ((3 * bw) / 8)), float64((h / 2) - 10 - ((3 * bh) / 8)))
        case 4:
            gm.Translate(float64((w / 2) - 30 - ((3 * bw) / 4)), float64((h / 2) - 30 - ((3 * bh) / 4)))
        case 6:
            gm.Translate(float64((w / 2) - 50 - ((9 * bw) / 8)), float64((h / 2) - 50 - ((9 * bh) / 8)))
        }
        for a := 0; a < g.GetN(); a++ {
            for b := 0; b < g.GetN(); b++ {
                gm.Translate(-float64(2), -float64(2))
                screen.DrawImage(
                    ci, &ebiten.DrawImageOptions{
                        GeoM: gm})
                gm.Translate(float64(2), float64(2))
                screen.DrawImage(
                    cci, &ebiten.DrawImageOptions{
                        GeoM: gm})
                screen.DrawImage(
                    cards[[2]int{a, b}], &ebiten.DrawImageOptions{
                        GeoM: gm})
                if b < g.GetN() - 1 {
                    gm.Translate(float64(((3 * bw) / 8) + 20), float64(0))
                } else if a < g.GetN() - 1 {
                    gm.Translate(-float64((((3 * bw) / 8) + 20) * (g.GetN() - 1)), float64(((3 * bh) / 8) + 20))
                }
            }
        }
        if sleep {
            if sleept == 0 {
                time.Sleep(dur)
                sleept = 1
                sleep = false
                cards[cardclicked] = cardImage
                cards[flippedcard] = cardImage
            } else {
                sleept--
            }
        }
        if !start {
            sgm := ebiten.GeoM{}
            sgm.Scale(0.125, 0.125)
            sgm.Translate(float64(w) - (float64(w) / 8.714), float64(h) - (float64(h) / 3.835))
            screen.DrawImage(
                startImage, &ebiten.DrawImageOptions{
                    GeoM: sgm})
        }
    }
    if start {
        if len(t) == 0 {
            t[0] = time.Now()
        }
        if clicked && !isAlreadyDone(cardclicked) {
            if !flipped {
                cards[cardclicked] = cardz[cardclicked]
                flippedcard = cardclicked
                flipped = true
                clicked = false
            } else {
                if cardclicked != flippedcard {
                    if cardz[cardclicked] == cardz[flippedcard] {
                        cards[cardclicked] = cardz[cardclicked]
                        flippedcards = append(flippedcards, flippedcard, cardclicked)
                        pairs++
                        if pairs == g.GetN() * (g.GetN() / 2) {
                            won = true
                            start = false
                        }
                    } else {
                        cards[cardclicked] = cardz[cardclicked]
                        sleep = true
                    }
                    flipped = false
                    clicked = false
                }
            }
        }
    }
    if won {
        r := text.BoundString(fo4, "YOU WON")
        wid := r.Max.X - r.Min.X
        hei := r.Max.Y - r.Min.Y
        wgm := ebiten.GeoM{}
        wgm.Translate(float64((w / 2) - (wid / 2) - 22), float64((h / 2) - (hei / 2) - 22))
        wwi := ebiten.NewImage(wid + 44, hei + 44)
        wwi.Fill(color.Black)
        screen.DrawImage(
            wwi, &ebiten.DrawImageOptions{
                GeoM: wgm})
        wgm.Translate(float64(2), float64(2))
        wi := ebiten.NewImage(wid + 40, hei + 40)
        wi.Fill(color.RGBA{241, 162, 47, 255})
        screen.DrawImage(
            wi, &ebiten.DrawImageOptions{
                GeoM: wgm})
        text.Draw(screen, "YOU WON", fo4, (w / 2) - (wid / 2), (h / 2) + (hei / 2), color.RGBA{22, 154, 26, 204})
        r2 := text.BoundString(fo2, "Change Difficulty")
        wid2 := r2.Max.X - r2.Min.X
        hei2 := r2.Max.Y - r2.Min.Y
        dgm := ebiten.GeoM{}
        dgm.Translate(float64(w) / 9.0, float64(h) / 1.375)
        cdi := ebiten.NewImage(wid2 + 24, hei2 + 24)
        cdi.Fill(color.Black)
        screen.DrawImage(
            cdi, &ebiten.DrawImageOptions{
                GeoM: dgm})
        dgm.Translate(float64(2), float64(2))
        di := ebiten.NewImage(wid2 + 20, hei2 + 20)
        di.Fill(color.RGBA{241, 162, 47, 255})
        screen.DrawImage(
            di, &ebiten.DrawImageOptions{
                GeoM: dgm})
        text.Draw(screen, "Change Difficulty", fo2, int(float64(w) / 9.0) + 12, int(float64(h) / 1.375) + 30, color.Black)
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return outsideWidth, outsideHeight
}

func isAlreadyDone(c [2]int) bool {
    for _, a := range flippedcards {
        if a == c {
            return true
        }
    }
    return false
}

func init() {
    fon, err = truetype.Parse(goregular.TTF)
    if err != nil {
        log.Fatal(err)
    }
    fo2 = truetype.NewFace(fon, &truetype.Options{Size: 20})
    fo4 = truetype.NewFace(fon, &truetype.Options{Size: 40})

    bgimage, _, err := image.Decode(bytes.NewReader(platformer.Background_png))
    if err != nil {
        log.Fatal(err)
    }
    bgImage = ebiten.NewImageFromImage(bgimage)

    cardimage, _, err := image.Decode(bytes.NewReader(assets.Card_png))
    if err != nil {
        log.Fatal(err)
    }
    cardImage = ebiten.NewImageFromImage(cardimage)

    quitimage, _, err := image.Decode(bytes.NewReader(assets.QuitButton_png))
    if err != nil {
        log.Fatal(err)
    }
    quitImage = ebiten.NewImageFromImage(quitimage)

    leftimage, _, err := image.Decode(bytes.NewReader(platformer.Left_png))
    if err != nil {
        log.Fatal(err)
    }
    leftImage = ebiten.NewImageFromImage(leftimage)

    rightimage, _, err := image.Decode(bytes.NewReader(platformer.Right_png))
    if err != nil {
        log.Fatal(err)
    }
    rightImage = ebiten.NewImageFromImage(rightimage)

    frontimage, _, err := image.Decode(bytes.NewReader(platformer.MainChar_png))
    if err != nil {
        log.Fatal(err)
    }
    frontImage = ebiten.NewImageFromImage(frontimage)
    bw, bh := frontImage.Size()

    startimage, _, err := image.Decode(bytes.NewReader(assets.StartButton_png))
    if err != nil {
        log.Fatal(err)
    }
    startImage = ebiten.NewImageFromImage(startimage)

    playpauseimage, _, err := image.Decode(bytes.NewReader(assets.PlayPause_png))
    if err != nil {
        log.Fatal(err)
    }
    playpauseImage = ebiten.NewImageFromImage(playpauseimage)

    ci = ebiten.NewImage(bw + (32 / 3), bh + (32 / 3))
    ci.Fill(color.Black)

    cci = ebiten.NewImage(bw, bh)
    cci.Fill(color.RGBA{241, 162, 47, 255})
}

func main() {
    ebiten.SetWindowSize(1280, 720)
    ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
    ebiten.SetWindowSizeLimits(720, 561, -1, -1)
    ebiten.SetWindowTitle("Card Memory Game")

    game := &Game{}

    //go func() {
    //    log.Println(http.ListenAndServe("localhost:6060", nil))
    //}()

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}

