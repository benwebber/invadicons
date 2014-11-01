package invadicon

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/benwebber/bitboard"
	"github.com/disintegration/imaging"
)

type Invadicon struct {
	Hash       []byte
	Background color.RGBA
	Foreground color.RGBA
	Bitmap     uint64
}

// New creates a new invadicon from a seed string.
func New(s string) (*Invadicon, error) {
	// Calculate the MD5 hash of the input string.
	hash := md5.Sum([]byte(s))
	// Set the background and foreground to white and black, respectively,
	// for testing.
	bg := color.RGBA{0xff, 0xff, 0xff, 0xff}
	fg := color.RGBA{0x00, 0x00, 0x00, 0xff}
	bitmap, err := generateBitmap(hash[:4])
	if err != nil {
		return &Invadicon{}, err
	}
	return &Invadicon{
		Hash:       hash[:],
		Background: bg,
		Foreground: fg,
		Bitmap:     bitmap,
	}, nil
}

// Render the invadicon at a given size.
func (i *Invadicon) Render(w, h int) image.Image {
	// Draw the base image, including a 1px border.
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	// Fill in a uniform background.
	draw.Draw(img, img.Bounds(), &image.Uniform{i.Background}, image.ZP, draw.Src)
	// Fill in foreground pixels.
	for y := 1; y < 9; y++ {
		for x := 1; x < 9; x++ {
			p := (y-1)*8 + (x - 1)
			if bitboard.GetBit(&i.Bitmap, p) == 1 {
				img.Set(x, y, i.Foreground)
			}
		}
	}
	// Resize the image.
	m := imaging.Resize(img, w, h, imaging.NearestNeighbor)
	return m
}

// Write the invadicon to a data stream.
func (i *Invadicon) Write(out io.Writer, w, h int) {
	img := i.Render(w, h)
	png.Encode(out, img)
}

// Save the invadicon to the given filename.
func (i *Invadicon) Save(file string, w, h int) error {
	out, err := os.Create(file)
	defer out.Close()
	if err != nil {
		return err
	}
	i.Write(out, w, h)
	return nil
}

// PrettyPrint prints a Unicode representation of the invadicon for debugging
// purposes.
func (i *Invadicon) PrettyPrint() {
	const (
		square = 0x25a0 // black square
		space  = 0x00b7 // middle dot
	)
	// Print the bitboard with a 1px border.
	fmt.Println(strings.Repeat(string(space), 10))
	for y := 0; y < 8; y++ {
		fmt.Printf(string(space))
		for x := 0; x < 8; x++ {
			p := y*8 + x
			if bitboard.GetBit(&i.Bitmap, p) == 1 {
				fmt.Printf(string(square))
			} else {
				fmt.Printf(string(space))
			}
		}
		fmt.Printf(string(space))
		fmt.Println()
	}
	fmt.Println(strings.Repeat(string(space), 10))
}

// generateBitmap computes a symmetrical bitboard from a slice of 4 bytes.
func generateBitmap(b []byte) (uint64, error) {
	if len(b) < 4 {
		return 0, fmt.Errorf("generateBitmap: b must be at least 4 bytes in length")
	}
	// Although the bitboard consists of 64 unique positions, it is mirrored
	// about a vertical axis. This means we only need 4 bytes (64/2/8 = 4) to
	// construct the image.
	//
	// We will construct a bitboard mirrored about the centre horizonal ranks,
	// then rotate it for the final image.
	//
	// Store the first 4 bytes of the byte slice.
	var byteSlice []byte
	for _, n := range b[:4] {
		byteSlice = append(byteSlice, n)
	}
	// Append the same sequence in reverse.
	for i := 3; i >= 0; i-- {
		byteSlice = append(byteSlice, b[i])
	}
	// Read the byte slice into a 64-bit integer (the bitboard).
	var bitmap uint64
	buf := bytes.NewReader(byteSlice)
	err := binary.Read(buf, binary.LittleEndian, &bitmap)
	if err != nil {
		return 0, err
	}
	// Rotate the bitboard so it is mirrored horizontally.
	return bitboard.Rotate90(bitmap), nil
}
