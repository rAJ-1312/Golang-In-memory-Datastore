To run the server type the following command in the terminal
(make sure you are in this directory before running)

go run main.go

By default the server will run on PORT=8000, in case you want to use different PORT just change the PORT number in the .env file


You can then send requests to the server using tools like `curl` or a web browser.
For example, to set the value of the key "foo" to "bar", you can send the following
request:

curl --location --request POST 'http://127.0.0.1:8000/set' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "key" : "foo",
    "value" : "bar"
}'


And to retrieve the value of the key "foo", you can send the following request:

curl --location --request GET 'http://127.0.0.1:8000/get/foo'


Check the Report.pdf for more detailed explanation.

