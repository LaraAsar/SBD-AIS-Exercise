We tried to run a multi-node Docker Swarm cluster, but Windows + Docker Desktop networking caused many issues for us, including duplicated hostnames (docker-desktop), routing issues and VM's not working. Even with multiple attempts (WSL reset, reinstall, network reconfiguration), it was not possible to maintain stable inter-node connectivity.

Because of this, as everything else failed, we decided to just run swarm manager alone locally. Therefore, the final setup runs on one node with Swarm Mode enabled, but all services are deployed exactly as required: Traefik, Orderservice, SWS, Postgres, Minio.

The full command list:

docker swarm init
docker secret create postgres_user docker/postgres_user_secret
docker secret create postgres_password docker/postgres_password_secret
docker secret create s3_user docker/s3_user_secret
docker secret create s3_password docker/s3_password_secret

docker stack deploy -c docker-compose.yml order
docker stack services order
docker stack ps order
docker service logs order_orderservice