<h1 align="center"> StarWars </h1>

<p align="left">
   <img src="https://img.shields.io/badge/STATUS-EN%20DESAROLLO-green">
   </p>

   ##Índice

   *[Descripción de la aplicación](#descripción-de-la-aplicacion)
   
   *[Tecnologías utilizadas](#tecnologías-utilizadas)
   
   *[Puesta en Marcha](#puesta_en_marcha)

   <h2>Descripción de la aplicación</h2>

   Servicio de Rastreo y Descifrado de Mensajes (Star Wars)

La aplicación es un servicio de rastreo y descifrado de mensajes inspirado en el universo de Star Wars. Su objetivo principal es recibir datos de distancia y pedazos de mensajes de diferentes satélites y determinar la posición y el mensaje original utilizando algoritmos de trilateración.

Características principales:

  1. Ruta /topsecret/: Permite enviar datos de distancia y mensajes de los satélites en una sola solicitud POST. La aplicación calculará la posición y descifrará el mensaje original utilizando la técnica de trilateración. Devuelve la posición y el mensaje descifrado.

  2. Ruta /topsecret_split/{satellite_name}: Permite enviar datos de distancia y mensaje para un satélite específico en una solicitud POST. Los datos se guardan en una base de datos PostgreSQL para su posterior procesamiento. Los datos se almacenan por satélite, lo que permite enviar la información en varias solicitudes POST. Cuando se tienen datos suficientes de los tres satélites, se calcula la posición y se descifra el mensaje original.

  3. Ruta /topsecret_split/: Permite realizar una solicitud GET para obtener la posición y el mensaje descifrado si se tienen suficientes datos de los tres satélites almacenados en la base de datos. Si no hay suficiente información, se devuelve un mensaje de error indicando que no hay suficiente información para determinar la posición y el mensaje.


<h2>Tecnologías utilizadas</h2>

1. Base de datos PostgreSQL: La aplicación utiliza PostgreSQL como base de datos para almacenar los datos de distancia y mensajes de los satélites. Se utiliza el ORM GORM para interactuar con la base de datos y se proporciona soporte para migración de la base de datos.

2. Código modularizado: El código se organiza en diferentes archivos y carpetas para una mayor modularidad y facilidad de mantenimiento. El archivo main.go se encuentra en la raíz del proyecto y se encarga de configurar y ejecutar el servidor HTTP. Los controladores, como topsecret.go y topsecret_split.go, se encuentran en la carpeta "controllers" y manejan las solicitudes y respuestas de cada ruta. El archivo models.go en la carpeta "models" define las estructuras de datos utilizadas en la aplicación.

3. Implementación de API con Gin: La aplicación utiliza el framework web Gin para implementar la API RESTful. Gin proporciona una sintaxis sencilla y un rendimiento eficiente para manejar las rutas, los middleware y las respuestas JSON.


<h2>Puesta en Marcha</h2>

Para poner en marcha el proyecto, sigue los siguientes pasos:

 1. Clona el repositorio desde el repositorio privado de GitHub.
```
git clone <URL del repositorio>

```
2. Asegúrate de tener Go instalado en tu máquina. Puedes descargarlo desde el sitio web oficial de Go: https://golang.org/dl/
3. Configura la base de datos PostgreSQL:
     a. Instala PostgreSQL en tu máquina si aún no lo has hecho. Puedes descargarlo desde el sitio web oficial de PostgreSQL: https://www.postgresql.org/download/
     b. Crea una base de datos en PostgreSQL con el nombre "starwars".
     c. Abre el archivo db/database.go y actualiza los valores de user y password en la cadena de conexión DSN de ConnectDB() con tus propias credenciales de PostgreSQL.
4. Abre una terminal y navega hasta el directorio raíz del proyecto.
   ````
   cd <ruta-del-proyecto>

   ```
5. Instala las dependencias del proyecto.
   ```
   go mod init guerra_galaxias
   go get -u github.com/gin-gonic/gin
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
   ```

6. Inicia el servidor.
   ```
   go run main.go
   ```
   El servidor se iniciará en el puerto 8080.

   Ahora puedes probar las siguientes rutas:

a. POST /topsecret/: Utiliza esta ruta para enviar un JSON con los datos de los satélites y obtener la posición y el mensaje descifrado. Asegúrate de enviar un JSON válido en el cuerpo de la solicitud.

b. POST /topsecret_split/{satellite_name}: Utiliza esta ruta para enviar los datos de un satélite específico. Reemplaza {satellite_name} con el nombre del satélite correspondiente (por ejemplo, "kenobi", "skywalker", "sato"). Asegúrate de enviar un JSON válido en el cuerpo de la solicitud.

c. GET /topsecret_split/: Utiliza esta ruta para obtener la posición y el mensaje descifrado si hay suficiente información de los tres satélites. Si no hay suficiente información, recibirás un mensaje de error indicando que no hay suficiente información.
   
