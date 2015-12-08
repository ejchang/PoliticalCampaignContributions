# Go based backend for our Campaign Contributions App

### Installation:
1. Install go
2. Clone this repository into your GOPATH
3. Run go get to install dependencies

### Starting the server:
1. Go to the root of the directory
2. run "go run main.go" to launch the server

### Database
This server should be connected to a psql server. It is currently configured to connect to an Amazon RDS instance. To connect to this server, create a config folder in the main directory, and create the file config.json. In this file add the following:
```JSON
{
    "Database": {
      "username": "*database username*",
      "password":  "*database password*"
    }
}
```

### API
The backend supports the following API Calls

GET /CONGRESS

Returns a json list of maps
i.e.
```JSON
[{"chamber":"sen","id":"N00003535","name":"Sherrod Brown","party":"Democrat","state":"OH"},{"chamber":"sen","id":"N00007836","name":"Maria Cantwell","party":"Democrat","state":"WA"},...{..}]
```

GET /MEMBER/{MEMBERID}

Returns the following information about a congressmember: 

Party

State

How they voted on bills

i.e.
```JSON
{"bills":[{"billid":"h511","billname":"","description":"On Agreeing to the Amendment: Amendment 5 to H R 348","vote":"No"},{"billid":"h523","billname":"Womenâ€™s Public Health and Safety Act","description":"On Motion to Recommit with Instructions: H R 3495 Womenâ€™s Public Health and Safety Act","vote":"No"},...],"donors":[{"amount":126000,"industry":"Defense aerospace contractors","pacs":{"Aerojet Rocketdyne":2000,...}},..]}
```

GET /BILL/{BILLID}

Returns a JSON map containing bill information as well as the how each congress member voted on the bill

i.e.
```JSON
{"billid":"bill1","billname":"bill1name","description":"bill1desc","voteDate":"bill1votedate","votes":[{"id":"N00036633","member":"rep1","party":"Republican","state":"state1","vote":"No"},...]}
```
