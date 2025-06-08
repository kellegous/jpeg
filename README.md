# A JPEG structure parser in Go

[JPEG](https://en.wikipedia.org/wiki/JPEG) is an image format that is mostly known for its lossy discrete cosine transform (DCT) compression. However, the internal structure of a JPEG file includes an array of different segments and a section of entropy-coded data. This parser only breaks a JPEG up into that structure. This is useful for extracting EXIF data or, as I am using it, for stripping out anything that is not the image data.
