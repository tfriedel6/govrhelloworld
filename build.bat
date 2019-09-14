@echo off

cd govrlib

set CGO_ENABLED=1
set CGO_LDFLAGS=-Wl,-soname,libgovrlib.so -llog
set GOOS=android

set CC=%ANDROID_NDK_ROOT%\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-android21-clang
set GOARCH=arm64

echo "Building 64 bit library"

go build -buildmode=c-shared -tags egl -o libgovrlib.so
if %errorlevel% neq 0 (
    cd ..
    exit /b %errorlevel%
)
if not exist arm64-v8a mkdir arm64-v8a
move libgovrlib.so arm64-v8a\libgovrlib.so

set CC=%ANDROID_NDK_ROOT%\toolchains\llvm\prebuilt\windows-x86_64\bin\armv7a-linux-androideabi21-clang
set GOARCH=arm
set GOARM=7

echo "Building 32 bit library"

go build -buildmode=c-shared -tags egl -o libgovrlib.so
if %errorlevel% neq 0 (
    cd ..
    exit /b %errorlevel%
)
if not exist armeabi-v7a mkdir -p armeabi-v7a
move libgovrlib.so armeabi-v7a\libgovrlib.so

rem Not sure why this replacement is necessary, but it doesn't compile with this check included,
rem so this comments it out
powershell -Command "(gc libgovrlib.h) -replace 'typedef char _check_for', '// typedef char _check_for' | Out-File -encoding ASCII libgovrlib.h"

move libgovrlib.h govrlib.h

cd ..\Projects\Android

python build.py

cd ..\..

