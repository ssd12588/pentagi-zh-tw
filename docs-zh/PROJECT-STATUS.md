# 漢化專案狀態（交接文件）

> 最後更新：2026-11-15

## 一句話總結

PentAGI 的 **Web 介面已完成繁體中文化並通過建置**；安裝程式（TUI）與全部文件尚未漢化；測試套件有 13 項「斷言仍為英文」的既有落差待收尾。

## 來源與授權

| 項目 | 內容 |
| --- | --- |
| 上游 | <https://github.com/vxcontrol/pentagi> |
| 授權 | MIT（`LICENSE` 原樣保留）；另有 `EULA.md`，散布時請一併保留 |
| 漢化分支 | 本地 `main`（尚未推送；推送前請先確認遠端倉庫與帳號） |

## 已完成

### 0. GitHub 發布

- Repo：<https://github.com/ssd12588/pentagi-zh-tw>（fork 自 vxcontrol/pentagi）
- 已設定中文描述與 12 個 topics

### 0.1 README（GitHub 首頁）

- 已漢化：目錄、總覽、功能特色、目前的能力邊界、快速開始（系統需求／安裝程式／前置權限／Web 設定範圍／仍需伺服器端設定）
- 做法：標題改中文並在標題前插入 `<a id="原英文錨點"></a>`，確保既有連結與目錄不失效
- 未漢化：架構章節、供應商設定細節、進階設定、開發／建置等（約 2,500 行說明文）

### 0.2 安裝程式（TUI）

- 已完成 **313 / 997** 條字串（歡迎畫面、系統檢查、疑難排解、導覽提示、狀態訊息、主要欄位標籤與說明）
- 工具：`tools/i18n/go-extract.mjs`、`go-apply.mjs`、`dict-go/*.json`
- 驗證：`go build ./cmd/installer/...` 通過

### 1. Web 介面（frontend）

- 161 個檔案變更（其中 160 個在 `frontend/src`、43 個是測試檔）
- 翻譯字典 **772 條**（`tools/i18n/dict/*.json`），實際套用約 **1,300 處**字面量
- 覆蓋範圍：導覽、儀表板、流程／任務／子任務、知識庫、資源、範本、設定（供應商／提示詞／API Token／帳號）、檔案管理器、Markdown 編輯器、登入與錯誤頁
- `frontend/index.html`：`<html lang="zh-TW">`，移除 `translate="no"` 與 `notranslate`（原本是為了防止瀏覽器翻譯造成 React DOM 錯亂；介面已是中文，改為允許翻譯）

### 2. 漢化工具鏈（`tools/i18n/`）

| 檔案 | 作用 |
| --- | --- |
| `extract.mjs` | 以 TypeScript AST 抽出「顯示用」字串（JSX 文字、顯示屬性、物件顯示欄位、toast、JSX 運算式內字面量） |
| `apply.mjs` | 套用字典（含去重與重疊保護）；`--list` 列出未翻譯；`--tests` 只處理測試的查詢／斷言參數 |
| `dict/*.json` | 英 → 繁中對照表（依主題分檔） |
| `check-syntax.mjs` | 檢查是否有檔案被改壞（AST 解析失敗），目前 **0** |
| `check-mixed.mjs` | 找出中英混雜的句子（`{變數}` 把句子切斷造成） |
| `strings.json`、`pending*.tsv`、`baseline.txt`、`final*.txt` | 中間產物與測試基線（可刪） |

### 3. 驗證結果

| 檢查 | 結果 |
| --- | --- |
| `node tools/i18n/check-syntax.mjs` | ✅ 0 個無法解析檔案 |
| `pnpm exec tsc -b` | ✅ exit 0 |
| `pnpm run build` | ✅ 成功（built in 18.34s） |
| `pnpm run test` | ⚠️ 313 failed / 2226 passed（**基線為 300 failed**，見下方說明） |

**測試基線說明**：上游測試在 Windows 環境本來就有 300 項失敗（`localStorage.clear()` 不可用、`e2e/tools/*` 依賴 bash 等環境因素），與漢化無關。漢化後為 313 項，多出的 13 項是「測試仍以英文斷言已翻譯的介面文字」。

## 待辦（建議順序）

1. **收尾 13 項測試落差**（每檔 1 項，多為 `document.title`、`formatPromptId`、`connectionStatus` 等固定字串斷言）：
   `src/components/shared/document-title.test.tsx`、`version-panel.test.tsx`、`connection-status.test.tsx`、
   `src/lib/route-titles/format-prompt-id.test.ts`、`src/features/flows/tasks/flow-task.test.tsx`、
   `src/components/shared/file-manager/file-manager.test.tsx`、`file-manager-utils.test.ts`、
   `src/components/shared/markdown-editor/markdown-editor-marked.test.ts`、
   `src/features/knowledges/knowledge-form-blocker.test.tsx`、`src/pages/settings/settings-provider.test.tsx`（6 項）
2. **安裝程式 TUI 漢化**：所有字串集中在 `backend/cmd/installer/wizard/locale/locale.go`（約 150 KB）與 `backend/cmd/installer/processor/locale.go`，屬純常數、機械式替換，需注意中文字寬（lipgloss 排版）。
3. **文件漢化**：`README.md`（4,247 行）建議只加中文導讀與連結；`backend/docs/**`、`frontend/docs/**` 可分批處理。
4. **推送**：確認 GitHub 帳號與遠端後，`git push` 到自己的 fork，並在 README 加上「關於本漢化版」區塊（見下）。

## 刻意不動的地方

| 項目 | 原因 |
| --- | --- |
| `backend/pkg/templates/prompts/**` | 給 LLM 的提示詞，翻譯會改變代理判斷品質。**語言已由 `language_chooser.tmpl` 自動偵測**，用中文提問即得中文回報 |
| `LICENSE`、`EULA.md`、`NOTICE` | 法律文件 |
| GraphQL 欄位、狀態值、程式識別字 | 邏輯依賴 |
| `backend/migrations/**` | SQL 結構，翻譯會壞 |

## 給下一個對話的開場提示

```text
這是 PentAGI（https://github.com/vxcontrol/pentagi）的繁體中文漢化版，位於 pentagi-zh/。
Web 介面已漢化完成並通過 build。請先讀 docs-zh/PROJECT-STATUS.md，
接著處理「待辦」第 1 項（13 項測試落差），再往安裝程式 TUI 漢化推進。
漢化工具在 tools/i18n/，字典在 tools/i18n/dict/*.json。
```
