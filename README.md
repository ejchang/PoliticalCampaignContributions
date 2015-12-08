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

### Expected Schema
If you load in your own database, it should be in the following schema
```SQL
CREATE TABLE CongressMembers
(congress_id VARCHAR(32) NOT NULL PRIMARY KEY,
 name VARCHAR(64) NOT NULL,
 chamber VARCHAR(16) NOT NULL,
 state VARCHAR(16) NOT NULL,
 party VARCHAR(16) NOT NULL);

CREATE TABLE Bill
(bill_id VARCHAR(8) NOT NULL PRIMARY KEY,
name VARCHAR(512) NOT NULL,
description VARCHAR(512) NOT NULL,
date_voted VARCHAR(32) NOT NULL);

CREATE TABLE VOTED
(bill_id VARCHAR(8) NOT NULL REFERENCES Bill(bill_id),
name VARCHAR(32) NOT NULL,
state VARCHAR(4) NOT NULL,
vote VARCHAR(256) NOT NULL
CHECK(vote in ('Yes', 'No', 'Abstain')));

CREATE TABLE PAC
(pacID VARCHAR(64) NOT NULL PRIMARY KEY,
name VARCHAR(64) NOT NULL,
industry VARCHAR(64) NOT NULL);

CREATE TABLE PAC_Donations
(pacdonation_id BIGINT NOT NULL PRIMARY KEY,
pac_id VARCHAR(64) NOT NULL REFERENCES PAC(pacID),
congress_id VARCHAR(32) NOT NULL REFERENCES CongressMembers(congress_id),
amount INTEGER NOT NULL,
industry VARCHAR(64) NOT NULL);

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
