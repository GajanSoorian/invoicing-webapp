# Parallax Invoicing Web Application

A web app that allows creating, updating, and displaying invoices.

## API endpoints

| Action  | Endpoint                        |
|---------|---------------------------------|
| Create  | POST   /v1/invoice/create       |
| Update  | PATCH  /v1/invoice/update/{id}  |
| Display | GET    /v1/invoice/display/{id} |

"v1" corresponds to the api version.

## commands to run application

go to path "parallax-invoicing" , which had the "makefile". Then run on the below commands:

    `make deps` : fetch all dependent libraries.

    `make test` : Run all unit tests.

    `make run`  : Build and run the application.  
    
    `make all`  : Do all the above actions in sequence.

Todo Now:

- Nice DB handler
- Create handlers
- Update handlers
- Modify handlers
- Error handling

## Design decision

- Test Driven Development
- Loosely coupled architecture

## Corner of hope (did not get to these)

- JWT authorization for rest endpoints.
- Integration test cases with mocked Database.
- create end points for gRPC
- Dockerize everything
- Create an analytics page to
        - display customer purchasing pattern to predict future order timelines.

## Out of scope

- List view for multiple invoices.
- Login and authentication.
- API security (Token based authentication and TLS setup).
- Timeouts(api request/ db query), rate limiting and retry strategies.
- Logger.
- Test cases into test suites.
- Integration test cases.
- Support for different unit of cost and conversions.
- Tax logic for items.
- Incremental "vertical slice" commits for easy review and faster feedback.

## Go information

- Packages help  organize related Go source files together into a single unit, making them modular, reusable, and maintainable.
- Methods that start with a capital letter are services offered by a package (Similar to public methods in OOP languages).
- Units tests are part of the same packages, placed in the folder and "_test" suffix at the end.
