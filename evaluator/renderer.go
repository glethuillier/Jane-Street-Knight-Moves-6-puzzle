package evaluator

import (
	"fmt"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"image"
	"image/color"
	"image/png"
	"jane/board"
	"jane/paths"
	"os"
)

type segment struct {
	start, end board.Coordinates
}

// SaveAsPNG saves the solution as a PNG file
func SaveAsPNG(cb *board.ChessBoard, values *Values, path0, path1 *paths.Path) error {
	const (
		cellSize = 100
		padding  = cellSize / 2
	)

	boardWidthHeight := cb.Size * cellSize
	canvasWidthHeight := boardWidthHeight + padding*2
	img := image.NewRGBA(image.Rect(0, 0, canvasWidthHeight, canvasWidthHeight))
	dc := gg.NewContextForRGBA(img)

	// fill background
	dc.SetRGB255(255, 255, 255)
	dc.Clear()

	colors := getColorMappings()
	face := basicfont.Face7x13

	for c, square := range cb.Squares {
		startX, startY := computeStartCoordinates(c, cb.Size, cellSize, padding)
		boardColor := colors[square]
		drawSquare(dc, startX, startY, cellSize, boardColor, color.Black)

		value := getSquareValue(square, values)
		valueStr := fmt.Sprintf("%d", value)
		addLabel(dc, startX, startY, cellSize, valueStr, color.Black, face)
	}

	pathSegments := make(map[segment]int)
	drawPath(dc, path0, color.RGBA{R: 0, G: 0, B: 255, A: 255}, cb.Size, cellSize, padding, pathSegments)
	drawPath(dc, path1, color.RGBA{R: 255, G: 0, B: 0, A: 255}, cb.Size, cellSize, padding, pathSegments)

	for seg, count := range pathSegments {
		if count > 1 {
			drawLine(dc, seg.start, seg.end, color.RGBA{R: 0, G: 0, B: 0, A: 255}, cb.Size, cellSize, padding)
		}
	}

	for i := 0; i < cb.Size; i++ {
		letter := fmt.Sprintf("%c", 'a'+i)
		addLabel(dc, i*cellSize+cellSize/2, cb.Size*cellSize+cellSize/4, cellSize, letter, color.Black, face)
		number := fmt.Sprintf("%d", cb.Size-i)
		addLabel(dc, -cellSize/4, (i+1)*cellSize-cellSize/2, cellSize, number, color.Black, face)
	}

	filename := fmt.Sprintf(
		"solutions/%d-%d.png",
		values.A+values.B+values.C,
		len(path0.Squares)+len(path1.Squares),
	)

	return saveImage(img, filename)
}

// drawSquare draws a rectangle and border representing the square.
func drawSquare(dc *gg.Context, startX, startY, cellSize int, fillColor, borderColor color.Color) {
	dc.SetColor(fillColor)
	dc.DrawRectangle(float64(startX), float64(startY), float64(cellSize), float64(cellSize))
	dc.Fill()
	dc.SetColor(borderColor)
	dc.DrawRectangle(float64(startX), float64(startY), float64(cellSize), float64(cellSize))
	dc.Stroke()
}

// addLabel adds a text label to the image at the specified coordinates
func addLabel(dc *gg.Context, x, y, cellSize int, label string, col color.Color, face font.Face) {
	dc.SetColor(col)
	dc.SetFontFace(face)
	dc.DrawStringAnchored(label, float64(x+cellSize/2), float64(y+cellSize/2), 0.5, 0.5)
}

// drawLine draws a line between two coordinates
func drawLine(dc *gg.Context, start, end board.Coordinates, col color.RGBA, boardSize, cellSize, padding int) {
	startX, startY := computeCoordinates(start, boardSize, cellSize, padding)
	endX, endY := computeCoordinates(end, boardSize, cellSize, padding)
	dc.SetRGBA255(int(col.R), int(col.G), int(col.B), int(col.A))
	dc.SetLineWidth(3)
	dc.DrawLine(float64(startX), float64(startY), float64(endX), float64(endY))
	dc.Stroke()
}

// computeStartCoordinates calculates the top-left pixel coordinates for a given board coordinate
func computeStartCoordinates(c board.Coordinates, boardSize, cellSize, padding int) (int, int) {
	return int(c.File-'a')*cellSize + padding, (boardSize-c.Rank)*cellSize + padding
}

// computeCoordinates calculates the pixel coordinates for a given board coordinate
func computeCoordinates(c board.Coordinates, boardSize, cellSize, padding int) (int, int) {
	return int(c.File-'a')*cellSize + cellSize/2 + padding, (boardSize-c.Rank)*cellSize + cellSize/2 + padding
}

// makeSegment creates a segment from two coordinates
func makeSegment(start, end board.Coordinates) segment {
	if start.File > end.File || (start.File == end.File && start.Rank > end.Rank) {
		start, end = end, start
	}
	return segment{start, end}
}

// drawPath draws a single path on the chess board
func drawPath(dc *gg.Context, path *paths.Path, col color.RGBA, boardSize, cellSize, padding int, pathSegments map[segment]int) {
	for i := 0; i < len(path.Squares)-1; i++ {
		seg := makeSegment(path.Squares[i], path.Squares[i+1])
		pathSegments[seg]++
		drawLine(dc, seg.start, seg.end, col, boardSize, cellSize, padding)
	}
}

// saveImage saves the image as a PNG file
func saveImage(img *image.RGBA, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	return png.Encode(file, img)
}

// getColorMappings provides the color mappings for different squares
func getColorMappings() map[board.Square]color.RGBA {
	return map[board.Square]color.RGBA{
		"A": {R: 237, G: 237, B: 237, A: 255}, // light grey for A
		"B": {R: 221, G: 235, B: 247, A: 255}, // light blue for B
		"C": {R: 226, G: 239, B: 218, A: 255}, // light green for C
	}
}
