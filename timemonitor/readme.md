
### Q: 限制魔百盒每天使用时间，到时间自动关机

1. 魔百盒用ttl线连接电脑
    ```
    https://www.znds.com/tv-1227500-1-1.html
    ```
2. 执行命令 "cat /proc/cpuinfo"，显示处理器为kunlun的ARMv7架构
    ```
    processor       : 0
    Processor       : ARMv7 Processor rev 0 (v7l)
    BogoMIPS        : 24.00
    Features        : half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm
    CPU implementer : 0x41
    CPU architecture: 7
    CPU variant     : 0x2
    CPU part        : 0xd05
    CPU revision    : 0
    
    processor       : 1
    Processor       : ARMv7 Processor rev 0 (v7l)
    BogoMIPS        : 24.00
    Features        : half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm
    CPU implementer : 0x41
    CPU architecture: 7
    CPU variant     : 0x2
    CPU part        : 0xd05
    CPU revision    : 0
    
    processor       : 2
    Processor       : ARMv7 Processor rev 0 (v7l)
    BogoMIPS        : 24.00
    Features        : half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm
    CPU implementer : 0x41
    CPU architecture: 7
    CPU variant     : 0x2
    CPU part        : 0xd05
    CPU revision    : 0
    
    processor       : 3
    Processor       : ARMv7 Processor rev 0 (v7l)
    BogoMIPS        : 24.00
    Features        : half thumb fastmult vfp edsp neon vfpv3 tls vfpv4 idiva idivt vfpd32 lpae evtstrm
    CPU implementer : 0x41
    CPU architecture: 7
    CPU variant     : 0x2
    CPU part        : 0xd05
    CPU revision    : 0
    
    Hardware        : kunlun
    Revision        : 0000
    Serial          : 0000000000000000
    ```
3. 修改build-linux.cmd的linux目标架构并执行，编辑为linux可执行程序
4. 执行 "mount -rw -o remount /system" 重新挂载/system为可读写 
5. 将编译后的"time_monitor"文件复制到机顶盒 "/system/bin/timeMonitor.out" 下，
   添加可执行权限 chmod +x timeMonitor.out
6. 设置开机自启动脚本，修改 "/system/etc/init.kunlun.sh" 和 "/system/etc/init.sunniwell.sh" 文件,在末尾加上
    ```
    # 限制魔百盒每天使用时间，到时间自动关机
    /system/bin/timeMonitor.out
    ```