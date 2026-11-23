package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/makiuchi-d/gozxing"
	zxqr "github.com/makiuchi-d/gozxing/qrcode"
	qrcode "github.com/skip2/go-qrcode"
)

const maxTextLen = 300

// validateURL accepts any non-empty text of up to 300 characters (links included).
func validateURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", errors.New("please enter a link or some text")
	}
	if n := utf8.RuneCountInString(raw); n > maxTextLen {
		return "", fmt.Errorf("that's %d characters. A garden fits up to %d", n, maxTextLen)
	}
	return raw, nil
}

var matrixCache sync.Map // text -> []string; a text always gives the same code

// cachedMatrix is qrMatrix with a memory: the gallery asks for the same codes again and again.
func cachedMatrix(text string) ([]string, error) {
	if v, ok := matrixCache.Load(text); ok {
		return v.([]string), nil
	}
	rows, err := qrMatrix(text)
	if err == nil {
		matrixCache.Store(text, rows)
	}
	return rows, err
}

// qrMatrix builds a QR code with the highest error correction (level Q) and
// returns one string of '1'/'0' per row (no quiet zone). It decodes its own
// image first, so a code that cannot be read back is never returned.
func qrMatrix(link string) ([]string, error) {
	q, err := qrcode.New(link, qrcode.High)
	if err != nil {
		return nil, err
	}

	// verify with a quiet zone, the way a phone camera sees it
	bmp, err := gozxing.NewBinaryBitmapFromImage(q.Image(512))
	if err != nil {
		return nil, err
	}
	res, err := zxqr.NewQRCodeReader().Decode(bmp, nil)
	if err != nil || res.GetText() != link {
		return nil, errors.New("verification failed: generated code could not be read back")
	}

	q.DisableBorder = true
	bm := q.Bitmap()
	rows := make([]string, len(bm))
	for y, row := range bm {
		var b strings.Builder
		for _, on := range row {
			if on {
				b.WriteByte('1')
			} else {
				b.WriteByte('0')
			}
		}
		rows[y] = b.String()
	}
	return rows, nil
}
