@echo off
:: ====================================================================
:: 【配置区域】所有配置项全部集中在此处，按需修改：
:: 注意：set 变量时等号左右两边千万不要留空格！
:: ====================================================================

:: 1. 设置项目所在的子目录名称（如果批处理与项目同级，可直接留空）
set PROJECT_DIR=client

:: 2. 设置构建/运行模式:
::    - spa      : 纯网页单页应用
::    - ssr      : 服务端渲染 (SEO 优化)
::    - electron : 桌面端应用
set BUILD_MODE=electron

:: 3. 设置执行动作:
::    - dev      : 启动开发服务器（热重载）
::    - build    : 生产环境打包构建
set ACTION=dev

:: ====================================================================
:: 以下为脚本核心逻辑，无需修改
:: ====================================================================

:: 设置编码为 UTF-8，防止中文路径或提示乱码
chcp 65001 >nul

:: 切换到当前批处理文件所在的目录
cd /d %~dp0

:: 进入指定的项目目录
if "%PROJECT_DIR%"=="" goto :RUN_TASK

if exist "%PROJECT_DIR%" (
    cd %PROJECT_DIR%
    echo [信息] 已成功切换至目录: %CD%
) else (
    :: 【修改点】找不到目录时直接打印错误，并跳转到结束点中止运行
    echo [错误] 未找到指定的项目目录 "%PROJECT_DIR%"，脚本终止运行！
    goto :END
)

:RUN_TASK
echo ========================================
echo 正在执行: Quasar [%ACTION%] 模式 [%BUILD_MODE%]
echo ========================================

:: 使用标签跳转替代嵌套 if，完美避开 Windows 批处理的变量预解析坑
if /i "%ACTION%"=="dev" goto :DO_DEV
if /i "%ACTION%"=="build" goto :DO_BUILD

:: 如果都不是，则报错
echo [错误] 未知的 ACTION 参数 [%ACTION%]，请检查脚本配置！
goto :END

:DO_DEV
echo 正在启动开发服务器...
call pnpm quasar dev -m %BUILD_MODE%
goto :END

:DO_BUILD
echo 正在开始生产环境打包...
if /i "%BUILD_MODE%"=="electron" (
    echo [配置] 已启用仅构建 Windows 平台目标 (Win64)
    call pnpm quasar build -m electron -w
) else (
    call pnpm quasar build -m %BUILD_MODE%
)
goto :END

:END
echo ========================================
echo 任务已结束。
echo ========================================
pause