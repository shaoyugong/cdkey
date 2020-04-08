PROTO_PATH="api/proto"
SAVE_PATH="pkg"

for file in ${PROTO_PATH}/*.proto
do
  protoc --gofast_out=plugins=grpc:${SAVE_PATH} ${file}
done