#!/usr/bin/env bash
export GOPATH=$GOPATH:/Users/jemy/QiniuCloud/Projects/qlive
gox -os="windows" -arch="386"   -output="../bin/qlive_windows_386"
gox -os="windows" -arch="amd64" -output="../bin/qlive_windows_amd64"
gox -os="darwin"  -arch="386"   -output="../bin/qlive_darwin_386"
gox -os="darwin"  -arch="amd64" -output="../bin/qlive_darwin_amd64"
gox -os="linux"   -arch="386"   -output="../bin/qlive_linux_386"
gox -os="linux"   -arch="amd64" -output="../bin/qlive_linux_amd64"
gox -os="linux"   -arch="arm"   -output="../bin/qlive_linux_arm"