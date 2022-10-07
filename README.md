# Parallax Invoicing Web Application

A web app that allows creating, updating, and displaying invoices.

## API endpoints

| Action  | Endpoint                        |
|---------|---------------------------------|
| Create  | POST   /v1/invoice/save         |
| Update  | POST  /v1/invoice/save          |
| Display | GET    /v1/invoice/view/{id}    |

"v1" is the api version.

## Install Go

### Option 1

Download go package from : https://go.dev/dl/
Go versions are backward compatible so please install any version that matches your OS version and architecture. That being said, newer version installation might not work on older OS( Go version 1.16.15 would be a safe choice).
For older OSes, archived version of Go installers can be found in `archived version` drop down in the same link.

This project was built using Go 1.16.15. To install the same version use one of the below links:

For Darwin amd64 (x86-64 arch): https://go.dev/dl/go1.16.15.darwin-amd64.pkg

For Darwin arm64 (ARM64 arch): https://go.dev/dl/go1.16.15.darwin-arm64.pkg

Follow the instuction to install Go. Once complete, if `go version` does not does not recognize go as a command- add `export PATH="/usr/local/go/bin:$PATH"` to ~/.bashrc or ~/.bash_profile to ensure `go` command works.

### Option 2

Go can also be installed using homebrew:
https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5

Steps to set up Go workspace are not required to build and run this application and can be ignored.

Once installed, running  `go version` should display a version number.

This application was testing on my partner's Mac using option 1, please let me know if you're any difficulties and I'd be happy to help!

## Commands to build and run invoice-service application

Navigate to the folder "parallax-invoicing" , which contains the "makefile" (A wrapper for Go build and run commands) Then run on the following commands:

    `make deps` : fetch all dependent libraries.

    `make test` : Run all unit tests.

    `make run`  : Build and run the application.  

Environment variables(host IP, port, DB info) needed for the service are stored in `invoice-service/invoice-service.env`. 
For this MVP, an online database has been setup for in https://bit.io/GajanSoorian/invoice_db and the relevant connection information have been set in the `invoice-service.env` file. I will add the reviewer as collaborator for this Database as well. Usage is intuitive, but please let me know if you have any questions.
Any SQL DB should work as long the env file is updated with connection information.

On successful launch of the invoice-service Go application, the following command prompt will be displayed:
`[GIN-debug] Listening and serving HTTP on :3000`

## Commands to start angular web application

Navigate to the `invoice-webapp` folder and run the below commands:

- npm install    : To install dependencies.
- npm run build  : Build the project.
- npm run test   : Run the test cases
- npm run start  : Run the application.

Use any browser to open: http://localhost:4200/
Note: This project works with Angular 14. It has not been tested in other angular versions.

Complete ng versions: Angular CLI: 14.2.4 , Node: 14.20.1 ,Package Manager: npm 8.19.2

## Corner of hope (did not get to these)

- Integration test cases with mocked Database.
- More frontend test cases.
- Better error handling and error propagation.
- Separate routes for different views for manageable code.
- JWT authorization for rest endpoints.
- Dockerize everything
- Create an analytics dashboard page as the login page to:
        - display customer purchasing pattern to predict future order timelines.
        - upcoming payments
        - Identify dip/increase in order requests for items.
- Decouple DB operations with the type of Database present.
- Pagination.

## Out of scope

- Login and authentication.
- API security (Token based authentication and TLS setup).
- Timeouts, rate limiting and retry strategies.
- Logger functionality.
- Alternate API endpoints- gRPC
- Test suites to better manager similar test cases.
- Integration test cases.
- Support for different unit of cost and conversions.
- Tax logic for items.
  
## Tech stack

- Angular 14
- Go 1.16
- Postgres

## Go information

- Packages help  organize related Go source files together into a single unit, making them modular, reusable, and maintainable.
- Methods that start with a capital letter are services offered by a package (Similar to public methods in OOP languages). Likewise, entities that start with a small letters are private to the package.
- Units tests are part of the same packages, placed in the folder and "_test" suffix at the end.
