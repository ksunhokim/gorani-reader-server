source ./docker-push.sh

PS3='enter: '
options=("push base" "push etl" "push api" "quit")
select opt in "${options[@]}"
do
    case $opt in
        "push base")
			docker build -t gorani-base -f ../../go/Dockerfile ../../go
            break
            ;;
        "push etl")
			push $ETLREPO ../../go/Dockerfile.etl ../../go
            break
            ;;
        "push api")
			push $APIREPO ../../go/Dockerfile.api ../../go
            break
            ;;
        "quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done
