#!/usr/bin/env bash
 
package_name="retroarch-links-generator"

# go tool dist list
platforms=("linux/amd64" "darwin/amd64" "windows/amd64" "linux/386" "windows/386")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	
	GOOS=${platform_split[0]}
	
	GOARCH=${platform_split[1]}

	output_name_cli=$package_name'-'$GOOS'-'$GOARCH

	if [ $GOOS = "windows" ];
	then
		output_name_cli+='.exe'
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -o dist/$output_name_cli $package

	if [ $? -ne 0 ];
	then
		echo 'An error has occurred! Aborting the script execution...'

		exit 1
	fi
done
