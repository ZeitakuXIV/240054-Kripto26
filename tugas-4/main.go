package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func channel(img *image.NRGBA, k int) (off, c int) {
	b := img.Bounds()
	pix := k / 3
	return img.PixOffset(b.Min.X+pix%b.Dx(), b.Min.Y+pix/b.Dx()), k % 3
}

func encode(img *image.NRGBA, text string) error {
	var bits []byte
	for i, ch := range []byte(text) {
		for j := range 8 {
			bits = append(bits, ch>>(7-j)&1)
		}
		if i == len(text)-1 {
			bits = append(bits, 1)
		} else {
			bits = append(bits, 0)
		}
	}

	b := img.Bounds()
	kapasitas := (b.Dx() * b.Dy() * 3 / 9) * 9
	if len(bits) > kapasitas {
		return fmt.Errorf("pesan terlalu panjang: butuh %d bit, kapasitas %d bit", len(bits), kapasitas)
	}

	for g := range len(bits) / 9 {
		for j := range 9 {
			off, c := channel(img, g*9+j)
			if img.Pix[off+c]&1 != bits[g*9+j] {
				img.Pix[off+c] ^= 1
			}
		}
	}
	return nil
}

func decode(img *image.NRGBA) string {
	b := img.Bounds()
	kelompok := b.Dx() * b.Dy() * 3 / 9

	var sb strings.Builder
	for g := range kelompok {
		var kar byte
		for j := range 8 {
			off, c := channel(img, g*9+j)
			kar = kar<<1 | (img.Pix[off+c] & 1)
		}
		sb.WriteByte(kar)

		off, c := channel(img, g*9+8)
		if img.Pix[off+c]&1 == 1 {
			break
		}
	}
	return sb.String()
}

func read(sc *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	sc.Scan()
	return sc.Text()
}

func encodeFlow(sc *bufio.Scanner) {
	cover := read(sc, "Enter image name (with extension): ")
	src, err := os.Open(cover)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	b := img.Bounds()
	dst := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	for y := range b.Dy() {
		for x := range b.Dx() {
			dst.Set(x, y, img.At(b.Min.X+x, b.Min.Y+y))
		}
	}

	data := read(sc, "Enter data to be encoded: ")
	if data == "" {
		fmt.Println("Error: Data is empty")
		return
	}
	if err := encode(dst, data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	out := read(sc, "Enter the name of new image (with extension): ")
	if e := strings.ToLower(filepath.Ext(out)); e != ".png" {
		if e != "" {
			out = strings.TrimSuffix(out, filepath.Ext(out))
			fmt.Printf("Catatan: %s diganti ke .png — JPEG lossy dan akan menghapus bit LSB.\n", e)
		}
		out += ".png"
	}

	f, err := os.Create(out)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()
	if err := png.Encode(f, dst); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Encoded %d characters into %s\n", len(data), out)
}

func decodeFlow(sc *bufio.Scanner) {
	name := read(sc, "Enter image name (with extension): ")
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	b := src.Bounds()
	dst := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	for y := range b.Dy() {
		for x := range b.Dx() {
			dst.Set(x, y, src.At(b.Min.X+x, b.Min.Y+y))
		}
	}
	fmt.Println("Decoded Word: " + decode(dst))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println(":: Welcome to Steganography ::")
	fmt.Println("1. Encode")
	fmt.Println("2. Decode")

	switch read(sc, "Choose (1/2): ") {
	case "1":
		encodeFlow(sc)
	case "2":
		decodeFlow(sc)
	default:
		fmt.Println("Invalid choice, exiting.")
	}
}
