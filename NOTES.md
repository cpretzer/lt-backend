# Test Commands

## cURL

* Add user
`curl -X POST http://localhost:3001/users/add -d '{"first_name":"Post", "last_name":"Test", "email_address":"post_test@chabrina.com", "username":"posttest", "password":"wordpass", "username": "posttest"}' -H "content-type: application/json; charset=UTF-8"`