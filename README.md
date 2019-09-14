# Hello world Oculus Mobile project written in Go

Because it just wouldn't be right if the Oculus Go couldn't be programmed in Go :-)

This is just a hello world project to get anything running at all. It only displays a triangle and uses the tracked head position, but nothing else. 

The basic project is created from the VrTemplate project in the Oculus Mobile samples. See [this very helpful Gist](https://gist.github.com/cnnid/6819aba7de6044871c597cadb363894e) for how to do that.

The `Src` directory contains C code, mostly from the VrTemplate project, but with a few hooks where some external functions are called. These will be run in the Go library.

The `govrlib` directory contains the go sources and the binaries built from them. At the moment it contains just enough GL code to load a shader and display a triangle with it.

The `govrlib\gl` directory contains a slightly modified version of the `github.com/go-gl/gl/v3.1/gles2` package. The only change is some LDFLAGS for the Android build, and no pkg-config for building under Windows.

Currently this sample can be built under Linux by running the `build.sh` file in the root of the directory. The Android NDK has to be installed for this to work, and currently it uses version 21. It should be no problem to use a different NDK version by adjusting the paths in the `build.sh`.

Compiling for Windows should work with the `build.bat` file, but the same path adjustments need to be made.

Some environment variables need to be set:

OCULUS_SDK_PATH = The root path of the Oculus Mobile SDK  
ANDROID_NDK_HOME = The root path of the Android NDK  

Python needs to be installed for the Oculus build.
