/**
 * Auth :   liubo
 * Date :   2018/10/12 9:29
 * Comment: 
 */

package main

import (
	_ "bufio"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"os"
)

func main() {
	dst := imaging.New(720, 1080, color.NRGBA{0, 0, 0, 0})

	src, err := imaging.Open("bg.jpg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dst = imaging.Paste(dst, src, image.Pt(0, 0))

	file ,err := os.OpenFile("erweima.png", os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		log.Println("打开文件失败：", err.Error())
	}
	//temp := bufio.NewReader(buff)
	// raw, err := imaging.Open("erweima.png")
	raw, err := imaging.Decode(file)
	//raw := imaging.New(370, 370, color.NRGBA{255,0,0,255})
	if err == nil {
		dst = imaging.Paste(dst, raw, image.Pt(178, 841))
	} else {
		log.Println("解码失败", err.Error())
	}

	imaging.Save(dst, "2.png")
}