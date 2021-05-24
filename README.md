# Solution

Entrypoints resides in cmd/{data,domain} dirs. Each of them are entrypoints for both domain and data services


# What's missing (no more time :( )

- Tests a lot o them. 
- Some structs should be initialized, configured in better way (e.g. not only single ones are conifiugred with envs etc.)
- No logs, metrics, repeating
- Not implemented requirements from task - e.g. graceful shutdown of GRPC server
- No docker-compose used 



# Running 

to run please disable local mongo or disable local mongo (no time for docker-compose)
and be sure to have ports 9090 and 9091 free (works on linux with network=host)

```sh
cp ${YOUR_PATH_TO_PORTS_JSON_FILE}/ports.json ./
make build     # builds docker files
make run       # runs everythuing on your local machine as detached docker containers be sure to have free ports


# when everything start 
curl -X POST http://localhost:9090/ports
```


