# Gozen

<img src="https://github.com/ignitedcms/gozen/blob/main/resources/images/goph.svg" width="250">

## About
Gozen is full stack minimalistic framework for golang.

Go developers are traditionally very picky when it comes to what and what 
they don't want. We've built Gozen with this in mind. It only uses go-chi,
a super lightweight Go router and a few Gorilla pacakges. 
Zero ORMs or any other third party libraries. Ensuring Gozen
is lightweight and giving you confidence that things won't break in the future.

## Who is this for?
If you're itching for a Laravel, Ruby on Rails or Django python like development
experience in Go, you'll fall in love with Gozen.

### Features
- Scaffolding 
- Middleware
- Templating
- Controllers 
- Models 
- Emails
- Authentication 
- Form validation 
- Web sockets
- Sessions
- Rate limiting
- Sql builders
- Cors
- API only mode (for usage with SPAs)
- Testing
- CSRF protection
- Flash data
- File uploads

### Small dependency chain
Gozen only relies on 'go-chi' and a few 'gorilla' repositories

### Installation
Please ensure you have at least go version 1.22

First Git clone the repository, then cd into that directory and run

```
git clone https://github.com/ignitedcms/gozen.git
cd gozen
go run -v .
```

Please note, it may take some time to download and compile the initial repository.
It can take upto 5 minutes. Please be patient. Subsequent execution will be fast,
typically less than 3 seconds.

This will spin up a server on

```
http://localhost:3000
```

We recommend using [Go Air](https://github.com/air-verse/air) if you wish to do continuous development with hot reloading.

After, your first compilation use

```
air app.go
```

### Important
This project is currently in development, so is subject to frequent changes.
