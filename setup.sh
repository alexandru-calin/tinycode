#!/bin/bash
set -eu

MIGRATOR_USER=migrator
WEB_USER=web
HOST=localhost

# Create Database
echo "Creating DB..."
read -p "Enter DB name: " DB_NAME

sudo mysql -e "CREATE DATABASE IF NOT EXISTS $DB_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Create Migrator user
echo "Creating Migrator user for $DB_NAME..."
read -p "Enter password for Migrator user: " PASSWORD

sudo mysql -e "CREATE USER IF NOT EXISTS '$MIGRATOR_USER'@'$HOST' IDENTIFIED BY '$PASSWORD'";
sudo mysql -e "GRANT CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, INDEX, LOCK TABLES ON $DB_NAME.* TO '$MIGRATOR_USER'@'$HOST'";

echo "export MIGRATOR_DB_DSN='mysql://migrator:$PASSWORD@tcp($HOST)/$DB_NAME?parseTime=true'" > .envrc

# Create Web user
echo "Creating Web user for $DB_NAME..."
read -p "Enter password for Web user: " PASSWORD

sudo mysql -e "CREATE USER IF NOT EXISTS '$WEB_USER'@'$HOST' IDENTIFIED BY '$PASSWORD';"
sudo mysql -e "GRANT SELECT, INSERT, UPDATE, DELETE ON $DB_NAME.* TO '$WEB_USER'@'$HOST';"

echo "export WEB_DB_DSN='$WEB_USER:$PASSWORD@tcp($HOST)/$DB_NAME?parseTime=true'" >> .envrc
