####################################################################################################
# Database Tasks
####################################################################################################

db/psql: # @HELP connect to locally running database
db/psql:
	psql ${DB_CONNECTION_STRING}

db/migrations/up: # @HELP run migrations to create database tables
db/migrations/up:
	migrate -path ./migrations -database ${DB_CONNECTION_STRING} up


db/migrations/down: # @HELP run migrations to teardown database tables
db/migrations/down:
	migrate -path ./migrations -database ${DB_CONNECTION_STRING} down


db/migrations/refresh: # @HELP refresh database tables
db/migrations/refresh: db/migrations/down db/migrations/up
