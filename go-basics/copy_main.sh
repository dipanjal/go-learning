skip_dirs=("app/" "array/" "utils/")
source_path=./array/main.go
for dir in */; do
    if ! [[ "${skip_dirs[@]}" =~ "$dir" ]]; then
        cp "${source_path}" "$dir"
    fi
done