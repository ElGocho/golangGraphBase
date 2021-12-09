## Correr el proyecto

Usar comando "docker-compose up -d"

---

## Programas necesarios

1. Docker

Opcionalmente

2. Golang
3. RabbitMQ
4. PostgreSQL

---
Es posible correr todo el proyecto solo con el docker, al lanzar el comando "docker-compose up" pero si se quiere prescindir del docker instalar el golang, rabbitMQ y postgreSQL. Para correr todo en local en vez de los contenedores de docker.

Es recomendable correr el proyecto con docker para aprovechar la configuración en el docker-compose.yml y toda la portabilidad que probee usar contenedores.

En el archivo .env.example esta un ejemplo de archivo de recursos el cual por defecto debe llamarse ".env". Puede renombrar el archivo ".env.example" como ".env"



