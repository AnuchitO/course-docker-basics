package main

import (
	"errors"
	"net/url"
	"strings"

	"github.com/makiuchi-d/gozxing"
	zxqr "github.com/makiuchi-d/gozxing/qrcode"
	qrcode "github.com/skip2/go-qrcode"
)

const maxURLLen = 500

// validateURL accepts only absolute http(s) URLs.
func validateURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", errors.New("please enter a link")
	}
	if len(raw) > maxURLLen {
		return "", errors.New("that link is too long for a tidy QR code (max 500 characters)")
	}
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return "", errors.New("that does not look like a valid link")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return "", errors.New("only http:// and https:// links are allowed")
	}
	return u.String(), nil
}

// qrMatrix builds a QR code with the highest error correction (level H) and
// returns one string of '1'/'0' per row (no quiet zone). It decodes its own
// image first, so a code that cannot be read back is never returned.
func qrMatrix(link string) ([]string, error) {
	q, err := qrcode.New(link, qrcode.Highest)
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
