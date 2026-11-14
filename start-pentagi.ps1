#requires -Version 7.0
<#
.SYNOPSIS
    PentAGI 繁體中文漢化版 · Windows 一鍵啟動器

.DESCRIPTION
    檢查環境（Docker／.env／連接埠）→ 必要時建立 .env → 啟動 docker compose → 開啟瀏覽器。
    全程以繁體中文提示。

.EXAMPLE
    pwsh -File .\start-pentagi.ps1
    pwsh -File .\start-pentagi.ps1 -Down      # 停止服務
    pwsh -File .\start-pentagi.ps1 -Logs      # 追蹤日誌
#>
[CmdletBinding()]
param(
    [switch]$Down,
    [switch]$Logs,
    [switch]$SkipChecks
)

$ErrorActionPreference = 'Stop'
$root = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $root

function Write-Title([string]$text) {
    Write-Host ''
    Write-Host ('─' * 62) -ForegroundColor DarkGray
    Write-Host "  $text" -ForegroundColor Cyan
    Write-Host ('─' * 62) -ForegroundColor DarkGray
}

function Write-Step([string]$text) { Write-Host "  → $text" -ForegroundColor Gray }
function Write-Ok([string]$text) { Write-Host "  [OK] $text" -ForegroundColor Green }
function Write-Warn([string]$text) { Write-Host "  [!]  $text" -ForegroundColor Yellow }
function Write-Err([string]$text) { Write-Host "  [X]  $text" -ForegroundColor Red }

Write-Title 'PentAGI 漢化版 · 啟動器'

# 1) Docker 檢查
if (-not $SkipChecks) {
    Write-Step '檢查 Docker'
    $docker = Get-Command docker -ErrorAction SilentlyContinue
    if (-not $docker) {
        Write-Err '找不到 docker 指令。請先安裝 Docker Desktop：https://docs.docker.com/get-docker/'
        exit 1
    }
    try {
        docker info *> $null
        Write-Ok 'Docker 正在執行'
    } catch {
        Write-Err 'Docker 服務未啟動，請開啟 Docker Desktop 後重試。'
        exit 1
    }
    docker compose version *> $null
    if ($LASTEXITCODE -ne 0) {
        Write-Err 'Docker Compose 外掛不可用（需要 docker compose v2）。'
        exit 1
    }
    Write-Ok 'Docker Compose 可用'
}

# 2) 停止 / 日誌模式
if ($Down) {
    Write-Step '停止服務'
    docker compose down
    Write-Ok '已停止'
    exit 0
}

if ($Logs) {
    Write-Step '追蹤日誌（Ctrl+C 結束）'
    docker compose logs -f
    exit 0
}

# 3) .env 檢查
Write-Step '檢查 .env'
if (-not (Test-Path '.env')) {
    if (Test-Path '.env.example') {
        Copy-Item '.env.example' '.env'
        Write-Warn '已從 .env.example 建立 .env —— 請先填入至少一個 LLM 供應商的 API Key 再繼續。'
        Write-Host ''
        Write-Host '  建議最小設定（取消註解並填入）：' -ForegroundColor Yellow
        Write-Host '    OPEN_AI_KEY=sk-...' -ForegroundColor DarkGray
        Write-Host '    ANTHROPIC_API_KEY=sk-ant-...' -ForegroundColor DarkGray
        Write-Host '    OLLAMA_SERVER_URL=http://ollama-server:11434   # 本機模型' -ForegroundColor DarkGray
        Write-Host ''
        Write-Host "  編輯完成後重新執行：pwsh -File .\$($MyInvocation.MyCommand.Name)" -ForegroundColor Yellow
        exit 0
    }
    Write-Err '找不到 .env 也找不到 .env.example。'
    exit 1
}
Write-Ok '.env 存在'

# 4) 啟動
Write-Step '啟動容器（第一次會下載映像，可能需要數分鐘）'
docker compose up -d
if ($LASTEXITCODE -ne 0) {
    Write-Err '啟動失敗，請執行：docker compose logs --tail 50'
    exit 1
}

Write-Step '等待服務就緒'
$deadline = (Get-Date).AddMinutes(3)
$ready = $false
while ((Get-Date) -lt $deadline) {
    $state = (docker compose ps --format '{{.Service}} {{.State}}' 2>$null) -join "`n"
    if ($state -match 'pentagi\s+running') {
        $ready = $true
        break
    }
    Start-Sleep -Seconds 3
}
if ($ready) { Write-Ok 'PentAGI 已啟動' } else { Write-Warn '尚未偵測到 pentagi 服務就緒，可稍後以 -Logs 檢視' }

# 5) 開啟瀏覽器
$url = 'https://localhost:8443'
Write-Host ''
Write-Host "  介面網址：$url" -ForegroundColor Cyan
Write-Host '  （自簽憑證，瀏覽器會提示風險，選擇「繼續前往」即可）' -ForegroundColor DarkGray
Write-Host '  停止服務：pwsh -File .\start-pentagi.ps1 -Down' -ForegroundColor DarkGray
Write-Host '  查看日誌：pwsh -File .\start-pentagi.ps1 -Logs' -ForegroundColor DarkGray
Start-Process $url
