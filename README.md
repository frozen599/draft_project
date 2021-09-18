## How to run

### Required
- Docker

### Run
$ git clone https://github.com/gopher5889/interview.git
$ cd interview
$ docker build -t gopher .
$ docker run --name container_name interview:latest ./interview /path/to/email_template.json /path/to/customers.csv /path/to/output_emails/ /path/to/errors.csv


### Trivia
- This project is written in Go version 1.17.1