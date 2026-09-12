@echo off
chcp 65001 >nul

:: 切换到 web 目录
cd web

:: 使用 call 执行 Electron 生产打包，并通过 -w 限制只构建 Windows 平台
call pnpm quasar build -m spa

:: 如果你想换成其他模式打包，可以改成下面对应的命令：
:: call pnpm quasar build -m electron -T win
:: call pnpm quasar build -m spa
:: call pnpm quasar build -m ssr

pause