#!/bin/sh

set -e

destination_directory='build'

rm -rf $destination_directory
mkdir -p $destination_directory

for os in windows linux darwin; do
  for arch in amd64; do
    for source_file in private_server.go public_server.go; do
      if [ "$os" == "windows" ]; then
        ext='.exe'
      else
        ext=''
      fi

      output_path=$destination_directory/${source_file/.go}.$os.$arch$ext

      echo "Building application for os $os and arch $arch: $output_path" >&2
      GOOS=$os GOARCH=$arch go build -o $output_path $source_file

      if [ "$os" == "linux" ]; then
        goupx $output_path
      else
        upx $output_path
      fi
    done
  done
done
