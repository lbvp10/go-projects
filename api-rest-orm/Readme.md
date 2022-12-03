# **API REST Golang + ORM gorm + Postgresql**

<HR>

## **Run App**

`docker-compose -f ./api-rest-orm/docker/app.yml up `

**cURLs**
----------

* **POST** :<br><br>`curl --location --request POST 'localhost:23000/api' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "fecha":"2022-01-01T23:59:59Z",
  "descripcion":"prueba descripcion"
  }'`<br>
* **GET ALL** : <br><br>`curl --location --request GET 'localhost:23000/api' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "numero_dias":10,
  "fecha_inicial":"2022-01-01T23:59:59Z",
  "fecha_final":"2022-12-31T00:00:00Z"
  }'`<br>

* **GET by ID**: <br><br>`curl --location --request GET 'localhost:23000/api/1' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "numero_dias":10,
  "fecha_inicial":"2022-01-01T23:59:59Z",
  "fecha_final":"2022-12-31T00:00:00Z"
  }'`<br>
* **DELETE**:<br><br>`curl --location --request DELETE 'localhost:23000/api/1' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "numero_dias":10,
  "fecha_inicial":"2022-01-01T23:59:59Z",
  "fecha_final":"2022-12-31T00:00:00Z"
  }'`<br>