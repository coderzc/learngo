file_path=$(cd `dirname $0`; pwd)
dir=$(ls $file_path/../protos) #定义遍历的目录
for i in $dir; do
  protoc -I=$file_path/../protos --go_out=plugins=grpc:$file_path/../src $file_path/../protos/$i
done
cd
