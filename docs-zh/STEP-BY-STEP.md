# PentAGI 繁體中文版 · 從零開始一步一步做

> 適用：Windows 10/11、macOS、Linux。全程以繁體中文介面操作。
> 假設你完全沒用過 PentAGI，只要照著做即可。

---

## 步驟 0：先確認你電腦的條件

| 項目 | 最低需求 | 怎麼確認 |
| --- | --- | --- |
| Docker Desktop | 已安裝且**正在執行** | 開 PowerShell 打 `docker info`，有輸出且沒報錯 |
| CPU | 2 核心以上 | 工作管理員 → 效能 |
| 記憶體 | 4 GB 以上（建議 8 GB） | 工作管理員 → 效能 |
| 磁碟 | 20 GB 可用 | 本機磁碟 C: |
| LLM API Key | 至少一個 | 沒有也能先跑，但代理無法思考 |

> 沒有 Docker？先裝：<https://docs.docker.com/get-docker/>（Windows 請裝 Docker Desktop 並開啟 WSL2）。

---

## 步驟 1：把專案 clone 下來

打開 PowerShell（或終端機），找一個你放專案的位置：

```powershell
cd C:\Users\你的帳號\Documents

git clone https://github.com/ssd12588/pentagi-zh-tw.git pentagi-zh-tw

cd pentagi-zh-tw
```

確認抓到的是漢化版：

```powershell
git log --oneline -3
```

你應該看到含 `zh-TW` 的 commit 訊息。

---

## 步驟 2：建立 `.env` 設定檔

```powershell
Copy-Item .env.example .env
```

用記事本（或 VS Code）打開 `.env`，**至少要填一個 LLM 供應商的金鑰**。挑一個你有的：

```dotenv
# 有 OpenAI
OPEN_AI_KEY=sk-...

# 有 Anthropic
ANTHROPIC_API_KEY=sk-ant-...

# 想用本機模型（免費）
OLLAMA_SERVER_URL=http://ollama-server:11434
```

選填但建議（搜尋能力）：

```dotenv
TAVILY_API_KEY=tvly-...
```

> 不確定要填哪個？先跳過，啟動後再到「設定 → 供應商」用介面新增也可以。

---

## 步驟 3：啟動（第一次要用「建置」）

⚠️ **關鍵**：`docker-compose.yml` 預設抓的是**上游英文映像**，一定要用本專案原始碼建置，中文介面才會出現。

**最簡單：直接雙擊 `start-pentagi.bat`** → 選單按 `2`（建置並啟動）。

或用命令列：

```powershell
pwsh -File .\start-pentagi.ps1 -Build
```

腳本會依序做：

1. 檢查 Docker 是否在跑
2. 檢查 `.env` 是否存在（沒有會幫你從範本建立並停下，提醒你填 Key）
3. 建置本地映像 `pentagi-zh-tw:latest`（**第一次約 5～15 分鐘**，會看到大量建置訊息，正常）
4. 啟動容器
5. 開啟瀏覽器到 <https://localhost:8443>

> 如果你偏好手動指令：
>
> ```powershell
> docker compose -f docker-compose.yml -f docker-compose.build.yml up -d --build
> ```

---

## 步驟 4：開瀏覽器登入

打開 <https://localhost:8443>

- 會出現「憑證不受信任」警告 → 點 **進階** → **繼續前往 localhost（不安全）**
  （自簽憑證是正常的，只影響本機）
- 用 `.env` 裡設定的管理員帳密登入
- 第一次登入可能要求你改密碼：需 16 字元以上，或 8-15 字元且含大小寫、數字、`!@#$&*` 其中之一

登入後你應該看到**繁體中文**的側邊欄：儀表板、流程、知識庫、範本、資源、設定。

---

## 步驟 5：設定 LLM 供應商（若 `.env` 沒填）

1. 左側 **設定** → **供應商**
2. 右上 **新增供應商**
3. 選擇類型（例如 OpenAI）→ 填入 **API 金鑰** → **名稱**（自己取，如 `openai-main`）
4. 按 **測試** 確認可用（會顯示測試結果）
5. 按 **建立供應商**

再依同樣方式設定要用於「摘要」「嵌入」的模型（若你的方案需要）。

