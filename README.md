# CRUD-Alumno

Requisitos del sistema
4GB RAM
4 CPUS

#Como contruir y ejecutar el Sistsema

Para este CRUD de Alumnos utiizamos docker Desktop y Kubernetes (puede utilizar otro controlador de contendores, pero tenga en cuenta que debe adaptar el controlador de ingress y cert-manager segun el controlador)

Intalar Docker Desktop y hablitar kubernetes
https://www.docker.com/products/docker-desktop/
 
instalar controlador de Ingress (NGINX) para docker
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.4.0/deploy/static/provider/aws/deploy.yaml

instalar cert-manager 
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.10.1/cert-manager.yaml

------------------------------



#Clone el repositorio y ejecute estos comandos sobre la carpeta principa:

docker build -t nodejs/module-manager-grud ./Frontend-ReactJS/Module-Manager-Grud

docker build -t nodejs/login ./Frontend-ReactJS/Module-Login

docker build -t golang/module-manager-grud-graphql ./GraphQL-GRPC-Backend-Golang/Module-Manager-CRUD-GraphQL

docker build -t golang/module-auth-graphql ./GraphQL-GRPC-Backend-Golang/Module-Auth-GraphQL

docker build -t golang/module-encryption-grpc ./GraphQL-GRPC-Backend-Golang/Module-Encryption-gRPC

docker build -t golang/module-token-grpc ./GraphQL-GRPC-Backend-Golang/Module-Token-gRPC

docker build -t python/module-db-users-grpc ./gRPC-DB-Python/Module-Users-DB-gRPC



#Entre a la carpeta /Kubernetes y ejecute lo siguientes comandos:

kubectl apply -f ./namespaces

kubectl apply -f .

El sistema se ejecutara y contrura en automatico. Solo espere su ejeucion 

Al Finalizar en un navegador al entrar en https://localhost/ o https://127.0.0.1/ deberia mostrar el sistemas ejecutandose

El usuario principal e inicial es
  usurio: root
  Password: 12345678
