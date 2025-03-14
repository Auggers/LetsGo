#!/bin/bash
set -e

mongosh admin --eval '
db.createRole({
  role: "testwriters",
  privileges: [],
  roles: [ "userAdminAnyDatabase", "clusterMonitor", "clusterManager", "clusterAdmin" ]
})
db.createRole({
  role: "testreaders",
  privileges: [],
  roles: [ "read" ]
})
db.createRole({
  role: "otherreaders",
  privileges: [],
  roles: [ "readAnyDatabase" ]
})
db.createRole({
  role: "otherwriters",
  privileges: [],
  roles: [ "readWriteAnyDatabase" ]
})
db.createRole({
  role: "testusers",
  privileges: [],
  roles: [ "readWrite" ]
})
db.createRole({
  role: "otherusers",
  privileges: [],
  roles: [ "read" ]
})
'

wget https://raw.githubusercontent.com/percona/percona-server-mongodb/master/support-files/ldap-sasl/ldap/deploy_openldap.sh
wget https://raw.githubusercontent.com/percona/percona-server-mongodb/master/support-files/ldap-sasl/ldap/generate_users_ldif.sh
wget https://raw.githubusercontent.com/percona/percona-server-mongodb/master/support-files/ldap-sasl/ldap/groups.ldif
wget https://raw.githubusercontent.com/percona/percona-server-mongodb/master/support-files/ldap-sasl/ldap/users.ldif
wget https://raw.githubusercontent.com/percona/percona-server-mongodb/master/support-files/ldap-sasl/settings.conf

sudo sed -i 's/..\/settings.conf/settings.conf/g' deploy_openldap.sh && sudo chmod +x deploy_openldap.sh | true
sudo sed -i 's/..\/settings.conf/settings.conf/g' generate_users_ldif.sh && sudo chmod +x generate_users_ldif.sh | true

# Confirm that slapd does not exist
output=$(sudo systemctl status slapd || true)

if systemctl is-active --quiet slapd; then
  echo "slapd is running. Stopping and removing it..."
  sudo systemctl stop slapd
  sudo apt-get purge -y slapd ldap-utils
  sudo rm -rf /etc/ldap /var/lib/ldap /var/run/slapd /var/log/slapd.log
  sudo deluser openldap || true
  sudo delgroup openldap || true
  sudo apt-get autoremove -y
  sudo apt-get autoclean
else
  echo "slapd is not running. Skipping removal."
fi

sleep 5

sudo ./deploy_openldap.sh

sudo sed -i '/^#security/a \
security:\n  authorization: "enabled"\n  ldap:\n    servers: "localhost"\n    transportSecurity: none\n    authz:\n      queryTemplate: "dc=percona,dc=com?dn?sub?(&(objectClass=groupOfNames)(member={PROVIDED_USER}))"' /etc/mongod.conf

sudo sed -i '/^#setParameter:/a \
setParameter:\n  authenticationMechanisms: "PLAIN"' /etc/mongod.conf

sudo systemctl restart mongod


