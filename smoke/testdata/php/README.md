# PHP Sample App using Built-in Webserver

## Building

`pack build php-webserver-sample --env BP_PHP_WEB_DIR=htdocs --buildpack docker.io/paketobuildpacks/php`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-webserver-sample`

## Viewing

`curl http://localhost:8080`
