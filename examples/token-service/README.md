# Token-Service Example
This example illustrates how Dockpit can help in development of an idomatic microservices. The token-service has the following behaviour:

1. It returns access tokens over HTTP to users who are able to authenticate using their username and password. 
2. In order to do so, the service will need to retrieve data from a *PostgreSQL* database.
4. The location of the database instance will need to be retrieved from the *etcd* service registry 
5. New accounts are created by another service in the system, upon creation the token-service will receive a new message on a *NSQ* message queue, after which a new token is generated for this user. 
3. The returned token also contains permission data, these permissions are fetched through the *HTTP API* of another microservice in the system.
