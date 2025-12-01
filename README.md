# OishiDes (_oh-ee-shee des_)
A SPA that helps manage a boutique restaurant and book tables for customers, built use Go, TS, react and vite. 

## Development 

### Dependencies 
* Backend  
    * Go [1.20.5 or higher](https://go.dev/doc/)
    * fiber [http server](https://gofiber.io/)
    * gorm [db ORM](https://gorm.io/)  
* Frontend  
    * Typescript v5 or higher
    * pnpm [package manager](https://pnpm.io/)
    * vite [bundler/dev server](https://vitejs.dev/guide/why.html)
    * react
### Starting database for dev

Start the docker container and grab the container IP from the commands below 
```bash 
docker run -d --name mysql-dev-server  -e MYSQL_ROOT_PASSWORD=password123 bitnami/mysql:latest
docker inspect <container-ID>
```
### Backend ENV variables 
```bash
export MYSQL_USERNAME=<dbuser>
export MYSQL_PASSWORD=<dbpass>
export MYSQL_IP=<dbIP:port>
export APP_AUTH=<path to secrets folder conataining generated keys>
```

### Starting backend for dev
```bash
cd oishi-server
go mod tidy
go run cmd/main.go
```
### Frontend ENV variables
```bash 
export API_GATEWAY=<IP where backend is hosted >
export API_SERVER_HOST=<IP where backend is hosted while deploying in a container>
export DEPLOY_ENV=<DEV||PROD>
```
### Starting frontend for dev
```bash
cd oishi-ui 
pnpm install
pnpm dev
```

### Generating RSA 256 private and public keys
```bash
cd ./secret
openssl genrsa -out private_key.pem 2048
openssl rsa -in test_key.pem -outform PEM -pubout -out test_key.pem.pub 
```
