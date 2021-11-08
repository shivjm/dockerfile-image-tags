#!/usr/bin/env
TAGS=`git tag`
OS_LIST=("linux" "windows" "darwin")

for tag in ${TAGS//\n/ }; do
    [[ $tag =~ ^v[0-9]\.[0-9]\.[0-9]$ ]] || continue

    echo "\nProcessing $tag…\n"

    git checkout "$tag"

    rm -rf build

    for os in ${OS_LIST[@]}; do
        output_path=build/$PACKAGE-amd64-$os

        if [[ $os == "windows" ]]; then
            output_path="$output_path.exe"
        fi

        GOOS=$os go build -tags netgo -o $output_path .
    done

    pushd build && gh release upload --clobber "$tag" ./* && popd
done
