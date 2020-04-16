#!/bin/bash
set -e

NAME=blackjack
VERSION=$(git describe --tags --always --long --dirty)
BASE=$(pwd)
RELEASE=$BASE/release/$VERSION

echo "Creating release of $NAME with version: $VERSION"

# build all binaries
make

# create release folder
rm -rf $RELEASE
mkdir -p $RELEASE

# copy model
cp ./machine/model.bin $RELEASE/model.bin

# package windows binary
echo Packaging Windows binary...
mv "$NAME"_windows_amd64.exe $RELEASE/$NAME.exe
cd $RELEASE
zip "$NAME"_windows.zip $NAME.exe model.bin
rm -rf $NAME.exe
cd $BASE

# package linux binary
echo Packaging Linux binary...
mv "$NAME"_linux_amd64 $RELEASE/$NAME
cd $RELEASE
tar -czvf "$NAME"_linux.tar.gz $NAME model.bin
rm -rf $NAME
cd $BASE

# package mac binary
echo Packaging macOS binary...
mv "$NAME"_darwin_amd64 $RELEASE/$NAME
cd $RELEASE
tar -czvf "$NAME"_mac.tar.gz $NAME model.bin
rm -rf $NAME
cd $BASE

# remove model
rm -rf $RELEASE/model.bin

echo "Release of $NAME created: $RELEASE"