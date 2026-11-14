@echo off
rem ============================================================
rem  PentAGI zh-TW launcher (double-click me)
rem  ASCII-only on purpose: cmd.exe mangles UTF-8 inside .bat.
rem  All Chinese UI text lives in tools\launcher\menu.ps1
rem ============================================================
setlocal
cd /d "%~dp0"

set "MENU=%~dp0tools\launcher\menu.ps1"

if not exist "%MENU%" (
    echo [ERROR] tools\launcher\menu.ps1 not found. Please re-clone the repository.
    pause
    exit /b 1
)

where pwsh >nul 2>nul
if not errorlevel 1 (
    pwsh -NoProfile -ExecutionPolicy Bypass -File "%MENU%" %*
    goto done
)

where powershell >nul 2>nul
if not errorlevel 1 (
    powershell -NoProfile -ExecutionPolicy Bypass -File "%MENU%" %*
    goto done
)

echo [ERROR] PowerShell not found. Windows 10/11 ships with it; check your PATH.
pause
exit /b 1

:done
endlocal
exit /b 0
