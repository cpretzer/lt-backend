# Test Commands

## cURL

* Add user
`curl -X POST http://localhost:3001/users/add -d '{"first_name":"Post", "last_name":"Test", "email_address":"post_test@chabrina.com", "user_name":"posttest", "password":"wordpass"}' -H "content-type: application/json; charset=UTF-8"`