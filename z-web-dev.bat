@echo off
chcp 65001 >nul

:: 切换到当前批处理文件所在的目录
cd web

:: 使用 call 强制批处理等待命令执行，防止直接闪退
call pnpm quasar dev -m spa

:: 如果你以后想改用 SSR 或 Electron，只需在这里换成对应的命令即可：
:: call pnpm quasar dev -m spa
:: call pnpm quasar dev -m ssr
:: call pnpm quasar dev -m electron

pause