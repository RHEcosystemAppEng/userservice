package test

var (
	KEYCLOAK_GET_TOKEN_URL      = "http://localhost:8080/realms/master/protocol/openid-connect/token"
	KEYCLOAK_GET_TOKEN_RESPONSE = `{
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJBWHNqSlFmbkMtdlRWVGpiSFViRTl1VS1mZUZPWnBEZVgwcUxnbWJlY0xZIn0.eyJleHAiOjE2ODM4MDg5NDcsImlhdCI6MTY4MzgwODg4NywianRpIjoiNTk3MmEyMWYtOWFhYy00M2M5LWI3MmMtYWIwZDgyZTE2YmVjIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9tYXN0ZXIiLCJzdWIiOiIxODE1NGJhNi0wYTY2LTQxNGQtODNjOS04YzYwZGZmNGMxYWQiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJhZG1pbi1jbGkiLCJzZXNzaW9uX3N0YXRlIjoiMzBkZTUyNzQtODhhYS00MmNmLWJlZWQtMDM1NmUwMjc3MGJlIiwiYWNyIjoiMSIsInNjb3BlIjoicHJvZmlsZSBlbWFpbCIsInNpZCI6IjMwZGU1Mjc0LTg4YWEtNDJjZi1iZWVkLTAzNTZlMDI3NzBiZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiYWRtaW4ifQ.J6P7-myYtYEjiVW-IWH-euOkWJP11N8tRs-Fs89pyAB0xHOLNMQpIRQ1mUW6Cr_axqZo5Z8AyXiWF-IQnHRzO8aWg6B2XwyzDgc887RXa__qrLXOZK0OUfc9T_O_t0VI7_y593E0KPkTodSi9M_7W2txE3XaQ0n1EvoTSINM9ZasLe20S5ij7ZTbMVER0g8P9AJxWOD6LU5xdWIF068WpVzGEvEKD2oHlBbcy8fdH_uavWloHoXXbb-jpJAxzYU3VsNIOLJYXXolkSAvkUOFXEuirvA764uYjCYOj0MIq4VWvPYVzEjt1BDELzLRYnd3UvwnTBFcVm4sh8cY89hOKg",
    "expires_in": 60,
    "refresh_expires_in": 1800,
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI5ZWFiNTBjOS05MWZiLTRlYjctOTQxYy02MjcyYTM2YTU2MmIifQ.eyJleHAiOjE2ODM4MTA2ODcsImlhdCI6MTY4MzgwODg4NywianRpIjoiYjYxOWMwMmQtZTZiNy00MjY0LWI5NGQtZmI1MzMxYzQyODljIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9tYXN0ZXIiLCJhdWQiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvcmVhbG1zL21hc3RlciIsInN1YiI6IjE4MTU0YmE2LTBhNjYtNDE0ZC04M2M5LThjNjBkZmY0YzFhZCIsInR5cCI6IlJlZnJlc2giLCJhenAiOiJhZG1pbi1jbGkiLCJzZXNzaW9uX3N0YXRlIjoiMzBkZTUyNzQtODhhYS00MmNmLWJlZWQtMDM1NmUwMjc3MGJlIiwic2NvcGUiOiJwcm9maWxlIGVtYWlsIiwic2lkIjoiMzBkZTUyNzQtODhhYS00MmNmLWJlZWQtMDM1NmUwMjc3MGJlIn0.zHs0LQCwi9q5VLDO9Q91Ian9AFYsbWROm3znGc84ivw",
    "token_type": "Bearer",
    "session_state": "30de5274-88aa-42cf-beed-0356e02770be",
    "scope": "profile email"
}`
	KEYCLOAK_USER_DATA1 = `[
    {
        "id": "779edba7-2dbe-44f5-89f8-38ccd3368b30",
        "username": "eng1-test",
        "email": "1@1.com",
        "firstName": "fn",
        "lastName": "ln",
        "org_admin": true,
        "is_internal": true,
        "type": "engineer",
        "attributes": {
            "is_internal": [
                "true"
            ],
            "org_admin": [
                "true"
            ],
            "org_id": [
                "rh"
            ],
            "type": [
                "engineer"
            ]
        }
    }
]`
	KEYCLOAK_USER_DATA2 = `[
    {
        "id": "6756946d-2bbc-423b-ab48-835567dae132",
        "username": "mgr1-test",
        "email": "2@2.com",
        "firstName": "fn",
        "lastName": "ln",
        "org_admin": true,
        "is_internal": true,
        "type": "manager",
        "attributes": {
            "is_internal": [
                "true"
            ],
            "org_admin": [
                "true"
            ],
            "org_id": [
                "rh"
            ],
            "type": [
                "manager"
            ]
        }
    }
]`
	KEYCLOAK_FIND_USERS_NO_PARAMS    = "http://localhost:8080/admin/realms/master/users"
	KEYCLOAK_FIND_USERS_BY_ORG_ID    = "http://localhost:8080/admin/realms/master/users?q=org_id:rh"
	KEYCLOAK_FIND_USERS_BY_EMAIL1    = "http://localhost:8080/admin/realms/master/users?email=1%401.com&q=org_id%3Arh"
	KEYCLOAK_FIND_USERS_BY_EMAIL2    = "http://localhost:8080/admin/realms/master/users?email=2%402.com&q=org_id%3Arh"
	KEYCLOAK_FIND_USERS_BY_USERNAME1 = "http://localhost:8080/admin/realms/master/users?q=org_id%3Arh&username=mgr1-test"
	KEYCLOAK_FIND_USERS_BY_USERNAME2 = "http://localhost:8080/admin/realms/master/users?q=org_id%3Arh&username=eng1-test"
	KEYCLOAK_FIND_USERS_BY_USERID1   = "http://localhost:8080/admin/realms/master/users?id=c2979a54-b50e-473a-8ff8-0710f701e64f&q=org_id%3Arh"
	KEYCLOAK_FIND_USERS_BY_USERID2   = "http://localhost:8080/admin/realms/master/users?id=3c577a73-d15a-4130-968b-1fdab10e0ee0&q=org_id%3Arh"
)
