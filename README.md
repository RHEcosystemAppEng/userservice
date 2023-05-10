# UserService

## Run Locally

* Run local instance of Keycloak on port 8080
```shell
    docker run -p 8080:8080 --env KEYCLOAK_ADMIN=admin --env KEYCLOAK_ADMIN_PASSWORD=admin quay.io/keycloak/keycloak start-dev
```

* Option 1: Build and run the userservice executable, which is exposed on port 8000.
```shell
    make build
    ./out/bin/userservice
```

* Option 2: Run userservice with the go run command, which is exposed on port 8000.
```shell
    go run .
```

* Get token from the userservice
```shell
curl --location 'http://localhost:8000/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'username=admin' \
--data-urlencode 'password=admin' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'client_id=admin-cli'
```

* A successful response means that the local Keycloak server and the userservice are working as expected.

## Using FindUser API
### Setup local Keycloak server instance with following test data
* Get master realm token from Keycloak server
```shell
curl --location 'http://localhost:8080/realms/master/protocol/openid-connect/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'username=admin' \
--data-urlencode 'password=admin' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'client_id=admin-cli'
```

* Create Engineers group
```shell
curl --location 'http://localhost:8080/admin/realms/master/groups' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <TOKEN>' \
--data '{
    "name": "fr1-eng"
}'
```

* Create Managers group
```shell
curl --location 'http://localhost:8080/admin/realms/master/groups' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <TOKEN>' \
--data '{
    "name": "fr1-mgr"
}'
```

* Create an Engineer User
```shell
curl --location 'http://localhost:8080/admin/realms/master/users' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <TOKEN>' \
--data-raw '{
    "username": "eng1-test",
    "groups" : ["fr1-eng"],
    "enabled" : "true",
    "email" : "1@1.com",
    "firstName": "fn",
    "lastName": "ln",
    "attributes": { "org_id" : "rh", "is_internal" : "true", "org_admin": "true", "type" : "engineer"},
    "credentials" : [
        { "credentialData" : "fr1-eng", "temporary" : "false", "type" : "password", "value" : "fr1-eng" }
    ]
}'
```

* Create a Manager User
```shell
curl --location 'http://localhost:8080/admin/realms/master/users' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <TOKEN>' \
--data-raw '{
    "username": "mgr1-test",
    "groups" : ["fr1-mgr"],
    "enabled" : "true",
    "email" : "2@2.com",
    "firstName": "fn",
    "lastName": "ln",
    "attributes": { "org_id" : "rh", "is_internal" : "true", "org_admin": "true", "type" : "manager"},
    "credentials" : [
        { "credentialData" : "fr1-mgr", "temporary" : "false", "type" : "password", "value" : "fr1-mgr" }
    ]
}'
```

### Use Find Users API in userservice
* Call Find Users API with following criteria
```shell
# Search users with usernames
curl --location 'http://localhost:8000/users?org_id=rh&usernames=mgr1-test%2Ceng1-test'
```
```shell
# Search users with emails
curl --location 'http://localhost:8000/users?org_id=rh&emails=1%401.com%2C2%402.com'
```
```shell
# Search users with userids, remember to replace user ids below as they will be different in your case.
curl --location 'http://localhost:8000/users?org_id=rh&user_ids=c2979a54-b50e-473a-8ff8-0710f701e64f%2C3c577a73-d15a-4130-968b-1fdab10e0ee0'
```

## Docker Tasks
* Build userservice Docker image
```shell
docker build --rm --tag userservice .
```

* Run userservice Docker image on port 8000
```shell
docker run -p 8000:8000 userservice
```

## Using different environment variables with .env files
* Currently the .env file is used for running the userservice locally on the host machine.
* .env.docker file is used for running the userservice in docker on the host machine with Keycloak service.
