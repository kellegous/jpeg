# A JPEG structure parser in Go

[JPEG](https://en.wikipedia.org/wiki/JPEG) is an image format that is mostly known for its lossy discrete cosine transform (DCT) compression. However, the internal structure of a JPEG file includes an array of different segments and a section of entropy-coded data. This parser only breaks a JPEG up into that structure. This is useful for extracting EXIF data or, as I am using it, for stripping out anything that is not the image data.

## Example

```go
b, err := os.ReadFile("JPG_Test.jpg")
if err != nil {
	log.Panic(err)
}

chunks, err := jpeg.Parse(b)
if err != nil {
	log.Panic(err)
}

for _, chunk := range chunks {
	switch t := chunk.(type) {
	case *jpeg.Segment:
		fmt.Printf("Segment Type=%s %d bytes\n", t.Type().Name(), t.Len())
	case *jpeg.EntropyCodedData:
		fmt.Printf("Entropy-Coded Data %d bytes\n", len(t.Data()))
	case *jpeg.Detritus:
		fmt.Printf("Detritus %d bytes\n", len(t.Data()))
	}
}
```
