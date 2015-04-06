echo "building ..."
rm -rf ./dbweb
mkdir ./dbweb
GOOS=linux GOARCH=amd64 go build -o ./dbweb/dbweb
cp -r ./static ./dbweb/
cp -r ./templates ./dbweb/
cp ./*.pem ./dbweb/
tar zcvf ./dbweb.tar.gz ./dbweb
echo "done."