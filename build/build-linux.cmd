@echo off

echo start building...

cd ..

    :: 魔百盒CM311-5 "cat /proc/cpuinfo" 为armv7架构
    :: arm架构
    SET GOARCH=arm
    SET CGO_ENABLED=0
    SET GOOS=linux

    :: 魔百盒时间同步工具
    cd ./timesync
    echo ^> build [linux arm]: time_sync
    go build -o ../bin/time_sync
    cd ..

    :: 魔百盒每天使用时长监控
    cd ./timemonitor
    echo ^> build [linux arm]: time_monitor
    go build -o ../bin/time_monitor
    cd ..

cd ./build
echo complete

:: 注：
::
:: 1. 如果编译失败，提示引入包报错，执行"go mod tidy"指令进行依赖的下载
::
:: 2. 如果下载依赖超时，修改 go mod 代理
::
:: 3. -ldflags "-H windowsgui" 不显示windows黑窗口,-w关闭DWARF调试信息,-s关闭Go符号表的生成