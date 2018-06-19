cp ../../db/schema.sql .
docker build -t ksunhokim/gorani-test-mysql .
docker push ksunhokim/gorani-test-mysql
