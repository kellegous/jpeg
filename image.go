package jpeg

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type Image []Chunk

// Strip removes all segments of the JPEG where non-essential metadata is stored. This includes all APPn segments and any COM segment. This
// will also truncate the image at the first EOI segment. This seems to be similar in function to `exiftool -all=“
func (i Image) Strip() Image {
	var slim Image
	for _, chunk := range i {
		if seg, ok := chunk.(*Segment); ok {
			switch seg.Type() {
			case APP0, APP1, APP2, APP3, APP4, APP5, APP6, APP7, APP8, APP9, APP10, APP11, APP12, APP13, APP14, APP15, COM:
				continue
			case EOI:
				return append(slim, chunk)
			}
		}
		slim = append(slim, chunk)
	}
	return slim
}

// Write writes the image to a writer.
func (i Image) Write(w io.Writer) (int, error) {
	var t int
	for _, chunk := range i {
		n, err := w.Write(chunk.Data())
		t += n
		if err != nil {
			return t, err
		}
	}
	return t, nil
}

// Parse parses a JPEG image into chunks, either a segment or entropy-coded data.
func Parse(b []byte) (Image, error) {
	chunks := make([]Chunk, 0, 4)
	for len(b) > 0 {
		// we should begin with a segment marker
		if b[0] != 0xff {
			return nil, fmt.Errorf("expected segment marker, got %02x", b[0])
		} else if len(b) < 2 {
			return nil, errors.New("unexpected end of file")
		}

		x := SegmentType(b[1])
		if x == EOI {
			chunks = append(chunks, &Segment{data: b[:2]})
			b = b[2:]

			if len(b) == 0 {
				break
			}

			// Apple's Preview uses a strategy where tone maps are appended as a second complete
			// JPEG image following the initial EOI. Most parsers stop at the EOI, but that
			// strategy means you potentially miss the entire second image. So if we find that
			// there are bytes remaining, look ahead to see if we have an SOI. Otherwise, we return
			// it as detritus (trailing junk that would otherwise be ignored).
			if len(b) < 2 || (b[0] != 0xff && b[1] != byte(SOI)) {
				chunks = append(chunks, &Detritus{data: b})
				break
			}
		} else if x == SOI {
			chunks = append(chunks, &Segment{data: b[:2]})
			b = b[2:]
		} else if x >= RST0 && x <= RST7 {
			// RST segments are only allowed in the entropy-coded data.
			return nil, fmt.Errorf("unexpected RST%d segment", x&0xf)
		} else if len(b) < 4 {
			return nil, errors.New("unexpected end of file")
		} else {
			n := binary.BigEndian.Uint16(b[2:])
			chunks = append(chunks, &Segment{data: b[:n+2]})
			b = b[n+2:]

			if SegmentType(x) == SOS {
				n := findEntropyCodedDataLength(b)
				if n == -1 {
					return nil, errors.New("invalid entropy encoded data")
				}
				chunks = append(chunks, &EntropyCodedData{data: b[:n]})
				b = b[n:]
			}
		}
	}
	return Image(chunks), nil
}

// findEntropyCodedDataLength finds the length of the entropy-coded data by scanning
// the buffer until it finds a marker that is not either a byte shuffle or an RST segment.
func findEntropyCodedDataLength(b []byte) int {
	for i, n := 1, len(b); i < n; i++ {
		// possible marker
		if b[i-1] != 0xff {
			continue
		}

		// skip byte shufflling and RST segments
		if b[i] == 0x00 || (b[i] >= byte(RST0) && b[i] <= byte(RST7)) {
			continue
		}

		return i - 1
	}
	return -1
}
