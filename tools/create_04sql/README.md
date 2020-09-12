

echo 'select id,features from estate' |  mysql -u isucon -p isuumo > estate_features.tsv
echo 'select id,features from chair' |  mysql -u isucon -p isuumo > chair_features.tsv

go run main.go > ../../webapp/mysql/db/3_DummyFeatures.sql