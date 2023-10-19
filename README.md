
# Crypto App Server
This project contains the backend for our mobile crypto app that I am developing with my friend who is an android developer in Kotlin.

Using this application, you can monitor the currency market using charts, as well as make fake investments (Currency rates are taken from a third-party API CoinMarket)


### Technology or what I learned
In this project, I built my own APIGateway, which redirects requests to the necessary microservices, performs logging and authentication using JWT

A separate microservice is a separate project that is started using docker-compose

I built a clean code microservice architecture using Gin

So that my front-end colleague can use my API, I use swagger to generate API documentation

### How to use
1. Clone this repository and rewrite .yaml configs so that DB url matches your db url on your local machine

2. Run docker-compose file, it'll build and run all apps together

#### NOTE: this projects is not finished yet. All implemented logic are listed in "Technology or what I learned" section






