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

    "ebitengine-poc/assets"

    "golang.org/x/image/font/gofont/goregular"

    "github.com/golang/freetype/truetype"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text"
    "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
)

var (
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
    newn int = 2
    clicked bool = false
    cardclicked [2]int
    pairs int = 0
    won bool = false
    sleep bool = false
    sleept int = 1
    t0 time.Time
    dura time.Duration
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
    t0 = time.Time{}
    rand.Seed(time.Now().UnixNano())
    for a := 0; a < g.GetN(); a++ {
        for b := 0; b <= g.GetN() / 2; b++ {
            cards[[2]int{a, b}] = frontImage
            if rand.Intn(2) == 1 {
                cardz[[2]int{a, b}] = rightImage
            } else {
                cardz[[2]int{a, b}] = leftImage
            }
        }
    }
    for c := 0; c < g.GetN(); c++ {
        for d := g.GetN() / 2; d < g.GetN(); d++ {
            cards[[2]int{c, d}] = frontImage
            if cardz[[2]int{c, d - (g.GetN() / 2)}] == leftImage {
                cardz[[2]int{c, d}] = rightImage
            } else {
                cardz[[2]int{c, d}] = leftImage
            }
        }
    }
}

func (g *Game) Update() error {
    w, h := ebiten.ScreenSizeInFullscreen()
    bw, bh := frontImage.Size()
    sw, sh := startImage.Size()
    qw, qh := quitImage.Size()
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        xx, yy := ebiten.CursorPosition()
        if xx >= w - (w / 22) && xx <= w - (w / 22) + (qw / 8) && yy >= h / 38 && yy <= (h / 38) + (qh / 8) {
            os.Exit(0)
        }
    }
    if won {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            x, y := ebiten.CursorPosition()
            if x >= w / 8 && x <= (w / 8) + 344 && y >= int(float64(h) / 1.5) && y <= int(float64(h) / 1.5) + 84 {
                g.Restart()
                initial = true
            }
        }
    }
    if !start {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            x, y := ebiten.CursorPosition()
            if x >= w - int(float64(w) / 4.5) && x <= w - int(float64(w) / 4.5) + (sw / 4) && y >= h - int(float64(h) / 3.835) && y <= h - int(float64(h) / 3.835) + (sh / 4) {
                g.Restart()
                start = true
            }
        }
    } else {
        if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
            x, y := ebiten.CursorPosition()
            switch g.GetN() {
                case 2:
                    switch {
                        case x >= (w / 2) - 10 - (bw / 2) && x <= (w / 2) - 10:
                            switch {
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{0, 0}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{1, 0}
                            }
                        case x >= (w / 2) + 10 && x <= (w / 2) + 10 + (bw / 2):
                            switch {
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{0, 1}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{1, 1}
                            }
                    }
                case 4:
                    switch {
                        case x >= (w / 2) - 30 - bw && x <= (w / 2) - 30 - (bw / 2):
                            switch {
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{0, 0}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{1, 0}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{2, 0}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{3, 0}
                            }
                        case x >= (w / 2) - 10 - (bw / 2) && x <= (w / 2) - 10:
                            switch {
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{0, 1}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{1, 1}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{2, 1}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{3, 1}
                            }
                        case x >= (w / 2) + 10 && x <= (w / 2) + 10 + (bw / 2):
                            switch {
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{0, 2}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{1, 2}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{2, 2}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{3, 2}
                            }
                        case x >= (w / 2) + 30 + (bw / 2) && x <= (w / 2) + 30 + bw:
                            switch {
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{0, 3}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{1, 3}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{2, 3}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{3, 3}
                            }
                    }
                case 6:
                    switch {
                        case x >= (w / 2) - 50 - ((3 * bw) / 2) && x <= (w / 2) - 50 - bw:
                            switch {
                                case y >= (h / 2) - 50 - ((3 * bh) / 2) && y <= (h / 2) - 50 - bh:
                                    clicked = true
                                    cardclicked = [2]int{0, 0}
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{1, 0}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{2, 0}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{3, 0}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{4, 0}
                                case y >= (h / 2) + 50 + bh && y <= (h / 2) + 50 + ((3 * bh) / 2):
                                    clicked = true
                                    cardclicked = [2]int{5, 0}
                            }
                        case x >= (w / 2) - 30 - bw && x <= (w / 2) - 30 - (bw / 2):
                            switch {
                                case y >= (h / 2) - 50 - ((3 * bh) / 2) && y <= (h / 2) - 50 - bh:
                                    clicked = true
                                    cardclicked = [2]int{0, 1}
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{1, 1}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{2, 1}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{3, 1}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{4, 1}
                                case y >= (h / 2) + 50 + bh && y <= (h / 2) + 50 + ((3 * bh) / 2):
                                    clicked = true
                                    cardclicked = [2]int{5, 1}
                            }
                        case x >= (w / 2) - 10 - (bw / 2) && x <= (w / 2) - 10:
                            switch {
                                case y >= (h / 2) - 50 - ((3 * bh) / 2) && y <= (h / 2) - 50 - bh:
                                    clicked = true
                                    cardclicked = [2]int{0, 2}
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{1, 2}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{2, 2}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{3, 2}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{4, 2}
                                case y >= (h / 2) + 50 + bh && y <= (h / 2) + 50 + ((3 * bh) / 2):
                                    clicked = true
                                    cardclicked = [2]int{5, 2}
                            }
                        case x >= (w / 2) + 10 && x <= (w / 2) + 10 + (bw / 2):
                            switch {
                                case y >= (h / 2) - 50 - ((3 * bh) / 2) && y <= (h / 2) - 50 - bh:
                                    clicked = true
                                    cardclicked = [2]int{0, 3}
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{1, 3}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{2, 3}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{3, 3}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{4, 3}
                                case y >= (h / 2) + 50 + bh && y <= (h / 2) + 50 + ((3 * bh) / 2):
                                    clicked = true
                                    cardclicked = [2]int{5, 3}
                            }
                        case x >= (w / 2) + 30 + (bw / 2) && x <= (w / 2) + 30 + bw:
                            switch {
                                case y >= (h / 2) - 50 - ((3 * bh) / 2) && y <= (h / 2) - 50 - bh:
                                    clicked = true
                                    cardclicked = [2]int{0, 4}
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{1, 4}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{2, 4}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{3, 4}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{4, 4}
                                case y >= (h / 2) + 50 + bh && y <= (h / 2) + 50 + ((3 * bh) / 2):
                                    clicked = true
                                    cardclicked = [2]int{5, 4}
                            }
                        case x >= (w / 2) + 50 + bw && x <= (w / 2) + 50 + ((3 * bw) / 2):
                            switch {
                                case y >= (h / 2) - 50 - ((3 * bh) / 2) && y <= (h / 2) - 50 - bh:
                                    clicked = true
                                    cardclicked = [2]int{0, 5}
                                case y >= (h / 2) - 30 - bh && y <= (h / 2) - 30 - (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{1, 5}
                                case y >= (h / 2) - 10 - (bh / 2) && y <= (h / 2) - 10:
                                    clicked = true
                                    cardclicked = [2]int{2, 5}
                                case y >= (h / 2) + 10 && y <= (h / 2) + 10 + (bh / 2):
                                    clicked = true
                                    cardclicked = [2]int{3, 5}
                                case y >= (h / 2) + 30 + (bh / 2) && y <= (h / 2) + 30 + bh:
                                    clicked = true
                                    cardclicked = [2]int{4, 5}
                                case y >= (h / 2) + 50 + bh && y <= (h / 2) + 50 + ((3 * bh) / 2):
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
    w, h := ebiten.ScreenSizeInFullscreen()
    dur, err := time.ParseDuration("1s")
    if err != nil {
        log.Fatal(err)
    }
    fon, err := truetype.Parse(goregular.TTF)
    if err != nil {
        log.Fatal(err)
    }

    if initial {
        ebitenutil.DebugPrintAt(screen, "Number of rows and columns: " + fmt.Sprint(newn), w / 2, h / 2)
        ebitenutil.DebugPrintAt(screen, "Press either 2, 4, or 6 to update", w / 2, h / 2 + 20)
        if ebiten.IsKeyPressed(ebiten.KeyEnter) {
            g.N(newn)
            rand.Seed(time.Now().UnixNano())
            for a := 0; a < newn; a++ {
                for b := 0; b <= newn / 2; b++ {
                    cards[[2]int{a, b}] = frontImage
                    if rand.Intn(2) == 1 {
                        cardz[[2]int{a, b}] = rightImage
                    } else {
                        cardz[[2]int{a, b}] = leftImage
                    }
                }
            }
            for c := 0; c < newn; c++ {
                for d := newn / 2; d < newn; d++ {
                    cards[[2]int{c, d}] = frontImage
                    if cardz[[2]int{c, d - (newn / 2)}] == leftImage {
                        cardz[[2]int{c, d}] = rightImage
                    } else {
                        cardz[[2]int{c, d}] = leftImage
                    }
                }
            }
            initial = false
        }
        if ebiten.IsKeyPressed(ebiten.Key2) {
            newn = 2
        }
        if ebiten.IsKeyPressed(ebiten.Key4) {
            newn = 4
        }
        if ebiten.IsKeyPressed(ebiten.Key6) {
            newn = 6
        }
    } else {
        fo := truetype.NewFace(fon, &truetype.Options{Size: 20})
        if !won {
            if t0.IsZero() {
                dura, err = time.ParseDuration("0s")
                if err != nil {
                    log.Fatal(err)
                }
            } else {
                dura = time.Now().Sub(t0)
            }
        }
        mi := int(dura.Minutes()) % 60
        se := int(dura.Seconds()) % 60
        clearBG()
        text.Draw(bgImage, fmt.Sprintf("%02d:%02d", mi, se), fo, 20, 40, color.Black)
        screen.DrawImage(bgImage, &ebiten.DrawImageOptions{})
        qgm := ebiten.GeoM{}
        qgm.Scale(0.125, 0.125)
        qgm.Translate(float64(w - (w / 22)), float64(h / 38))
        screen.DrawImage(
            quitImage, &ebiten.DrawImageOptions{
                GeoM: qgm})
        bw, bh := frontImage.Size()
        gm := ebiten.GeoM{}
        gm.Scale(0.5, 0.5)
        switch g.GetN() {
            case 2:
                gm.Translate(float64((w / 2) - 10 - (bw / 2)), float64((h / 2) - 10 - (bh / 2)))
            case 4:
                gm.Translate(float64((w / 2) - 30 - bw), float64((h / 2) - 30 - bh))
            case 6:
                gm.Translate(float64((w / 2) - 50 - ((3 * bw) / 2)), float64((h / 2) - 50 - ((3 * bh) / 2)))
        }
        for a := 0; a < g.GetN(); a++ {
            for b := 0; b < g.GetN(); b++ {
                screen.DrawImage(
                    cards[[2]int{a, b}], &ebiten.DrawImageOptions{
                        GeoM: gm})
                if b < g.GetN() - 1 {
                    gm.Translate(float64((bw / 2) + 20), float64(0))
                } else if a < g.GetN() - 1 {
                    gm.Translate(-float64(((bw / 2) + 20) * (g.GetN() - 1)), float64((bh / 2) + 20))
                }
            }
        }
        if sleep {
            if sleept == 0 {
                time.Sleep(dur)
                sleept = 1
                sleep = false
                cards[cardclicked] = frontImage
                cards[flippedcard] = frontImage
            } else {
                sleept--
            }
        }
        if !start {
            sgm := ebiten.GeoM{}
            sgm.Scale(0.25, 0.25)
            sgm.Translate(float64(float64(w) - (float64(w) / 4.5)), float64(float64(h) - (float64(h) / 3.835)))
            screen.DrawImage(
                startImage, &ebiten.DrawImageOptions{
                    GeoM: sgm})
        }
    }
    if start {
        if t0.IsZero() {
            t0 = time.Now()
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
        wgm := ebiten.GeoM{}
        wgm.Translate(float64((w / 2) - 122), float64((h / 2) - 37))
        wwi := ebiten.NewImage(244, 84)
        wwi.Fill(color.Black)
        screen.DrawImage(
            wwi, &ebiten.DrawImageOptions{
                GeoM: wgm})
        wgm.Translate(float64(2), float64(2))
        wi := ebiten.NewImage(240, 80)
        wi.Fill(color.RGBA{241, 162, 47, 255})
        screen.DrawImage(
            wi, &ebiten.DrawImageOptions{
                GeoM: wgm})
        fo := truetype.NewFace(fon, &truetype.Options{Size: 40})
        text.Draw(screen, "YOU WON", fo, (w / 2) - 100, (h / 2) + 20, color.RGBA{22, 154, 26, 204})
        dgm := ebiten.GeoM{}
        dgm.Translate(float64(w / 8), float64(h) / 1.5)
        cdi := ebiten.NewImage(344, 84)
        cdi.Fill(color.Black)
        screen.DrawImage(
            cdi, &ebiten.DrawImageOptions{
                GeoM: dgm})
        dgm.Translate(float64(2), float64(2))
        di := ebiten.NewImage(340, 80)
        di.Fill(color.RGBA{241, 162, 47, 255})
        screen.DrawImage(
            di, &ebiten.DrawImageOptions{
                GeoM: dgm})
        text.Draw(screen, "Change Difficulty", fo, (w / 8) + 20, int(float64(h) / 1.5) + 60, color.Black)
        //bos := text.BoundString(fo, "Change Difficulty")
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

func clearBG() {
    bgimage, _, err := image.Decode(bytes.NewReader(platformer.Background_png))
    if err != nil {
        log.Fatal(err)
    }
    bgImage = ebiten.NewImageFromImage(bgimage)
}

func main() {
    bgimage, _, err := image.Decode(bytes.NewReader(platformer.Background_png))
    if err != nil {
        log.Fatal(err)
    }
    bgImage = ebiten.NewImageFromImage(bgimage)

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

    startimage, _, err := image.Decode(bytes.NewReader(assets.StartButton_png))
    if err != nil {
        log.Fatal(err)
    }
    startImage = ebiten.NewImageFromImage(startimage)

    ebiten.SetFullscreen(true)
    ebiten.SetWindowTitle("Card Memory Game")

    game := &Game{n: 4}

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}

