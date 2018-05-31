// -----------------------------------------------------------------------------------------------------------------------------
// GENERAL
// -----------------------------------------------------------------------------------------------------------------------------
//
// author: Sebastiaan Van Hoecke
// mail: sebastiaan@sevaho.io
//
// NOTE:
//
// -----------------------------------------------------------------------------------------------------------------------------

package main

import (


)

func get_content (s string) string {

    f := ""

    switch (s) {

//     case "":
//
//         f =
// `
// `
//         break
//
        case "usage":

        f =
`USAGE: ae [OPTIONS]... [ARGS]...

    ae, file and folder structure templater

OPTIONS:
   -h, --help       display the help and exit

   TO IMPLEMENT: 
   -d, --debug      debug, makes a debug folder and use that as root folder

ARGUMENTS:

   TO IMPLEMENT: 

   init {type} # initializes a whole project

       types:

       nginx
       python
       go
       nodejs
       bare

   file {type} # initializes only 1 file

       types:

       makefile
       gitignore
       dockerignore
       docker {type}

           types:

           nginx
           python
           go
           nodejs
           bare

       makefile {type}

           types:

           nodejs
           python

EXAMPLES:

   ./ae file python

NOTE:

    the -d flag is not implemented yet
    update is not implemented yet
`

    break

    case "makefile":

        f =
`ROOT_DIR := <<<REPLACE>>>
IP 		:= 1.1.1.1

help: ## Show this help message (default)]
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m \033[35m%s\033[0m\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## DEV : run locally
	echo "run locally, uncomment what you need in the Makefile"
	#npm install --prefix app # nodejs
	#npm start --prefix app # nodejs
	#pip install --user -r app/requirements.txt
	#python app/app.py # python
	#gunicorn --workers=2 --bind=0.0.0.0:5000 --chdir app app:app # python-flask

cleanup: ## DEV : cleanup
	echo "install locally"
	#rm -rf app/node_modules # nodejs
	#rm -rf app/etc # nodejs

install: ## DEV: install locally
	echo "install locally"
	echo "TO IMPLEMENT"

docker: ## DEV DOCKER: running docker
	echo "running docker locally, change docker run to docker run -p ***:*** if you need network capabilities"
	docker build -t "$(ROOT_DIR)" .
	docker run "$(ROOT_DIR)"

docker-daemon: ## DEV DOCKER: daemonize your docker container to the local system
	echo "running docker locally, change docker run to docker run -p ***:*** if you need network capabilities"
	docker build -t "$(ROOT_DIR)" .
	docker run -d --restart=always --name="$(ROOT_DIR)" "$(ROOT_DIR)"

docker-cleanup: ## DEV DOCKER : docker cleanup
	echo "stop docker container and remove the image so the app name can be used again"
	docker stop "$(ROOT_DIR)"
	docker rm "$(ROOT_DIR)"

prod-build: ## make local docker
	rm -rf app/node_modules
	docker build -t docker.sevaho.io/"$(ROOT_DIR)" .
	docker push docker.sevaho.io/"$(ROOT_DIR)"

prod-deploy: ## make local docker
	ssh -fL 2201:192.168.0.4:22 sevahoSSHServer sleep 5
	ssh sevaho@localhost -p 2201 "docker pull docker.sevaho.io/\"$(ROOT_DIR)\""
	ssh sevaho@localhost -p 2201 "docker run -d --net production --ip \"$(IP)\" --restart=always --name  \"$(ROOT_DIR)\" --dns 172.32.0.2 docker.sevaho.io/\"$(ROOT_DIR)\""

prod-cleanup: ## clean up docker
	ssh -fL 2201:192.168.0.4:22 sevahoSSHServer sleep 5
	ssh sevaho@localhost -p 2201 "docker ps | grep \"$(ROOT_DIR)\" && docker stop \"$(ROOT_DIR)\" || echo \"no live container $(ROOT_DIR)\""
	ssh sevaho@localhost -p 2201 "docker ps -a | grep \"$(ROOT_DIR)\" && docker rm \"$(ROOT_DIR)\" || echo \"no live container $(ROOT_DIR)\""
`

        break

    case "dockerfile_bare":

        f =
`# vim: set ft=dockerfile:

FROM alpine

MAINTAINER <<<REPLACE>>>

RUN apk --update add

ADD ./app /app

WORKDIR /app

RUN addgroup -S -g 1000 user && adduser -S -u 1000 -D -G user user
RUN chown -R user:user /app

USER user

ENTRYPOINT ["", ""]
`

        break

    case "dockerfile_python":

        f =
`# vim: set ft=dockerfile:

FROM python:3-alpine

MAINTAINER <<<REPLACE>>>

RUN apk --update add

ADD ./app /app

WORKDIR /app

RUN addgroup -S -g 1000 user && adduser -S -u 1000 -D -G user user
RUN chown -R user:user /app

USER user

RUN pip install -r requirements.txt

ENTRYPOINT ["python", "app.py"]
`

        break

    case "dockerfile_go":

        f =
`# vim: set ft=dockerfile:

FROM golang:1.9.5-alpine

MAINTAINER <<<REPLACE>>>

RUN apk --update add git

ENV APP_DIR $GOPATH/src/app

RUN mkdir -p $APP_DIR

ADD ./app $APP_DIR

WORKDIR $APP_DIR

RUN addgroup -S -g 1000 user && adduser -S -u 1000 -D -G user user
RUN chown -R user:user $APP_DIR

USER user

RUN sh requirements.sh
RUN go build -o app

ENTRYPOINT ["./app"]
`
        break

    case "dockerfile_nginx":

        f =
`# vim: set ft=dockerfile:

FROM nginx:alpine

MAINTAINER <<<REPLACE>>>

RUN apk --update add

ADD ./dist /var/www/html

ADD ./config/nginx.conf /etc/nginx/nginx.conf

RUN touch /var/run/nginx.pid && \
	chown -R nginx:nginx /var/run/nginx.pid && \
	chown -R nginx:nginx /var/cache/nginx

RUN chown -R nginx:nginx /var/www/html

USER nginx

EXPOSE 8000 44300
`
        break

    case "dockerfile_nodejs":

        f =
`# vim: set ft=dockerfile:

FROM node:9-alpine

MAINTAINER <<<REPLACE>>>

RUN apk --update add

ADD ./app /app

WORKDIR /app

RUN chown -R node:node /app

USER node

RUN npm install

ENTRYPOINT ["npm", "start"]
`

        break

    case "nginx_https":

        f =
`# vim: set ft=nginx:
###########################################################################################################
# www.example.be
###########################################################################################################
#
#	NOTES:
#
#
###########################################################################################################

# user nginx; # this is not needed anymore if we already run as non-root ea: nginx in an nginx-alpine container, if running as root we will need this line
worker_processes 1;

error_log  /var/log/nginx/error.log warn;
pid        /var/run/nginx.pid; # this is the socket where we talk to as nginx / root user, as non-root this file need to be chowned to the non-root user


events {

	worker_connections 1024;

}

http {

	include       /etc/nginx/mime.types;
	default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  65;

    gzip  on;

    ##
    # Security
    ##

    add_header X-Robots-Tag "none";
    add_header X-Frame-Options "SAMEORIGIN"; # this causes problems if you work with iframes or embeddings, so comment this out if needed
    add_header X-Content-Type-Options "nosniff";
    add_header X-XSS-Protection "1; mode=block";
    add_header X-Permitted-Cross-Domain-Policies "none";
    add_header X-Download-Options "noopen";

    ##
    # Map
    ##

    map $sent_http_content_type $expires {

        default                    off;
        text/html                  epoch;
        text/css                   max;
        application/javascript     max;
        ~image/                    max;

    }

    ##
    # HTTP
    ##

    server {

        listen 8000;
        listen [::]:8000;

        server_name <<<REPLACE>>>;

        server_tokens off;

        error_page 404 /404.html;
        error_page 500 502 503 504 /50x.html;

        return 301 https://$server_name$request_uri;

    }

    ##
    # HTTPS
    ##

    server {

        listen 44300 ssl; # http2 proxy_protocol;
        listen [::]:44300 ssl; # http2 proxy_protocol;

        server_name <<<REPLACE>>>;

        ssl                         on;
        ssl_session_timeout         1d;
        ssl_session_cache           shared:SSL:50m;
        ssl_session_tickets         off; 
        ssl_prefer_server_ciphers   on;
        ssl_dhparam                 /etc/dhparam/dh8192.pem;
        ssl_stapling                on;
        ssl_stapling_verify         on;
        ssl_protocols               TLSv1 TLSv1.1 TLSv1.2;

        ssl_ciphers 'ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:DHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES256-SHA256:DHE-RSA-AES256-SHA:DHE-RSA-CAMELLIA256-SHA:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-RSA-SEED-SHA:DHE-RSA-CAMELLIA128-SHA:HIGH:!aNULL:!eNULL:!LOW:!3DES:!MD5:!EXP:!PSK:!SRP:!DSS';

        ssl_certificate             /etc/ssl/tls.crt;
        ssl_certificate_key         /etc/ssl/tls.key;

        server_tokens off;

        index index.html;

        error_page 404 /404.html;
        error_page 500 502 503 504 /50x.html;

        expires $expires;   

        location ~ ^/(scripts|styles)/(.*)$ {

            root .tmp/;
            error_page 404 =200 @asset_pass;
            try_files $uri =404;
            break;

        }

        location @asset_pass {

            root /var/www/html;
            try_files $uri =404;

        }

        location / {

            root /var/www/html;
            expires -1;
            add_header Pragma "no-cache";
            add_header Cache-Control "no-store, no-cache, must-revalicate, post-check=0 pre-check=0";
            try_files $uri $uri/ /index.html =404;

        }

    }

}
`
        break

    case "nginx_http":

        f =
`# vim: set ft=nginx:
###########################################################################################################
# www.example.be
###########################################################################################################
#
#	NOTES:
#
#
###########################################################################################################

# user nginx; # this is not needed anymore if we already run as non-root ea: nginx in an nginx-alpine container, if running as root we will need this line
worker_processes 1;

error_log  /var/log/nginx/error.log warn;
pid        /var/run/nginx.pid; # this is the socket where we talk to as nginx / root user, as non-root this file need to be chowned to the non-root user

events {

	worker_connections 1024;

}

http {

	include       /etc/nginx/mime.types;
	default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  65;

    gzip  on;

    ##
    # Security
    ##

    add_header X-Robots-Tag "none";
    add_header X-Frame-Options "SAMEORIGIN"; # this causes problems if you work with iframes or embeddings, so comment this out if needed
    add_header X-Content-Type-Options "nosniff";
    add_header X-XSS-Protection "1; mode=block";
    add_header X-Permitted-Cross-Domain-Policies "none";
    add_header X-Download-Options "noopen";

    ##
    # Map
    ##

    map $sent_http_content_type $expires {

        default                    off;
        text/html                  epoch;
        text/css                   max;
        application/javascript     max;
        ~image/                    max;

    }

    ##
    # HTTP
    ##

    server {

        listen 8000;
        listen [::]:8000;

        server_name <<<REPLACE>>>;

        server_tokens off;

        error_page 404 /404.html;
        error_page 500 502 503 504 /50x.html;

        index index.html;

        error_page 404 /404.html;
        error_page 500 502 503 504 /50x.html;

        expires $expires;   

        location ~ ^/(scripts|styles)/(.*)$ {

            root .tmp/;
            error_page 404 =200 @asset_pass;
            try_files $uri =404;
            break;

        }

        location @asset_pass {

            root /var/www/html;
            try_files $uri =404;

        }

        location / {

            root /var/www/html;
            expires -1;
            add_header Pragma "no-cache";
            add_header Cache-Control "no-store, no-cache, must-revalicate, post-check=0 pre-check=0";
            try_files $uri $uri/ /index.html =404;

        }

    }

}
`
        break

    case "dockerignore":

        f =
`node_modules
npm-debug.log
`
        break

    case "gitignore":

        f =
`node_modules
app/node_modules
dist
`
        break

    case "app.go":

        f =
`// -----------------------------------------------------------------------------------------------------------------------------
// GENERAL
// -----------------------------------------------------------------------------------------------------------------------------
//
// author: Sebastiaan Van Hoecke
// mail: sebastiaan@sevaho.io
//
// NOTE:
//
// -----------------------------------------------------------------------------------------------------------------------------

package main

import (

    "fmt"
    "os"
    "flag"

)

// -----------------------------------------------------------------------------------------------------------------------------
// CONSTANTS
// -----------------------------------------------------------------------------------------------------------------------------

var target      = ""
var port        = 0

// -----------------------------------------------------------------------------------------------------------------------------
// FUNCTIONS
// -----------------------------------------------------------------------------------------------------------------------------

func usage () {

    fmt.Printf(` + "`" + `Usage: %s [OPTIONS]... [ARGS]...

    description

OPTIONS:

    -h, *               display the help and exit
    -t, --target        target
    -p, --port          port

EXAMPLES:

    %s -t 127.0.0.1 -p 8000

NOTE:

` + "`" + `, os.Args[0], os.Args[0])

    os.Exit(0)

}

// -----------------------------------------------------------------------------------------------------------------------------
// MAIN
// -----------------------------------------------------------------------------------------------------------------------------

func main () {

    argument_help := flag.Bool("h", false, "")
    argument_target := flag.String("t", "", "")
    argument_port := flag.Int("p", 0, "")

    flag.Parse()

    if (*argument_help) {
    
        usage()

    }

    target = *argument_target
    port = *argument_port

}
`
        break

    case "requirements.sh":

        f =
`// go get github.com/maintainer/pkg/..
`
        break

    case "ae_lib.go":

        f =
`// ae go library
`

        break


    case "app.py":

        f =
`#! /usr/bin/python

# -----------------------------------------------------------------------------------------------------------------------------
# GENERAL
# -----------------------------------------------------------------------------------------------------------------------------
#
# author: Sebastiaan Van Hoecke
# mail: sebastiaan@sevaho.io
#
# NOTE:
#
# -----------------------------------------------------------------------------------------------------------------------------

import sys
import getopt

# -----------------------------------------------------------------------------------------------------------------------------
# GLOBAL VARIABLES
# -----------------------------------------------------------------------------------------------------------------------------

target          = ""
port            = 0

# -----------------------------------------------------------------------------------------------------------------------------
# CLASSES
# -----------------------------------------------------------------------------------------------------------------------------

# -----------------------------------------------------------------------------------------------------------------------------
# FUNCTIONS
# -----------------------------------------------------------------------------------------------------------------------------


def usage ():

    print("Usage: python %s [OPTIONS]... [ARGS]... \
            \n \
            \n\tdescription\
            \n \
            \nOPTIONS:\
            \n \
            \n\t-h, *               display the help and exit \
            \n\t-t, --target        target \
            \n\t-p, --port          port \
            \n \
            \nEXAMPLES:\
            \n \
            \n\tpython %s -t 127.0.0.1 -p 8000 \
            \n \
            \nNOTE:\
            \n \
            " % (sys.argv[0], sys.argv[0]))

    sys.exit(1)


def banner ():

    print("A Banner \
            \n \
            ")

# -----------------------------------------------------------------------------------------------------------------------------
# MAIN
# -----------------------------------------------------------------------------------------------------------------------------


def main (argv):

    global target, port

    banner()

    # if not len(argv):
    #     usage()

    try:

        opts, args = getopt.getopt(argv, "ht:p:", [
            "help",
            "target",
            "port"])

    except getopt.GetoptError as err:

        print("error: %s" % (str(err)))
        usage()

    for opt, arg in opts:

        if opt in ("-h", "--help"):
            usage()
        elif opt in ("-t", "--target"):
            target = arg
        elif opt in ("-p", "--port"):
            port = int(arg)
        else:
            assert False, "Unhandled Option"

    print("What would you like to do captain? \
            \n \
            \n  1) TODO \
            \n  99) exit \
            \n \
            ")

    mode = input()
    mode = int(mode)

    if mode == 1:
        print("mode")

    elif mode == 99:
        sys.exit(1)


if __name__ == "__main__":

    main(sys.argv[1:])
`

        break

    case "ae_lib.py":

        f =
`# ae python library
`

        break

    case "requirements.txt":

        f =
`
`

        break

    case "app.js":

        f =
`//! /usr/bin/node

// GENERAL
// -----------------------------------------------------------------------------------------------------------------------------
//
// author: Sebastiaan Van Hoecke
// mail: sebastiaan@sevaho.io
//
// NOTE:
//
// -----------------------------------------------------------------------------------------------------------------------------

'use strict';

/*
 *      load all variables from your .env file as process variables
 *      call them by process.env.VARIABLE
 */

// require('dotenv').config()

// const ae        = require('./ae_lib.js');

// -----------------------------------------------------------------------------------------------------------------------------
// GLOBAL VARIABLES
// -----------------------------------------------------------------------------------------------------------------------------

var target      = ""
var port        = 0
var options = process.argv;

// -----------------------------------------------------------------------------------------------------------------------------
// CLASSES
// -----------------------------------------------------------------------------------------------------------------------------

// -----------------------------------------------------------------------------------------------------------------------------
// FUNCTIONS
// -----------------------------------------------------------------------------------------------------------------------------

function usage () {

    console.log("Usage: node %s [OPTIONS]... [ARGS]..." + 
                "\n" +
                "\n\tdescription" +
                "\n" +
                "\nOPTIONS:" +
                "\n" +
                "\n\t-h, *            display the help and exit" +
                "\n\t-t, --target     target" +
                "\n\t-p, --port       port" +
                "\n" +
                "\nEXAMPLES:" +
                "\n" +
                "\n\tnode %s -t 127.0.0.1 -p 8000 " +
                "\n" +
                "\nNOTE:" +
                "\n" +
                "", process.argv[1], process.argv[1]);

    process.exit(0);

}

function banner () {

    console.log("A banner" +
                "\nMake something astonishing" +
                "\n"
                );

}

function argparser (option) {

    let opt = false
    let arg = ""

    if (options.indexOf(option) > -1) {
    
        opt = true;
        arg = options[options.indexOf(option) + 1];

    } 

    return {opt, arg}

}

// -----------------------------------------------------------------------------------------------------------------------------
// MAIN
// -----------------------------------------------------------------------------------------------------------------------------

function main () {

    options.splice(0, 2);

    // needs rework

    try {
    
        for (let i = 0; i < options.length; i ++) {

            if (argparser("-h").opt || argparser("--help").opt) {
            
                options.splice(options.indexOf("-h"), 1);
                options.splice(options.indexOf("--help"), 1);

                usage();

            } else if (argparser("-t").opt || argparser("--target").opt) {
            
                target = argparser("-t").arg || argparser("--target").arg;
                options.splice(options.indexOf("-t"), 1);

            } else if (argparser("-p").opt || argparser("--port").opt) {
            
                port = argparser("-p").arg || argparser("--port").arg;
                options.splice(options.indexOf("-p"), 1);

            }

        }

    } catch (e) {

        console.error(e);

    }

    console.log(target, port);

}

main();
`

        break

    case "dotenv":

        f =
`KEY=VALUE
ENCRYPTION_KEY=TODO
ENCRYPTION_IV=TODO
ENCRYPTION_ALGORITHM=aes256
`

        break

    case "ae_lib.js":

        f =
`// ae library

'use strict'

const crypto   = require('crypto');


module.exports = {

    /*
     *      @encrypt
     */

    encrypt: function (text, algo, key, iv) {

        let cipher = crypto.createCipheriv(algo, key, iv);
        let crypted = cipher.update(text, 'utf8', 'hex');

        crypted += cipher.final('hex');

        return crypted;

    },

    /*
     *      @decrypt
     */

    decrypt: function (text, algo, key, iv) {

        let decipher = crypto.createDecipheriv(algo, key, iv);
        let dec = decipher.update(text,'hex','utf8');

        dec += decipher.final('utf8');

        return dec;

    },

    /*
     *      @base64_encode
     */

    base64_encode: function (payload) {

        return Buffer.from(payload).toString('base64');

    },

    /*
     *      @base64_decode
     */

    base64_decode: function (payload) {

        return Buffer.from(payload, 'base64').toString('ascii');

    }

}
`

        break

    case "package.json":

        f =
`{
  "name": "<<<NAME>>>",
  "version": "1.0.0",
  "description": "app.js service",
  "main": "app.js",
  "scripts": {
    "start": "node app.js",
    "build": "ng build --prod"
  },
  "keywords": [
    "util",
    "functional",
    "server",
    "client",
    "browser"
  ],
  "author": "<<<AUTHOR>>>",
  "contributors": [],
  "dependencies": {
  }
}
`

        break

    case "completion":

        f =
`zsh completion
`

        break

    default:

        f = ""

    }

    return f

}
