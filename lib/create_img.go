package lib

import (
	"bufio"
	"encoding/base64"
	"image"
	"image/png"
	"io"
	"os"
	"crypto/md5"
)

func GetHash(params map[string]string) []byte {
	dataStr := ""

	for k, v := range params {
		dataStr += k + v
	}

	h := md5.New()
	h.Write([]byte(dataStr))
	return h.Sum(nil)
}

func bytesToInt(bytes []byte, x, y int) uint8 {
	length := (x ^ y) % len(bytes)
    var result uint8 = 0

	for i := length ; i < len(bytes) ; i++ {
		result = result | uint8(bytes[i])
	}

	return (result * uint8(x^y))
}

func Get2dHash(hash []byte, dx, dy int) [][]uint8 {
	s := make([][] uint8, dy)

	for y := range make([]int, dy)  {
		r := make([] uint8, dx)
		for x := range make([]int, dx)  {
			r[x] = bytesToInt(hash, x, y)
		}
		s[y] = r
	}

	return s
}

func GenerateIdenticon(params map[string]string, dx, dy int) image.Image {
	hash := GetHash(params)
	data := Get2dHash(hash, dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] =  uint8(v)
			m.Pix[i+1] = uint8(v)
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
	return m
}

func ShowImage(m image.Image) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	io.WriteString(w, "IMAGE:")
	b64 := base64.NewEncoder(base64.StdEncoding, w)
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(b64, m)
	if err != nil {
		panic(err)
	}
	b64.Close()
	io.WriteString(w, "\n")
}
