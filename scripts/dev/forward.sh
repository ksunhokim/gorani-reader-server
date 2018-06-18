kubectl port-forward svc/mysql 3306:3306 &
kubectl port-forward svc/minio-service 9000:9000 &
kubectl port-forward svc/redis-master 6379:6379 -n gorani &
