# Go-Digit

## Install Steps

- Install Ruby (this project was spiked using MRI 2.3.1). You can do this via a Ruby version manager like [rbenv](https://github.com/rbenv/rbenv) or [RVM](https://rvm.io/).
- Install PostgreSQL. On Mac, this is done with `brew install postgresql`
- Install all project dependencies using bundler. Install using `gem install bundler` and then run `bundle install` at the root of the project directory.
- Create the database using `rails db:create`. This will have Rails connect to Postgres as the default user and create the go-digit development and test databases.

## Running the Project

You need to serve the front end. This can be done in the npm folder, or just by running a `rake npm:serve` or a `rake gulp:serve`. This will start up the front end server at localhost:3000.
Then you can start up the server with rails s -p 3030. This will start up the server at localhost:3030.
