package imageProcessor

import (
    "os"
    "path/filepath"
    "image"
    "math"
    _ "image/jpeg"
//    "fmt"
)

const GRIDS int = 7

func mean(l []float64) float64 {
    sum := 0.0
    for _,num := range l {
        sum = sum + num
    }
    return sum/float64(len(l))
}

func variance(l []float64, mean float64) float64 {
    sum := 0.0
    for _,num := range l {
        sum = sum + math.Pow(num - mean,2)
    }
    return sum/float64(len(l))
}

func stdev(l []float64, mean float64) float64 {
    v := variance(l,mean)
    return math.Sqrt(v)
}

func to_f(in []uint32) []float64 {
    out := make([]float64,len(in))
    for i,val := range in {
        out[i] = float64(val)
    }
    return out
}

func vals(l []uint32) []float64 {
    lf := to_f(l)
    m := mean(lf)
    sd := stdev(lf,m)
    _ = sd
    return []float64{m,sd}
}

func calculate_values(b image.Rectangle, i image.Image) []float64 {
    var reds = make([]uint32,b.Dx()*b.Dy())
    var greens = make([]uint32,b.Dx()*b.Dy())
    var blues = make([]uint32,b.Dx()*b.Dy())
    index := 0
    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            c := i.At(x,y)
            red, green, blue, _ := c.RGBA()
            reds[index] = red
            greens[index] = green
            blues[index] = blue
            index = index + 1
        }
    }
    r := make([]float64,0)
    r = append(r,vals(reds)...)
    r = append(r,vals(greens)...)
    r = append(r,vals(blues)...)
    return r
}

//This is used in cases where the image size isn't divisible by GRIDS 
func Btoi(b bool) int {
    if b {
        return 1
    }
    return 0
 }

//Breaks the image into sections.  Finds the RGB averages and standard devaitions 
//for every section making 6 values per section.  Presently, there are 49 sections, 
//making a total of 49*6 = 294
func Process(filename string) []float64 {
    fImg, _ := os.Open(filename)
    defer fImg.Close()
    img, _, err := image.Decode(fImg)

	if err != nil {
		return nil
	}

    results := make([]float64,0)

    b := img.Bounds()

    yMin := b.Min.Y
    for i := 0; i < GRIDS; i++ {
        xMin := b.Min.X
        yMax := yMin + b.Dy()/GRIDS + Btoi(b.Dy() % GRIDS > i)
        for j := 0; j < GRIDS; j++ {
            xMax := xMin + b.Dx()/GRIDS+ Btoi(b.Dx() % GRIDS > j)
            vals :=calculate_values(image.Rect(xMin,yMin,xMax,yMax),img)
            results = append(results,vals...)
            xMin = xMax
        }
        yMin = yMax
    }

    return results
}   

func ProcessAsync(filename string, res chan []float64) {
    res <- Process(filename)
}

func ProcessDirectory(directory string) [][]float64 {
    results := make([][]float64,0)
    files, _ := filepath.Glob(directory)
    ch := make(chan []float64, len(files))
    for _, f := range files {
        go ProcessAsync(f,ch)
    }
    for _,_ = range files {
        fRes := <-ch
        if len(fRes) != 0 {
            results = append(results,fRes)
        }
    }
    return results
}

//This is what calls to ProcessDirectoy should look like
//func main() {
//    in := ProcessDirectory("./in/*.jpg")
//    out := ProcessDirectory("./out/*.jpg")
//}
