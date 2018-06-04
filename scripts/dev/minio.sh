sudo docker run -d -p 9000:9000 \
-e MINIO_ACCESS_KEY=admin \
-e MINIO_SECRET_KEY=password \
-v /mnt/data:/data \
-v /mnt/config:/root/.minio \
--restart=unless-stopped \
minio/minio server /data