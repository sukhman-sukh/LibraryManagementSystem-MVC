#!/bin/bash



go build -o mvc ./cmd/main.go

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
# DB_HOST="localhost"
# DB_NAME="LMS"
# read -p "Enter MySQL username: " DB_USER
# read -s -p "Enter MySQL password: " DB_PASSWORD
# DUMP_FILE="./schema/dump.sql"


# cat << EOF > data.yaml
# DB_USERNAME: $DB_USER
# DB_PASSWORD: $DB_PASSWORD
# DB_HOST: $DB_HOST
# DB_NAME: $DB_NAME
# EOF
# echo "data.yaml created with the variables."

# # Command to import the dump file
# mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD $DB_NAME < $DUMP_FILE
# # Check the exit code of the mysql command
# if [ $? -eq 0 ]; then
#     echo "Dump file imported successfully."
# else
#     echo "Error importing dump file."
# fi


./mvc