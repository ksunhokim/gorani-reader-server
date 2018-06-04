PS3='enter: '
options=("recreate gorani_reader" "recreate gorani_reader_test" "recreate both" "quit")
select opt in "${options[@]}"
do
    case $opt in
        "recreate gorani_reader")
            mysql -u root -p < recreate.sql
            break
            ;;
        "recreate gorani_reader_test")
            mysql -u root -p < recreate_test.sql
            break
            ;;
        "recreate both")
            cat recreate.sql recreate_test.sql | mysql -u root -p
            break
            ;;
        "quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done