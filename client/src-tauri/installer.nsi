; 包含 Tauri 默认的 NSIS 模板逻辑
!include "tauri.nsi"

; 在安装完成、准备创建快捷方式的阶段，重写或追加中文快捷方式
Section "MainSection" SEC01
    ; 如果默认脚本生成了英文快捷方式，我们可以直接在这里覆盖它，或者创建一个新的中文快捷方式
    ; $DESKTOP 是系统桌面路径，$INSTDIR 是你的安装目录
    CreateShortCut "$DESKTOP\幻想镇.lnk" "$INSTDIR\hxzl.exe" "" "$INSTDIR\hxzl.exe" 0
    
    ; 如果还需要在“开始菜单”创建中文快捷方式：
    CreateDirectory "$SMPROGRAMS\幻想镇"
    CreateShortCut "$SMPROGRAMS\幻想镇\幻想镇.lnk" "$INSTDIR\hxzl.exe" "" "$INSTDIR\hxzl.exe" 0
SectionEnd