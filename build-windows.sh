#!/bin/bash

# 设置交叉编译环境
export GOOS=windows
export GOARCH=amd64
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++

# 清理并构建
wails build -clean -platform windows/amd64 -nsis

# 输出构建结果
echo "Build completed. Check the 'build/bin' directory for output files."