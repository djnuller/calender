# Calender API

## Environment
In the root folder of the project, place a .env file containing:
```
mongouri="mongodb://mongouri"
database="CalenderApi"
userTable="users"
taskTable="tasks"
pendingTable="pending"
eventTable="events"
access_secret="supersecret"
```

## User routes
| Method        | Endpoint         | Description  |
| ------------- |------------------| ------------ |
| POST          | /users           | Register user|
| POST          | /users/signin    | Sign in user |
| DELETE        | /users/{user_id} | Delete user  |
| PUT           | /users/{user_id} | Update user  |
