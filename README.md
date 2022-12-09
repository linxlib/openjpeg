# openjpeg
golang wrapper for openjpeg library

# usage
refer to `example` dir
```go
import openjpeg "github.com/linxlib/openjpeg"
```

# build

## Windows

- install [MSYS2](http://www.msys2.org/)
- add `C:\msys64\mingw64\bin` to `PATH` environment
- open `MSYS2 MSYS` from start menu, install gcc with `pacman -S mingw-w64-x86_64-gcc`
- `git clone https://github.com/linxlib/openjpeg`
- `cd openjpeg\example`
- `go run -v .\main.go`
- then you get a `jp2000.jpg` in the same dir. you can change the example code to get different image file

## Mac OS X
- configure cgo compile environment by yourself
- `brew install openjpeg` and make sure `opj_decompress -h` can get something like `openjp2 library v2.5.0`
- `git clone https://github.com/linxlib/openjpeg`
- `cd openjpeg\example`
- `go run -v .\main.go`
- then you get a `jp2000.jpg` in the same dir. you can change the example code to get different image file

## Linux (Centos 7)
- `wget https://github.com/uclouvain/openjpeg/archive/refs/tags/v2.5.0.zip`
- `unzip v2.5.0.zip && cd openjpeg-2.5.0/`
- 
