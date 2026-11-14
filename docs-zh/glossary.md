# 術語對照表（English → 繁體中文）

漢化時一律照此表翻譯，避免同一詞多種譯法。「保留原文」的項目代表刻意不翻。

## 核心概念

| 英文 | 繁體中文 | 備註 |
| --- | --- | --- |
| Flow | 流程 | 一次完整的任務（一個 engagement） |
| Task | 任務 | 流程下的任務 |
| Subtask | 子任務 | |
| Assistant | 助理 | 與使用者對話的入口 |
| Automation | 自動化 | 流程執行期間的訊息串 |
| Agent | 代理 | 執行任務的 AI 角色 |
| Pentester | 滲透測試代理 | 保留「Pentester」於程式碼中，介面譯為滲透測試代理 |
| Coder | 程式代理 | |
| Memorist | 記憶代理 | |
| Searcher | 搜尋代理 | |
| Adviser | 顧問代理 | |
| Reflector | 反思代理 | |
| Enricher | 補充代理 | |
| Generator | 生成代理 | |
| Refiner | 精煉代理 | |
| Reporter | 報告代理 | |
| Summarizer | 摘要代理 | |
| Knowledge / Knowledges | 知識庫 | |
| Resource | 資源 | 上傳的檔案／資料夾 |
| Template | 範本 | 任務範本 |
| Prompt | 提示詞 | |
| Human Prompt | 使用者提示詞 | 對應 system prompt 之相對 |
| System Prompt | 系統提示詞 | |
| Provider | 供應商 | LLM 供應商 |
| Model | 模型 | |
| Tool Call | 工具呼叫 | |
| Terminal | 終端機 | |
| Container | 容器 | |
| Vector Store | 向量儲存 | |
| Screenshot | 截圖 | |
| Dashboard | 儀表板 | |
| Analytics | 分析 | |
| Token | Token | **保留原文** |
| Reasoning | 推理 | |
| Reasoning Effort | 推理強度 | |
| Reasoning Mode | 推理模式 | |
| Adaptive | 自適應 | |
| Budget | 預算 | |
| Temperature | 溫度 | 取樣參數，保留「溫度」 |

## 常用的動作詞

| 英文 | 繁體中文 |
| --- | --- |
| Create | 建立 |
| Edit | 編輯 |
| Delete | 刪除 |
| Rename | 重新命名 |
| Duplicate / Clone | 建立副本 / 複製一份 |
| Save | 儲存 |
| Submit | 送出 |
| Cancel | 取消 |
| Reset | 重設 |
| Retry / Try again | 重試 |
| Upload | 上傳 |
| Download | 下載 |
| Attach | 附加 |
| Promote | 升階 |
| Copy | 複製 |
| Move | 移動 |
| Refresh | 重新整理 |
| Reload | 重新載入 |
| Filter | 篩選 |
| Search | 搜尋 |
| Select all | 全選 |
| Expand all / Collapse all | 全部展開 / 全部收合 |

## 計費與用量

| 英文 | 繁體中文 |
| --- | --- |
| Input Price | 輸入價格 |
| Output Price | 輸出價格 |
| Cache Read Price | 快取讀取價格 |
| Cache Write Price | 快取寫入價格 |
| Total Cost | 總成本 |
| Tokens In / Out | 輸入 Token / 輸出 Token |
| Cache In / Out | 快取讀取 / 快取寫入 |
| Usage | 用量 |

## 刻意保留原文

| 項目 | 原因 |
| --- | --- |
| PentAGI、OpenAI、Anthropic、Gemini、Bedrock、DeepSeek、Qwen、GLM、Kimi、MiniMax、Mistral、xAI、MoonshotAI、Ollama、Swagger UI、OAuth | 品牌／產品名 |
| XSS、SSRF、IDOR、JWT、CVE、CVSS、PoC、API、URL、JSON、GraphQL、PDF、MD | 產業慣用縮寫 |
| Top K / Top P | 取樣參數原名 |
| `prompts/**`、`skills/**` 內文 | 屬於給模型的指令，翻譯會降低判斷品質 |
| 狀態值（`running`、`finished`…）、GraphQL 欄位 | 程式邏輯依賴 |
| `LICENSE`、`EULA.md` | 法律文件 |

## 風格規範

- 繁體中文，全角標點（，。：；（）「」）。
- 中英之間不加空格：`使用 Docker`；數字與單位加空格：`5 分鐘`。
- 指令、參數、路徑、環境變數一律原樣呈現：`docker compose up -d`、`OPEN_AI_KEY`。
- 破折號用 `——`，引號用「」。
