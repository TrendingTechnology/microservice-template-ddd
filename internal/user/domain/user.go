//go:generate protoc -I/usr/local/include -I. --go_out=Minternal/user/domain/user.proto=.:. --go_opt=paths=source_relative user.proto

package domain
