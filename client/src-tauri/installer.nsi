; 不要包含 tauri.nsi，直接在默认安装的后续步骤中追加或执行
Section "MainSection" SEC01
    ; 在桌面上创建中文快捷方式
    CreateShortCut "$DESKTOP\幻想镇.lnk" "$INSTDIR\hxzl.exe" "" "$INSTDIR\hxzl.exe" 0
    
    ; 在开始菜单创建中文快捷方式
    CreateDirectory "$SMPROGRAMS\幻想镇"
    CreateShortCut "$SMPROGRAMS\幻想镇\幻想镇.lnk" "$INSTDIR\hxzl.exe" "" "$INSTDIR\hxzl.exe" 0
SectionEnd