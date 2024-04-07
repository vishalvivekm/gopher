<details><summary>curl reqs</summary>

```bash
curl localhost:8080/api -v
curl  curl https://nominatim.openstreetmap.org/search?q=\san%20francisco\&format\=json | jq

curl localhost:9090/api\?q=san%20francisco | jq

curl localhost:9090/api\?q=san%20francisco -v | jq
```

</details>

<details><other refs</summary>
json to struct:
https://mholt.github.io/json-to-go/

</details>
