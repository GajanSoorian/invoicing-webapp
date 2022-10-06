# Parallax Invoicing Web Application

A web app that allows creating, updating, and displaying invoices.

## API endpoints

| Action  | Endpoint                        |
|---------|---------------------------------|
| Create  | POST   /v1/invoice/save         |
| Update  | POST  /v1/invoice/save          |
| Display | GET    /v1/invoice/view/{id}    |

"v1" denotes the api version.

## Commands to run invoice-service application

Navigate to the folder "parallax-invoicing" , which contains the "makefile". Then run on the below commands:

    `make deps` : fetch all dependent libraries.

    `make test` : Run all unit tests.

    `make run`  : Build and run the application.  
    
    `make all`  : Do all the above actions in sequence.

Environment variables(host IP, port, DB info) needed for the service are stored in invoice-service/invoice-service.env

## Commands to start angular web application

Navigate to the `invoice-webapp` folder and run the below commands:

- npm install  : To install dependencies.
- npm run build  : Build the project.
- npm run start -o  : Run the application.

Note: This has not been tested in newer/older angular versions.

complete ng versions: Angular CLI: 14.2.4 , Node: 14.20.1 ,Package Manager: npm 8.19.2

## Corner of hope (did not get to these)

- Integration test cases with mocked Database.
- Angular test cases for frontend.
- Separate endpoint for Modify (PATCH)
- JWT authorization for rest endpoints.
- Dockerize everything
- Create an analytics dashboard page as the login page to:
        - display customer purchasing pattern to predict future order timelines.
        - upcoming payments
        - Identify dip/increase in order requests for items.
- Use interfaces to make Database queries and operations independent Database type.
- Pagination.

## Out of scope

- Login and authentication.
- API security (Token based authentication and TLS setup).
- Timeouts(api request/ db query), rate limiting and retry strategies.
- Logger functionality.
- Alternate API endpoints- gRPC
- Test suites to better manager similar test cases.
- Integration test cases.
- Support for different unit of cost and conversions.
- Tax logic for items.
- Incremental "vertical slice" commits for easy review and faster feedback.
  
## Tech stack

- Angular 14
- Go 1.16
- Postgres

## Go information

- Packages help  organize related Go source files together into a single unit, making them modular, reusable, and maintainable.
- Methods that start with a capital letter are services offered by a package (Similar to public methods in OOP languages).
- Units tests are part of the same packages, placed in the folder and "_test" suffix at the end.
