# react-nodejs-boilerplate by Wonjinsin

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app) + Craco(https://github.com/gsoft-inc/craco) + Redux(https://redux.js.org/) + Nginx(https://www.nginx.com/) + Node.js(https://nodejs.org) + Docker(https://www.docker.com/)

## Additional Info

Client folder is for react & server folder is for nodejs.<br />
Client communicate with server using proxypass(8000 port) and Default url for proxypass is /front/api, can be changed with client/src/setupProxy.<br />
Server session is based on Redis.<br />
you can also deplay your server using docker-compose.

### Command for client
```
// for local
$ cd client &&  yarn install && yarn start:local
// for docker
$ cd client &&  yarn install && yarn build:docker
// for build
$ cd client &&  yarn install && yarn build
```

### Command for server
```
// for local
$ cd server &&  yarn install && yarn start:local
// for docker
$ cd server &&  yarn install && yarn start:docker
// for server
$ cd client &&  yarn install && yarn start
```

### Docker command
```
docker create network dummy_network
docker-compose up -d
```