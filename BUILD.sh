#!/usr/bin/env bash
package_name=json-files
OUTPUT_DIR=build
if [ -d "$OUTPUT_DIR" ]; then echo | rm -Rf $OUTPUT_DIR; fi
echo | mkdir -p $OUTPUT_DIR
platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/arm" "freebsd/amd64" "freebsd/386" "netbsd/386" "netbsd/amd64" "solaris/amd64")
for platform in "${platforms[@]}"
do
platform_split=(${platform//\// })
GOOS=${platform_split[0]}
GOARCH=${platform_split[1]}
output_name=$OUTPUT_DIR'/'$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
done
