@echo off
protoc --go_out=.\src .\proto\*.proto
pause