# NGINX Server Sample Application

## Building

```bash
pack build my-nginx-app --buildpack docker.io/paketobuildpacks/nginx
```

## Running

```bash
docker run --tty --env PORT=8080 --publish 8080:8080 my-nginx-app
```

## Viewing

```bash
curl -s localhost:8080
```
