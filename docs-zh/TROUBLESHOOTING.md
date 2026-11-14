# 疑難排解（中文錯誤訊息對照）

介面已是繁體中文，以下用「你看到的中文訊息」對照原因與解法。

## 登入與帳號

| 你看到的訊息 | 原因 | 解法 |
| --- | --- | --- |
| 登入失敗 | 帳號或密碼錯誤 | 確認 `.env` 中的管理員帳密；或見下方「忘記密碼」 |
| 工作階段已過期，請重新登入。 | Token 過期或被伺服器重啟清空 | 重新登入即可 |
| 需要變更密碼 / 繼續前需要先變更密碼。 | 首次登入或管理員重設後強制改密 | 到「設定 → 帳號」變更 |
| 密碼長度須超過 15 個字元，或至少 8 個字元且包含數字、小寫、大寫與特殊字元（!@#$&*） | 密碼強度不足 | 依訊息調整 |
| 兩次輸入的密碼不一致 | 確認密碼打錯 | 重新輸入 |

**忘記密碼**：以環境變數重設或直接更新資料庫中的使用者密碼雜湊，詳見上游 `backend/docs/`；最簡單的方式是停掉服務、修改 `.env` 中的管理員密碼變數後重啟。

## LLM 供應商

| 你看到的訊息 | 原因 | 解法 |
| --- | --- | --- |
| 尚未設定供應商 / 沒有可用的供應商 | 沒有可用的 LLM 設定 | 「設定 → 供應商」新增，或到 `.env` 補 Key |
| 載入供應商資料失敗 | 後端或資料庫未就緒 | `docker compose ps` 檢查 `pentagi`、`pgvector` 是否 healthy |
| 供應商測試結果：失敗 | Key 無效、模型名稱錯、或網路不通 | 用 curl 直接測 API；確認模型名稱（例如 `gpt-4o` 而非 `GPT-4o`） |
| 必須是有效的 JSON 物件 | 「額外 Body（JSON）」欄位格式錯誤 | 貼上合法 JSON，如 `{"chat_template_kwargs":{"enable_thinking":false}}` |
| 此模型僅支援自適應思考，無法停用。 | 該模型的推理參數限制 | 保持「自適應」或換模型 |

## 流程執行

| 你看到的訊息 | 原因 | 解法 |
| --- | --- | --- |
| 容器未執行 | Worker 容器被停止或映像缺失 | `docker compose logs worker`；必要時 `docker compose up -d --force-recreate` |
| 未收到更新——正在重新連線 | WebSocket 斷線（常見於反向代理逾時） | 檢查 nginx/Caddy 的 `proxy_read_timeout`，或直接連 `https://localhost:8443` |
| 沒有進行中的任務 / 找不到任務 | 代理尚未規劃，或流程已結束 | 看「助理」分頁的訊息；必要時重新建立流程 |
| 上傳失敗 / 附加失敗 | 檔案過大或 Worker 掛載目錄不可寫 | 單檔上限 300 MB、單次上傳上限 2 GB |
| 儲存為資源失敗 | 資源庫目錄權限不足 | 檢查 `docker-compose.yml` 中 `resources/` 的 volume 權限 |

## 安裝程式（英文 TUI）

安裝程式介面仍為英文（尚未漢化），常見對照：

| 英文 | 中文 |
| --- | --- |
| System Requirements Not Met | 系統需求未達成 |
| Docker Not Installed / Docker Daemon Not Running | 未安裝 Docker／Docker 服務未啟動 |
| Environment Configuration Missing | 找不到或無法讀取 `.env` |
| Write Permissions Required | 需要寫入權限 |
| Welcome to PentAGI | 歡迎畫面 |
| Maintenance → Update PentAGI | 維護 → 更新 PentAGI |

## 漢化相關

| 現象 | 原因 | 解法 |
| --- | --- | --- |
| 介面仍有零星英文 | 動態字串（樣板字串）或第三方元件文案 | 見 `PROJECT-STATUS.md` 的待辦清單 |
| 介面出現中英混雜的句子 | 句子被 `{變數}` 切成多段，只有前段被翻 | 跑 `node tools/i18n/check-mixed.mjs` 找出位置後手動修 |
| 改了程式碼但畫面沒變 | 只改了原始碼，沒重建前端 | `cd frontend && pnpm run build`，或重建 Docker 映像 |
| 想改回英文 | 漢化是直接改字串 | `git revert` 漢化 commit，或切回上游 tag |
