package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)



var AsciiSet = " .:co#%?@PO"
//var AsciiSet = "PO@?%#co:. "
var count uint32 = 0
var ScalingFac = 16


func luminous_intensity(img image.Image,x int,y int) (float64,color.RGBA) {
	var redsum uint32 = 0.0
	var greensum uint32 =0.0
	var bluesum uint32 =0.0
	var col color.RGBA

	for i:=0;i<=(y+ScalingFac);i++{
		for j:=0;j<=(x+ScalingFac);j++{
			col := img.At(x,y)
			r,g,b,_ := col.RGBA()
			 
			r = r >> 8
			g	= g >> 8
			b = b >> 8

		//	fmt.Printf("r,g,b ---> %v  , %v ,%v \n",r,g,b)
			redsum = redsum+r
			bluesum = bluesum+b
			greensum = greensum+g
			count++
		}
	}

	redsum 	 = redsum   /	count 
	greensum = greensum / count
	bluesum  = bluesum  / count

	//fmt.Printf("R %f \n",redSum)

	Y := (0.2126*float64(redsum)) + (0.7152*float64(greensum)) + (0.0722*float64(bluesum))
	col.R = uint8(redsum)
	col.G = uint8(greensum)
	col.B = uint8(greensum)
	col.A = 0
	
	//fmt.Printf("Y %f %v \n",Y,count)
	count = 0
	return Y,col
}

func main() {
	 // red := "\033[31m"
   // green := "\033[32m"
   // yellow := "\033[33m"
   // blue := "\033[34m"
   // reset := "\033[0m"

	fmt.Print("\x1b[38;5;196m [USAGE] ./program <filepath> \x1b[0m \n")
	filepathname:=os.Args[1]	

	fmt.Printf("Image name/Path --> %s \n",string(filepathname))
	
	File,err:= os.Open(filepathname)	
	if err != nil {
		log.Fatal("[USAGE] ./program <filepath>")
	}
	defer File.Close()

	ImgFile , err := png.Decode(File)
	if (err!=nil){
		fmt.Print("[ERROR] : File couldn't be Loaded \n",err)
	}

	ImageWidth:=ImgFile.Bounds().Max.X
	Imageheight:= ImgFile.Bounds().Max.Y 

	fmt.Printf("Image Dimensions : %v x %v \n",Imageheight,ImageWidth)

	for y:=0;y<=Imageheight;y++ {
		for x:=0;x<=ImageWidth;x++{
			Y,col:=luminous_intensity(ImgFile,x,y)
			
			index := int(Y*10/256.0)

			a :=AsciiSet[index]
			
			r,g,b,_:= col.RGBA()
			r = r >> 8
			g	= g >> 8 
			b = b >> 8
		
			//  fmt.Printf("\n %v %v %v  \n", r, g, b)  \x1b[48;5;16mHello \x1b[0m
				fmt.Printf("\x1b[48;5;16m\x1b[38;2;%v;%v;%vm %s \x1b[0m\x1b[0m",r,g,b,string(a) )			

			Y = 0.0
		}
		fmt.Print("\n")
	}

}








