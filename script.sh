#!/bin/bash


# commandExists() {
#   command -v "$1" >/dev/null 2>&1
# }

# if ! commandExists "mysql"; then
#   echo "
#   mysql is not installed on your system.
#   Please install mysql and try again.
#   "
#   exit 1
# fi

# if ! commandExists "go"; then
#   echo "
#   go is not installed on your system.
#   Please install go and try again.
#   "
#   exit 1
# fi
# if commandExists "open"; then
#   open http://localhost:8000
# else
#   echo "Couldn't automatically open the website in the browser. You can manually open http://localhost:8000"
# fi


# # Set Up database settings
# # MySQL Database Information

# read -p "Enter MySQL username: " DB_USER
# read -s -p "Enter MySQL password: " DB_PASSWORD
# read -p "Enter host name: " DB_HOST
# read -s -p "Enter MySQL Databse Name: " DB_NAME


# cat << EOF > config.yaml
# DB_USERNAME: $DB_USER
# DB_PASSWORD: $DB_PASSWORD
# DB_HOST: $DB_HOST
# DB_NAME: $DB_NAME
# EOF
# echo "config.yaml created with the variables."

# # Command to run migrate files
# migrate -path migration/ -database "mysql://DB_USERNAME:DB_PASSWORD@tcp(DB_HOST:3306)/DB_NAME" -verbose up

# # Check the exit code of the mysql command
# if [ $? -eq 0 ]; then
#     echo "Databse created successfully."
# else
#     echo "Error importing dump file."
# fi


go build -o mvc ./cmd/main.go

./mvc