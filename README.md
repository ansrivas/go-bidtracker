bid-tracker:
Bid Tracker
---

![Test](https://github.com/ansrivas/go-bidtracker/workflows/Test/badge.svg)

A simple project to implement bid-tracker interface and concrete implementation with the following functionality:

- record a user’s bid on an item;
- get the current winning bid for an item;
- get all the bids for an item;
- get all the items on which a user has bid;
- build simple REST API to manage bids.
- No persistent store(events are for reporting only). 
- Try and use `fiber` for this.

#### Why does it exist?
Just wanted to try using `fiber` for a relatively okay-sized project.

#### Run using docker
```bash
docker build --network=host --build-arg VCS_REF=`git rev-parse --short HEAD` --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"`  -t ansrivas/bid-tracker:latest -f Dockerfile .

docker run -p 3000:3000 -i --rm ansrivas/bid-tracker:latest
```

Now navigate to `http://localhost:3000/swagger`

#### Examples:
1. Insert a new bid:
    ```
    curl -H 'Content-Type: application/json' -d '{"useruuid":"ae8f7716-867b-4479-b455-c5769e7475ba", "itemuuid": "b2f9ee6d-79fe-4b14-9c19-35a69a89219a", "timestamp": 1212321, "amount":32}' http://localhost:3000/api/v1/bids | jq
    ```
