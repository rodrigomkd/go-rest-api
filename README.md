# GO REST API
=================

Basic template to develop an API with Go!

Usage
-----

Clone the repo:

    git clone https://github.com/rodrigomkd/go-rest-api
    cd go-rest-api

# Download packages:
    * gorilla/mux: `go get github.com/gorilla/mux`
    * BurntSushi/toml: `go get github.com/BurntSushi/toml`
    * stretchr/testigy: `go get -u github.com/stretchr/testify`

Set properties variables

    ServerPort=3000 <SERVER PORT> 
    DataSource="data.csv" <CSV PATH>
    DataSourceWorker="data_worker.csv" <CSV PATH>
    ApiUri="https://fakerestapi.azurewebsites.net/api/v1/Activities" <API URL>

Run the server

    go run main.go

Try the endpoints:

    curl -XGET http://localhost:3000/api/v1/items
    curl -XGET http://localhost:3000/api/v1/items/1
    curl -XGET http://localhost:3000/api/v1/items/workers/type=odd&items=10&items_per_workers=2
    curl -XPOST http://localhost:3000/api/v1/items/sync

License
-------

MIT, see LICENSE file