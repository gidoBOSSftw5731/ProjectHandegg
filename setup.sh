#!/bin/sh

echo "Please have already installed postgres and have it working with hashed passwords for local users kthx"

if [ -f .env ]
then
    echo ".env exists, please rename or remove."
    exit 1
fi

password=`openssl rand -base64 48`
name="handegg"
user="handegg_usr"

sudo -u postgres createdb $name
sudo -u postgres createuser $user

sudo -u postgres psql -c "alter user ${user} with encrypted password '${password}';"
sudo -u postgres psql -c "grant all privileges on database ${name} to  ${user};"

echo "DBUSER=${user}">.env
echo "DBPASS=${password}">>.env
echo "DBPNAME=${name}">>.env