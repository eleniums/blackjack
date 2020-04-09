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

# copy model and predict.py
cp ./machine/model.bin $RELEASE/model.bin
cp ./machine/predict.py $RELEASE/predict.py

# package windows binary
echo Packaging Windows binary...
mv "$NAME"_windows_amd64.exe $RELEASE/$NAME.exe
cd $RELEASE
zip "$NAME"_windows.zip $NAME.exe model.bin predict.py
rm -rf $NAME.exe
cd $BASE

# package linux binary
echo Packaging Linux binary...
mv "$NAME"_linux_amd64 $RELEASE/$NAME
cd $RELEASE
tar -czvf "$NAME"_linux.tar.gz $NAME model.bin predict.py
rm -rf $NAME
cd $BASE

# package mac binary
echo Packaging macOS binary...
mv "$NAME"_darwin_amd64 $RELEASE/$NAME
cd $RELEASE
tar -czvf "$NAME"_mac.tar.gz $NAME model.bin predict.py
rm -rf $NAME
cd $BASE

# remove model and predict.py
rm -rf $RELEASE/model.bin
rm -rf $RELEASE/predict.py

echo "Release of $NAME created: $RELEASE"