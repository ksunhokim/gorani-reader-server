for d in ../../server/proto/*/; do  
    DIR="$(echo "$d" | rev | cut -d'/' -f2 | rev)";
    echo "$DIR compiling"
    protoc -I="$d" "$d$DIR.proto" --go_out=plugins=grpc:"$d"
done