#/bin/sh
# create temp dir
DOWNLOAD_DIR="$(mktemp -d)"

if [ -z $VERSION ]; then
    echo "VERSION is required"
    exit 1
fi

cd $DOWNLOAD_DIR
git clone https://github.com/volatiletech/sqlboiler
cd sqlboiler
git pull --tags
git checkout $VERSION
go install .
go install ./drivers/sqlboiler-psql
sqlboiler --version
