LazyPDF
=======

This began as a rollback fork of https://github.com/Nitro/lazypdf to return it
to use as a library.

[![](https://travis-ci.org/Nitro/lazypdf.svg?branch=master)](https://travis-ci.org/Nitro/lazypdf)

This is a rasterizing engine for PDF documents, built around MuPDF. It exports
a Go interface that allows the creation of a `Rasterizer` for each document you
wish to be able to rasterize to images.  It generated Go `image.Image`s and you
can then render these as PNG/JPEG/etc by using the Go stdlib functions for
doing that (`image/jpeg`, `image/png`).

Building
--------

You can simply run `./build` and the shell script will do all the right things
on either Linux or OSX.

Installing
----------

You cannot simply `go get` this library. It pulls in and links against C code
which is external to this repository for licensing reasons (the MuPDF library
is AGPL). Thus, in order to build this library, you need to run the `build`
shell script to pull down and compile MuPDF first.

Anything that you then want to link against this library will need to use the
correct Cgo configuration.  That should look like:

```go
// #cgo CFLAGS: -I. -I./mupdf-1.12.0-source/include -I./mupdf-1.12.0-source/include/mupdf -I./mupdf-1.12.0-source/thirdparty/openjpeg -I./mupdf-1.12.0-source/thirdparty/jbig2dec -I./mupdf-1.12.0-source/thirdparty/zlib -I./mupdf-1.12.0-source/thirdparty/jpeg -I./mupdf-1.12.0-source/thirdparty/freetype -g
// #cgo LDFLAGS: -L./mupdf-1.12.0-source/build/release -lmupdf -lmupdfthird -lm -ljbig2dec -lz -lfreetype -ljpeg -lcrypto -lpthread
// #include <faster_raster.h>
import "C"
```

`<relative path>` in this case should be the relative location of this library
to your project. If you are building a Go project under the `github.com/Nitro`
folder, this would usually be `../lazypdf`.
