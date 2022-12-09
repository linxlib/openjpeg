# openjpeg ![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/linxlib/openjpeg?style=flat-square)
golang wrapper for openjpeg library

# description
only tested on `windows x64 (Windows 7 Windows 10)` `Mac OS X 13` `centos 7` `ubuntu 22.04`. you should test it yourself for other platforms.

# usage
refer to `example` dir
```go
import openjpeg "github.com/linxlib/openjpeg"
```

# build

install golang > 1.18 first

## Windows

- install [MSYS2](http://www.msys2.org/)
- add `C:\msys64\mingw64\bin` to `PATH` environment
- open `MSYS2 MSYS` from start menu, install gcc with `pacman -S mingw-w64-x86_64-gcc`
- `git clone https://github.com/linxlib/openjpeg`
- `cd openjpeg\example`
- `go run -v .\main.go`
- then you get a `jp2000.jpg` in the same dir. you can change the example code to get different image file


## Linux (Centos 7)
The version of openjpeg in package manager may not the latest (v2.5.0), so we need to build from source.

### build and install openjpeg 2.5.0 (optional)
- install gcc cmake ...etc
- `wget https://github.com/uclouvain/openjpeg/archive/refs/tags/v2.5.0.zip`
- `unzip v2.5.0.zip && cd openjpeg-2.5.0/`
- `mkdir build && cd build`
- `cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS:bool=on`
- `make`
- `make install`
- you will find openjpeg(libopenjp2) library and  in `/usr/local/lib` `/usr/local/include` ...

### run
- `git clone https://github.com/linxlib/openjpeg`
- `cd openjpeg\example`
- `go run main.go`
- then you get a `jp2000.jpg` in the same dir. you can change the example code to get different image file


## Mac OS X
May be you don't need to install openjpeg just like linux.

- configure cgo compile environment by yourself
- if you have installed ffmpeg, so openjpeg 2.5.0 already installed, just skip the next step
- `brew install openjpeg` and make sure `opj_decompress -h` can get something like `openjp2 library v2.5.0`
- `git clone https://github.com/linxlib/openjpeg`
- `cd openjpeg/example`
- `go run main.go`
- then you get a `jp2000.jpg` in the same dir. you can change the example code to get different image file


## TODO

- [ ] cross compile (compile executable file in only one platform (may be docker))
- [ ] add some other platforms support




