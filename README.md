# transmission-rest

a CRUD REST interface for transmission's RPC. Written in Go using `github.com/hekmon/transmissionrpc/v2`

### Configuration

Empty config files are stored in `config\.env.[local|prod]`<BR>
If `IS_PROD` env variable is set to true, then env.prod is used. Otherwise env.local <br>
Config sample: <br>

```
# transmission (if running in docker on localhost, use "host" network driver to reach transmission host)
TRANSMISSION_HOST=XXXXX # ip adress of a host where transmission instance is running
TRANSMISSION_PORT=XXXXX # port which is used by TRANSMISSION_HOST
TRANSMISSION_PROTOCOL=X # protocol [http | https]
TRANSMISSION_USERNAME=X # username
TRANSMISSION_PASSWORD=X # password

# server
SERVER_PORT=XXXXXXXX # port on which transmission-rest is listening
SERVER_BASE_PATH=XXX # part of the URI
SERVER_MAIN_ENTITY=X # part of the URI

# log
LOG_LEVEL=INFO # log level [INFO | DEBUG | WARN]
```

### Build

to generate docker image for your local system run:<br>
`make build-docker-image`

to generate docker image for Synology DS920+ run:<br>
`make build-docker-image-synology-920p`

files will be generated in `bin/` directory
