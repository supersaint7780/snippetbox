docker pull mysql:latest

docker run \
   --name snippetbox-mysql \
   -e MYSQL_ROOT_PASSWORD=9026985224 \
   -p 3306:3306 \
   -v snippetbox-mysql-data:/var/lib/mysql \
   -d mysql

docker exec -it snippetbox-mysql mysql -u root -p

docker stop snippetbox-mysql

docker start snippetbox-mysql