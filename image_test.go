package jpeg

import (
	"os"
	"path/filepath"
	"testing"
)

func expectSegmentOf(st SegmentType, n int) func(t *testing.T, chunk Chunk) {
	return func(t *testing.T, chunk Chunk) {
		if seg, ok := chunk.(*Segment); ok {
			if seg.Type() != st {
				t.Fatalf("expected segment of type %s, got %s", st.Name(), seg.Type().Name())
			}
			if seg.Len() != n {
				t.Fatalf("expected segment of length %d, got %d", n, seg.Len())
			}
		} else {
			t.Fatalf("expected segment of type %s, got %T", st.Name(), chunk)
		}
	}
}

func expectEntropyCodedDataOf(n int) func(t *testing.T, chunk Chunk) {
	return func(t *testing.T, chunk Chunk) {
		if ecd, ok := chunk.(*EntropyCodedData); ok {
			if len(ecd.Data()) != n {
				t.Fatalf("expected entropy-coded data of length %d, got %d", n, len(ecd.Data()))
			}
		} else {
			t.Fatalf("expected entropy-coded data, got %T", chunk)
		}
	}
}

func expectDetritusOf(n int) func(t *testing.T, chunk Chunk) {
	return func(t *testing.T, chunk Chunk) {
		if detritus, ok := chunk.(*Detritus); ok {
			if len(detritus.Data()) != n {
				t.Fatalf("expected detritus of length %d, got %d", n, len(detritus.Data()))
			}
		} else {
			t.Fatalf("expected detritus, got %T", chunk)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		Filename string
		Expected []func(t *testing.T, chunk Chunk)
	}{
		{
			"JPG_Test.jpg",
			[]func(t *testing.T, chunk Chunk){
				expectSegmentOf(SOI, 2),
				expectSegmentOf(APP2, 542),
				expectSegmentOf(DQT, 69),
				expectSegmentOf(DQT, 69),
				expectSegmentOf(SOF0, 19),
				expectSegmentOf(DHT, 31),
				expectSegmentOf(DHT, 92),
				expectSegmentOf(DHT, 30),
				expectSegmentOf(DHT, 76),
				expectSegmentOf(SOS, 14),
				expectEntropyCodedDataOf(182960),
				expectSegmentOf(EOI, 2),
			},
		},
		{
			"jpeg400jfif.jpg",
			[]func(t *testing.T, chunk Chunk){
				expectSegmentOf(SOI, 2),
				expectSegmentOf(APP0, 18),
				expectSegmentOf(DQT, 69),
				expectSegmentOf(SOF0, 13),
				expectSegmentOf(DRI, 6),
				expectSegmentOf(DHT, 212),
				expectSegmentOf(SOS, 10),
				expectEntropyCodedDataOf(44734),
				expectSegmentOf(EOI, 2),
			},
		},
		{
			"jpeg420exif.jpg",
			[]func(t *testing.T, chunk Chunk){
				expectSegmentOf(SOI, 2),
				expectSegmentOf(APP1, 8819),
				expectSegmentOf(DQT, 134),
				expectSegmentOf(DHT, 420),
				expectSegmentOf(SOF0, 19),
				expectSegmentOf(SOS, 14),
				expectEntropyCodedDataOf(759198),
				expectSegmentOf(EOI, 2),
			},
		},
		{
			"jpeg422jfif.jpg",
			[]func(t *testing.T, chunk Chunk){
				expectSegmentOf(SOI, 2),
				expectSegmentOf(APP0, 18),
				expectSegmentOf(DQT, 199),
				expectSegmentOf(DHT, 420),
				expectSegmentOf(SOF0, 19),
				expectSegmentOf(SOS, 14),
				expectEntropyCodedDataOf(1548392),
				expectSegmentOf(EOI, 2),
			},
		},
		{
			"jpeg444.jpg",
			[]func(t *testing.T, chunk Chunk){
				expectSegmentOf(SOI, 2),
				expectSegmentOf(APP0, 18),
				expectSegmentOf(APP1, 70),
				expectSegmentOf(APP2, 70),
				expectSegmentOf(APP3, 70),
				expectSegmentOf(APP4, 70),
				expectSegmentOf(APP5, 70),
				expectSegmentOf(APP6, 70),
				expectSegmentOf(APP7, 70),
				expectSegmentOf(APP8, 70),
				expectSegmentOf(APP9, 70),
				expectSegmentOf(APP10, 70),
				expectSegmentOf(APP11, 70),
				expectSegmentOf(APP12, 70),
				expectSegmentOf(APP13, 70),
				expectSegmentOf(APP14, 70),
				expectSegmentOf(APP15, 70),
				expectSegmentOf(SOF0, 19),
				expectSegmentOf(DQT, 69),
				expectSegmentOf(DQT, 69),
				expectSegmentOf(DQT, 69),
				expectSegmentOf(DHT, 420),
				expectSegmentOf(SOS, 14),
				expectEntropyCodedDataOf(3934),
				expectSegmentOf(EOI, 2),
				expectDetritusOf(1),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Filename, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("testdata", test.Filename))
			if err != nil {
				t.Fatalf("failed to read test file: %v", err)
			}

			img, err := Parse(data)
			if err != nil {
				t.Fatalf("failed to parse image: %v", err)
			}

			if len(img) != len(test.Expected) {
				t.Fatalf("expected %d chunks, got %d", len(test.Expected), len(img))
			}

			for i, expected := range test.Expected {
				expected(t, img[i])
			}
		})
	}
}