---

## 步驟 6：建立第一個流程

1. 左側 **流程** → 右上 **新增流程**
2. **流程標題**：例如 `example.com 外部滲透測試`
3. **描述任務**（用中文寫，AI 看得懂）：
   > 對 example.com 進行外部滲透測試。重點檢查登入頁面的 SQL 注入與 XSS，並列出所有開放的高風險服務。只做被動與低噪音掃描，不要進行暴力破解。
4. 選擇 **供應商／模型**（建議挑推理能力強的，如 GPT-5、Claude Sonnet 等）
5. 按 **建立**

接著你會看到：

- **助理** 分頁：代理用中文回報進度與計畫
- **任務** 分頁：拆解出的任務與子任務
- **終端機** 分頁：實際執行的指令（英文／原始輸出，正常現象）
- **工具** 分頁：每次工具呼叫與結果
- **儀表板** 分頁：Token 用量與成本

---

## 步驟 7：看報告

流程結束後，到 **流程報告** 分頁：

- **網頁檢視**：直接在瀏覽器看
- **複製到剪貼簿**：貼到你的報告範本
- **下載 MD／PDF**：交付用

---

## 步驟 8：日常操作

| 想做的事 | 指令 |
| --- | --- |
| 啟動（之後都很快） | `pwsh -File .\start-pentagi.ps1` |
| 看日誌 | `pwsh -File .\start-pentagi.ps1 -Logs` |
| 停止 | `pwsh -File .\start-pentagi.ps1 -Down` |
| 更新到最新漢化版 | `git pull` 後再 `pwsh -File .\start-pentagi.ps1 -Build` |

---

## 疑難排解（對照你看到的中文訊息）

| 訊息／現象 | 原因 | 解法 |
| --- | --- | --- |
| 介面還是英文 | 用了上游英文映像 | 一定要 `-Build` 重新建置 |
| 找不到 Docker／Docker 服務未執行 | Docker Desktop 沒開 | 開啟 Docker Desktop，等它顯示 Running |
| `port is already allocated` | 8443 被佔用 | 改 `.env` 的 `PENTAGI_LISTEN_PORT=9443` |
| 登入後一片空白 | 後端還沒起來 | `pwsh -File .\start-pentagi.ps1 -Logs` 觀察，等 1-2 分鐘 |
| 無法連線到更新服務 | 對外網路受限 | 可忽略；更新狀態會顯示未知 |
| 沒有可用的供應商 | 沒設 LLM | 步驟 5 設定，或補 `.env` |
| 代理不動作／一直等待 | 模型不支援工具呼叫 | 換支援 function calling 的模型 |
| 容器未執行 | worker 映像缺失 | `docker compose up -d --force-recreate` |

---

## 進階：不同部署情境

| 情境 | 建議 |
| --- | --- |
| 只有一台電腦試玩 | 預設（內嵌 pgvector + 單機） |
| 團隊共用 | 設定 `PENTAGI_LISTEN_IP=0.0.0.0`、`CORS_ORIGINS`、反向代理 + 正式憑證 |
| 不想讓代理碰主機 | 使用雙節點（worker 隔離到另一台），見上游 `docker-compose.yml` 的 `WORKER_*` 設定 |
| 想省錢 | 用 Ollama 本機模型，或 DeepSeek／Qwen 等低價 API |
| 要知識圖譜 | 啟用 Graphiti（`docker-compose-graphiti.yml`） |

---

## 想自己維護漢化

```powershell
# 前端字串
node tools/i18n/extract.mjs
node tools/i18n/apply.mjs --list --out tools/i18n/pending.tsv
# 編輯 tools/i18n/dict/*.json 後
node tools/i18n/apply.mjs; node tools/i18n/apply.mjs --tests
cd frontend; pnpm exec tsc -b; pnpm run build

# 安裝程式字串
node tools/i18n/go-apply.mjs --list
# 編輯 tools/i18n/dict-go/*.json 後
node tools/i18n/go-apply.mjs; cd backend; go build ./cmd/installer/...

# 品質檢查
node tools/i18n/check-syntax.mjs
node tools/i18n/check-mixed.mjs
```
