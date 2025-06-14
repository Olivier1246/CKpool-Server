###### Installation ######

go build -o ckpool-server *.go

./ckpool-server

http://0.0.0.0:9000
http://localhost:9000
http://0.0.0.0:9000/api/data


###### Update ######

cd /path/to/ckpool-server

git reset --hard
git pull

go build -o ckpool-server *.go


###### Use screen ######

screen -dmS CKpool-server ./ckpool-server