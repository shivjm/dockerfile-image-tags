$tags = (git tag) -split "`n"
$base = 'ghcr.io/shivjm/dockerfile-image-tags'

foreach ($tag in $tags) {
    Write-Output "Tag: $tag"

    git checkout $tag

    if (!(Test-Path Dockerfile)) {
        Write-Output "Skipping $tag (no Dockerfile)"
        continue
    }

    docker build . -t "${base}:$tag" -t latest
}

git checkout main
