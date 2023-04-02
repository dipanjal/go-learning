skip_dirs=("array/")

for dir in */; do
    if ! [[ "${skip_dirs[@]}" =~ "$dir" ]]; then
        rm -rf "$dir"/main.go
    fi
done
