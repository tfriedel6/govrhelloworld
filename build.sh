set -e

cd govrlib

export CGO_ENABLED=1
export CGO_LDFLAGS="-Wl,-soname,libgovrlib.so -llog"
export GOOS=android

export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang
export GOARCH=arm64

echo "Building 64 bit library"

go build -buildmode=c-shared -tags egl -o libgovrlib.so
mkdir -p arm64-v8a
mv libgovrlib.so arm64-v8a/libgovrlib.so

export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi21-clang
export GOARCH=arm
export GOARM=7

echo "Building 32 bit library"

go build -buildmode=c-shared -tags egl -o libgovrlib.so
mkdir -p armeabi-v7a
mv libgovrlib.so armeabi-v7a/libgovrlib.so

# Not sure why this sed is necessary, but it doesn't compile with this check included,
# so this comments it out
sed -i "s/typedef char _check_for_/\/\/typedef char _check_for_/" libgovrlib.h

mv libgovrlib.h govrlib.h

cd ../Projects/Android

python ./build.py
