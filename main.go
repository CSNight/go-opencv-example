package main

import (
	"gocv.io/x/gocv"
	"image"
	"math"
	"sync"
)

func main() {
	center := gocv.Point2f{
		X: 195, Y: 122,
	}
	angle := -24.5771
	src := gocv.IMRead("a.png", gocv.IMReadColor)
	var dest gocv.Mat = gocv.NewMat()
	rotateImage(src, &dest, -angle, center)
	gocv.IMWrite("b.png", dest)
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func rotateImage(src gocv.Mat, dst *gocv.Mat, angle float64, center gocv.Point2f) {
	if src.Empty() {
		return
	}
	translationMatrix := gocv.Zeros(2, 3, gocv.MatTypeCV32FC1)
	translationMatrix.SetFloatAt(0, 0, 1)
	translationMatrix.SetFloatAt(0, 2, float32(src.Cols())/2-center.X)
	translationMatrix.SetFloatAt(1, 1, 1)
	translationMatrix.SetFloatAt(1, 2, float32(src.Rows())/2-center.Y)

	var translationImage = gocv.NewMat()
	gocv.WarpAffine(src, &translationImage, translationMatrix, image.Point{
		X: src.Cols(),
		Y: src.Rows(),
	})
	//旋转
	alpha := -angle * 3.141592653 / 180.0

	var srcP = make([]gocv.Point2f, 3)
	var dstP = make([]gocv.Point2f, 3)
	srcP[0] = gocv.Point2f{X: 0, Y: float32(translationImage.Rows())}
	srcP[1] = gocv.Point2f{X: float32(translationImage.Cols()), Y: 0}
	srcP[2] = gocv.Point2f{X: float32(translationImage.Cols()), Y: float32(translationImage.Rows())}
	for i := 0; i < 3; i++ {
		dstP[i] = gocv.Point2f{
			X: srcP[i].X*float32(math.Cos(alpha)) - srcP[i].Y*float32(math.Sin(alpha)),
			Y: srcP[i].Y*float32(math.Cos(alpha)) + srcP[i].X*float32(math.Sin(alpha)),
		}
	}

	minx := math.Min(math.Min(math.Min(float64(dstP[0].X), float64(dstP[1].X)), float64(dstP[2].X)), 0)
	miny := math.Min(math.Min(math.Min(float64(dstP[0].Y), float64(dstP[1].Y)), float64(dstP[2].Y)), 0)
	maxx := math.Max(math.Max(math.Max(float64(dstP[0].X), float64(dstP[1].X)), float64(dstP[2].X)), 0)
	maxy := math.Max(math.Max(math.Max(float64(dstP[0].Y), float64(dstP[1].Y)), float64(dstP[2].Y)), 0)

	w := maxx - minx
	h := maxy - miny
	warpMat := gocv.Zeros(2, 3, gocv.MatTypeCV64F)
	warpMat.SetDoubleAt(0, 0, math.Cos(alpha))
	warpMat.SetDoubleAt(0, 1, 0-math.Sin(alpha))
	warpMat.SetDoubleAt(1, 0, math.Sin(alpha))
	warpMat.SetDoubleAt(1, 1, math.Cos(alpha))
	warpMat.SetDoubleAt(0, 2, 0-minx)
	warpMat.SetDoubleAt(1, 2, 0-miny)
	gocv.WarpAffine(translationImage, dst, warpMat, image.Point{
		X: int(w),
		Y: int(h),
	})
}
