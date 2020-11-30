# ProjectHandegg
Sports in wordpress, what could go wrong?


## Setup

Run setup.sh AFTER installing postgresql and configuring it to use hashed passwords for local users that are not "postgres". The script will run sudo as user postgres, if you do not have access to that unix account at time of installation, please either gain access or edit the script to log in using a set of postgres creds you have

This script will create a .env file with the credentials for the program. Do not remove this file unless you either know what you're doing or need to reset the app. Removing the .env file alone will not delete any data from the database (therefore it is not enough to destroy the entire configuration, only the username and password used by this script)