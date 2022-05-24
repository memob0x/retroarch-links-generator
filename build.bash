#!/usr/bin/env bash

package=github.com/memob0x/retroarch-links-generator

package_split=(${package//\// })
package_name=${package_split[-1]}

GOOS=windows

GOARCH=amd64

output_name=$package_name'-'$GOOS'-'$GOARCH 

env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name.exe $package

if [ $? -ne 0 ];
then
    echo 'An error has occurred! Aborting the script execution...'

    exit 1
fi 
