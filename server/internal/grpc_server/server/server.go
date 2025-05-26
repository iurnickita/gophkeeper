// Пакет grpcserver. Обработчики grpc
package grpcserver

// Установка grpc
// https://grpc.io/docs/languages/go/quickstart/

// Генерация go-файлов для grpc сервиса
// cd internal/grpc_server
// protoc --go_out=. --go_opt=paths=source_relative \
// --go-grpc_out=. --go-grpc_opt=paths=source_relative \
// proto/server.proto
