# Cherryfs: Protobuf

To update or generate pb.go files, go to the directory of the proto file, execute the command
below:

``` protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative ./<filename>?.proto ```
