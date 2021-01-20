# Test Commands

## cURL

* Add user
`curl -X POST http://localhost:3001/users/add -d '{"firstName":"Post", "lastName":"Test", "emailAddress":"post_test_4@chabrina.com", "username":"posttest", "password":"wordpass", "username": "posttest"}' -H "content-type: application/json; charset=UTF-8"`

* Get user
`curl http://localhost:3001/users?id=recbR2ySUkyjHN52C`

* Get user by email (airtable)
`curl https://api.airtable.com/v0/$AIRTABLE_BASE/users?filterByFormula=%7BemailAddress%7D%3D%22post_test%40chabrina.com%22 \
    -H "Authorization: Bearer $AIRTABLE_KEY"`
  

* Update user
`curl -X PATCH http://localhost:3001/users/update -d '{"firstName":"Update", "lastName":"Test Squirrel", "emailAddress":"post_test_4@chabrina.com", "username":"updatetest", "password":"wordpass", "username": "posttest"}' -H "content-type: application/json; charset=UTF-8"`

* Delete user
`curl -X DELETE http://localhost:3001/users/delete -d '{"firstName":"Update", "lastName":"Test Update", "emailAddress":"post_test@chabrina.com", "username":"updatetest", "password":"wordpass", "username": "posttest"}' -H "content-type: application/json; charset=UTF-8"`
