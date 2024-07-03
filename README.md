# Gozen

<img src="https://github.com/ignitedcms/gozen/blob/main/resources/images/goph.svg" width="250">

Gozen is full stack framework for Golang with a ridiculously low dependency chain.

## About
Go developers are traditionally very picky when it comes to what and what 
they don't want. We've built Gozen with this in mind. It only uses go-chi,
a super lightweight Go router and a few Gorilla pacakges. 
Zero ORMs or any other third party libraries. Ensuring Gozen
is lightweight and giving you confidence that things won't break in the future.

## Who is this for?
If you're itching for a Laravel, Ruby on Rails or Django python like development
experience in Go, you'll fall in love with Gozen.

### Features
Scaffolding, middleware, templating, controllers, models, emails,
authentication, form validation, web sockets, sessions, rate limiting,
sql builders, cors, testing, csrf protection, flash data, file uploads

### Small dependency chain
Gozen only relies on 'go-chi' and a few 'gorilla' repositories

### Installation
Please ensure you have at least go version 1.22
Git clone the repository

```
go run -v .
```

Please note, it may take some time to download and compile the initial repository.
It can take upto 5 minutes. Please be patient. Subsequent execution will be fast,
typically less than 3 seconds.

We recommend using 'go air' if you wish to do continuous development with hotreloading.

