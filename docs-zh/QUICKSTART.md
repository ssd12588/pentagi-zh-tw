# 快速開始（繁體中文）

從零到第一次登入，約 10 至 15 分鐘（不含映像下載時間）。

## 0. 你需要什麼

| 項目 | 需求 |
| --- | --- |
| Docker | Docker + Docker Compose（或 Podman） |
| CPU | 2 vCPU 以上 |
| 記憶體 | 4 GB 以上 |
| 磁碟 | 20 GB 可用空間 |
| 網路 | 可下載映像（第一次啟動會拉取數個映像） |
| LLM | 至少一個供應商的 API Key（OpenAI／Anthropic／Gemini／Bedrock／DeepSeek／GLM／Qwen／Ollama…） |

> 只想先在本機試玩：用 Ollama 跑本機模型即可，不必申請雲端 Key。

## 1. 安裝方式（擇一）

### 方式 A：互動式安裝程式（推薦）

官方安裝程式有終端機介面，會幫你檢查環境、產生 `.env`、設定供應商與 SSL。

- Windows (amd64)：<https://pentagi.com/downloads/windows/amd64/installer-latest.zip>
- Linux (amd64)：<https://pentagi.com/downloads/linux/amd64/installer-latest.zip>
- Linux (arm64)：<https://pentagi.com/downloads/linux/arm64/installer-latest.zip>
- macOS (Intel／Apple Silicon)：見上游 README 的對應連結

```bash
# Linux 範例
mkdir -p pentagi && cd pentagi
wget -O installer.zip https://pentagi.com/downloads/linux/amd64/installer-latest.zip
unzip installer.zip
sudo ./installer        # 需要讀取 /var/run/docker.sock
```

安裝程式會處理：系統檢查 → 建立 `.env` → 設定 LLM 供應商 → 設定搜尋引擎 → 安全強化（憑證）→ 以 docker compose 啟動。

> **注意**：安裝程式的介面目前仍是英文（屬於 TUI 常數檔 `backend/cmd/installer/wizard/locale/locale.go`，尚未漢化）。你可以照著畫面上的選項一路按 Enter 使用預設值。

### 方式 B：手動 docker compose

```bash
git clone <你的漢化版倉庫> pentagi-zh
cd pentagi-zh
cp .env.example .env
# 編輯 .env：至少填入一個 LLM 供應商的 Key
docker compose up -d
```

## 2. 最小 `.env` 設定

```dotenv
# 至少要有一個 LLM 供應商
OPEN_AI_KEY=sk-...

# 或 Anthropic
ANTHROPIC_API_KEY=sk-ant-...

# 或本機 Ollama
OLLAMA_SERVER_URL=http://ollama-server:11434

# 搜尋引擎（可選，但建議至少設一個）
TAVILY_API_KEY=tvly-...
```

其餘設定（觀測、Langfuse、Graphiti、工具參數）可先留預設。

## 3. 啟動與登入

```bash
docker compose ps          # 確認服務已啟動
docker compose logs -f pentagi   # 觀察啟動日誌
```

- Web 介面：<https://localhost:8443>（自簽憑證，瀏覽器會提示風險，選「繼續前往」）
- 第一次啟動會建立管理員帳號；若忘了密碼，見 `TROUBLESHOOTING.md`。

登入後你會看到（皆為繁體中文）：

| 位置 | 用途 |
| --- | --- |
| 儀表板 | 用量、成本、Token 趨勢 |
| 流程 | 建立／管理每一次滲透測試任務 |
| 知識庫 | 上傳文件作為向量知識，讓代理參考 |
| 範本 | 任務範本（滲透測試情境預設集） |
| 設定 → 供應商 | 新增／測試 LLM 供應商與模型 |
| 設定 → 提示詞 | 檢視系統／使用者提示詞（**不建議翻譯，會影響代理判斷**） |
| 設定 → API Token | 產生 Token 以便用程式呼叫 API |

## 4. 建立第一個流程

1. 左側選 **流程** → 右上 **新增流程**。
2. 填寫 **流程標題**，在 **描述任務** 欄位用中文描述目標，例如：
   > 對 example.com 進行外部滲透測試，重點檢查登入頁的 SQL 注入與 XSS。
3. 選一個 **供應商／模型**（建議挑推理能力較強的模型）。
4. 按 **建立**，代理會開始規劃任務；右側 **助理** 分頁可以看到它的回報（中文），**終端機** 分頁是實際執行的指令（英文／原始輸出，屬正常現象）。

## 5. 驗證漢化是否生效

```bash
cd frontend
pnpm exec tsc -b          # 型別檢查
pnpm run build            # 建置（會產生 dist/）
```

若你自己用 Docker 建置映像，`Dockerfile` 會在 Stage 1 一起建置前端，成品已含中文介面。
