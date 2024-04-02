set -e

mongosh <<EOF

use $MONGO_INITDB_DATABASE;

db.createUser({
    user: '$MONGODB_FLUENTD_USER',
    pwd: '$MONGODB_FLUENTD_PWD',
    roles: [
        {
            role: 'readWrite',
            db: '$MONGO_INITDB_DATABASE',
        },
    ],
})

EOF