package core

import (
	"github.com/Comdex/imgo"
	"i-go/image/util"
	"image"
	"io"
)

// AdjustSaturation 调整图像饱和度 percent 饱和度控制参数 (0,1)
func AdjustSaturation(file io.Reader, percent float64) (saturation []byte) {
	// 1.解析图片加载矩阵
	_image, _, err := image.Decode(file)
	if err != nil {
		return
	}
	imgMatrix, err := imgo.Read(_image)
	if err != nil {
		return
	}
	// 2.饱和度调整
	rgb2Gray := doAdjustSaturation(imgMatrix, percent)
	// 3.矩阵转为[]byte
	saturation, err = util.Matrix2Bytes(rgb2Gray)
	if err != nil {
		return
	}
	return
}

// doAdjustSaturation 调整饱和度
func doAdjustSaturation(src [][][]uint8, percent float64) [][][]uint8 {
	var (
		height = len(src)
		width  = len(src[0])
	)
	imgMatrix := imgo.NewRGBAMatrix(height, width)
	copy(imgMatrix, src)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var (
				R = imgMatrix[i][j][0]
				G = imgMatrix[i][j][1]
				B = imgMatrix[i][j][2]
			)

			nr, ng, nb := adjustSaturationByRGB(R, G, B, percent)

			imgMatrix[i][j][0] = nr
			imgMatrix[i][j][1] = ng
			imgMatrix[i][j][2] = nb
		}
	}
	return imgMatrix
}

// adjustSaturationByRGB 通过 RGB 值调整饱和度
/*
本转换主要是在YCrCb格式下完成
1.所以首先第一步就是需要将图像从RGB转换到YCrCb格式。
2.之后将YCrCb转换回RGB格式。
3.具体调整方式如下
	1.饱和度调整：当图像格式转换到YCrCb之后，直接对Cr、Cb分量乘上权重值，通过分别的权重值调整，便分别控制图像红色部分和蓝色分的饱和度。
具体公式如下：
     Cr = Cr * Wr
     Cb = Cb * Wb
   2、色偏调整：在YCrCb格式转换回RGB格式时候，在R和B分量计算中加入控制权重，即可以控制图像红色部分和蓝色部分的色偏。
具体公式如下：
     R = Y + 1.371 * Cr * Wr
     B = Y + 1.732 * Cb * Wb
参考链接 https://blog.csdn.net/u011630458/article/details/51782321
*/
func adjustSaturationByRGB(R, G, B uint8, percent float64) (nr, ng, nb uint8) {
	Y, Cr, Cb := RGB2YCrCb(float64(R), float64(G), float64(B))
	// 饱和度调整
	Cr = Cr * percent
	Cb = Cb * percent
	nr, ng, nb = YCrCb2RGB(Y, Cr, Cb)
	return
}

// RGB2YCrCb 颜色空间转换
func RGB2YCrCb(R, G, B float64) (Y, Cr, Cb float64) {
	Y = 0.299*R + 0.587*G + 0.114*B
	Cr = 0.500*R - 0.419*G - 0.081*B
	Cb = -0.169*R - 0.331*G + 0.500*B
	return
}

// YCrCb2RGB
func YCrCb2RGB(Y, Cr, Cb float64) (R, G, B uint8) {
	R = uint8(Y + 1.371*Cr)
	G = uint8(Y - 0.6982*Cr - 0.3365*Cb)
	B = uint8(Y + 1.732*Cb)
	return
}
