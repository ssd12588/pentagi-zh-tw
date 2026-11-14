#requires -Version 5.1
<#
.SYNOPSIS
    PentAGI 繁體中文版 · 互動式啟動選單（由 start-pentagi.bat 呼叫）
#>
[CmdletBinding()]
param(
    [string]$Choice = '',
    [switch]$Once
)

$ErrorActionPreference = 'Continue'
$root = Split-Path -Parent (Split-Path -Parent (Split-Path -Parent $MyInvocation.MyCommand.Path))
Set-Location $root
$script = Join-Path $root 'start-pentagi.ps1'

function Show-Menu {
    Clear-Host
    Write-Host '================================================================' -ForegroundColor DarkCyan
    Write-Host '   PentAGI 繁體中文版 · 啟動器' -ForegroundColor Cyan
    Write-Host '================================================================' -ForegroundColor DarkCyan
    Write-Host ''
    Write-Host '    [1] 啟動           使用已建置的中文映像啟動'
    Write-Host '    [2] 建置並啟動     第一次使用、或更新程式碼後必用'
    Write-Host '    [3] 查看日誌       追蹤容器輸出（Ctrl+C 返回）'
    Write-Host '    [4] 停止服務'
    Write-Host '    [5] 開啟網頁介面   https://localhost:8443'
    Write-Host '    [6] 編輯 .env      設定 API 金鑰等'
    Write-Host '    [7] 檢查環境       Docker / .env / 容器狀態'
    Write-Host ''
    Write-Host '    [0] 離開'
    Write-Host '================================================================' -ForegroundColor DarkCyan
}

function Invoke-Compose([string[]]$Arguments) {
    & docker compose @Arguments
}

function Test-Environment {
    Clear-Host
    Write-Host '================ 環境檢查 ================' -ForegroundColor Cyan
    Write-Host ''
    Write-Host '[1] Docker'
    $docker = Get-Command docker -ErrorAction SilentlyContinue
    if (-not $docker) {
        Write-Host '    X 找不到 docker 指令，請先安裝 Docker Desktop' -ForegroundColor Red
    } else {
        docker info *> $null
        if ($LASTEXITCODE -ne 0) {
            Write-Host '    X Docker 已安裝但服務未啟動（請開啟 Docker Desktop）' -ForegroundColor Red
        } else {
            Write-Host '    OK Docker 正在執行' -ForegroundColor Green
        }
        Write-Host ''
        Write-Host '[2] Docker Compose'
        docker compose version *> $null
        if ($LASTEXITCODE -ne 0) {
            Write-Host '    X 無法使用 docker compose v2' -ForegroundColor Red
        } else {
            Write-Host '    OK docker compose 可用' -ForegroundColor Green
        }
    }

    Write-Host ''
    Write-Host '[3] 設定檔 .env'
    $envPath = Join-Path $root '.env'
    if (-not (Test-Path $envPath)) {
        Write-Host '    ! 沒有 .env，請用選項 [6] 建立並填入 API Key' -ForegroundColor Yellow
    } else {
        Write-Host '    OK .env 存在' -ForegroundColor Green
        $lines = Get-Content $envPath -ErrorAction SilentlyContinue
        $keys = @('DEEPSEEK_API_KEY', 'OPEN_AI_KEY', 'ANTHROPIC_API_KEY', 'OLLAMA_SERVER_URL')
        $filled = @()
        foreach ($key in $keys) {
            $hit = $lines | Where-Object { $_ -match "^$key=(.+)$" } | Select-Object -First 1
            if ($hit) { $filled += $key }
        }
        if ($filled.Count -eq 0) {
            Write-Host '    ! 尚未填入任何 LLM 供應商設定' -ForegroundColor Yellow
        } else {
            Write-Host ('    OK 已設定：' + ($filled -join '、')) -ForegroundColor Green
        }
        $emb = $lines | Where-Object { $_ -match '^EMBEDDING_PROVIDER=' } | Select-Object -First 1
        if ($emb) { Write-Host ('    嵌入供應商：' + $emb) -ForegroundColor DarkGray }
    }

    Write-Host ''
    Write-Host '[4] 本地中文映像'
    $image = 'pentagi-zh-tw:latest'
    if (Test-Path $envPath) {
        $imageLine = (Get-Content $envPath | Where-Object { $_ -match '^PENTAGI_IMAGE=(.+)$' } | Select-Object -First 1)
        if ($imageLine) { $image = ($imageLine -split '=', 2)[1].Trim() }
    }
    docker image inspect $image *> $null
    if ($LASTEXITCODE -ne 0) {
        Write-Host "    ! 找不到映像 $image（請用選項 [2] 建置）" -ForegroundColor Yellow
    } else {
        Write-Host "    OK 映像已存在：$image" -ForegroundColor Green
    }

    Write-Host ''
    Write-Host '[5] 容器狀態'
    if (Get-Command docker -ErrorAction SilentlyContinue) {
        Invoke-Compose @('ps')
    }
    Write-Host ''
    if (-not $Once) {
        Read-Host '按 Enter 返回選單' | Out-Null
    }
}

function Start-WithPowerShell {
    param([string[]]$ExtraArgs)
    & pwsh -NoProfile -ExecutionPolicy Bypass -File $script @ExtraArgs 2>$null
    if ($LASTEXITCODE -ne 0 -and $null -ne $LASTEXITCODE) {
        & powershell -NoProfile -ExecutionPolicy Bypass -File $script @ExtraArgs
    }
    Write-Host ''
    Read-Host '按 Enter 返回選單' | Out-Null
}

while ($true) {
    if ($Choice) {
        $input = $Choice
        $Choice = ''
    } else {
        Show-Menu
        $input = Read-Host '請輸入選項後按 Enter'
    }

    switch ($input.Trim()) {
        '1' { Start-WithPowerShell @() }
        '2' { Start-WithPowerShell @('-Build') }
        '3' { Start-WithPowerShell @('-Logs') }
        '4' { Start-WithPowerShell @('-Down') }
        '5' { Start-Process 'https://localhost:8443' }
        '6' {
            $envPath = Join-Path $root '.env'
            if (-not (Test-Path $envPath)) {
                Copy-Item (Join-Path $root '.env.example') $envPath
                Write-Host '已從 .env.example 建立 .env，請填入至少一個 LLM 供應商的 API Key。' -ForegroundColor Yellow
            }
            Start-Process notepad.exe $envPath
        }
        '7' { Test-Environment }
        '0' { Write-Host '再見。' -ForegroundColor DarkGray; return }
        default { Write-Host '請輸入 0-7 的數字。' -ForegroundColor Yellow; Start-Sleep -Seconds 1 }
    }

    if ($Once) {
        return
    }
}
