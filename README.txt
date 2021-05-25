# KrakenJWT

A small server that demonstrates how to use [krakend](https://github.com/devopsfaith/krakend-ce)

## Run the server

	$ go run ./cmd/randomsrv

The server can be reached on [http://127.0.0.1:4680](http://127.0.0.1:4680).

## Connect the server to krakend

First, check out and build krakend-ce

	$ git clone https://github.com/devopsfaith/krakend-ce.git
	$ cd krakend-ce
	$ make build

Then start krakend with the configuration file in this repository

	$ ./krakend run -c ../krakend.json

krakend can be reached at [http://127.0.0.1:8080](http://127.0.0.1:8080/)

## Make requests

	$ curl -s http://127.0.0.1:8080/t/236 | jq .
	
```json
{
  "address": {
    "address": "1579 Port Village berg, North Bashirian, Kentucky 33261",
    "city": "North Bashirian",
    "country": "Somalia",
    "latitude": -86.33147,
    "longitude": 106.141713,
    "state": "Kentucky",
    "street": "1579 Port Village berg",
    "zip": "33261"
  },
  "car": {
    "brand": "Daewoo",
    "fuel": "Diesel",
    "model": "New Beetle",
    "transmission": "Automatic",
    "type": "Passenger car medium",
    "year": 1981
  },
  "contact": {
    "email": "leviwuckert@lemke.com",
    "phone": "6236687431"
  },
  "credit_card": {
    "exp": "07/22",
    "number": "XXXXXXXXXXXXXXXX",
    "type": "JCB"
  },
  "domain": "internationaleyeballs.name",
  "first_name": "Rafaela",
  "gender": "female",
  "job": {
    "company": "J.P. Morgan Chase",
    "descriptor": "Legacy",
    "level": "Interactions",
    "title": "Engineer"
  },
  "last_name": "Nikolaus",
  "profile_picture_url": "https://picsum.photos/300/300/people",
  "ssn": "808594433"
}

```

## License

MIT