# PentAGI 繁體中文漢化版 · 中文文件

本目錄是 PentAGI 繁體中文（zh-TW）漢化版的中文入門與維護文件。

| 文件 | 內容 |
| --- | --- |
| [QUICKSTART.md](QUICKSTART.md) | 從零到第一次登入：安裝、設定 LLM、啟動、驗證 |
| [TROUBLESHOOTING.md](TROUBLESHOOTING.md) | 常見錯誤（中文訊息 → 原因 → 解法） |
| [glossary.md](glossary.md) | 術語對照表（英文 → 繁體中文） |
| [PROJECT-STATUS.md](PROJECT-STATUS.md) | 漢化進度、驗證狀態、待辦與交接說明 |

## Windows 一鍵啟動

```powershell
pwsh -File .\start-pentagi.ps1 -Build    # 建置本地中文映像並啟動（第一次必用）
pwsh -File .\start-pentagi.ps1           # 之後啟動（沿用本地映像）
pwsh -File .\start-pentagi.ps1 -Logs     # 追蹤日誌
pwsh -File .\start-pentagi.ps1 -Down     # 停止服務
```

> **為什麼一定要 `-Build`？** `docker-compose.yml` 預設使用上游英文映像 `vxcontrol/pentagi:latest`；
> 只有用本倉庫的原始碼建置，才會包含繁體中文介面。等價指令：
> `docker compose -f docker-compose.yml -f docker-compose.build.yml up -d --build`

## 這個版本做了什麼

- **Web 介面**：全站文案改為繁體中文（約 1,100 處、140 個檔案），並把 `<html lang>` 改為 `zh-TW`、移除「禁止瀏覽器翻譯」的設定。
- **AI 回覆語言**：**不需要改提示詞**。PentAGI 內建語言偵測（`backend/pkg/templates/prompts/language_chooser.tmpl`），你用中文下指令，代理就會用中文回報；技術欄位（工具參數、指令、payload）仍保持英文以維持精準度。
- **未漢化的部分**：`backend/pkg/templates/prompts/**` 的提示詞、`LICENSE`／`EULA.md` 法律文件、程式碼識別字、GraphQL 欄位、狀態值。

## 想自己重跑一次漢化

漢化流程已腳本化，放在 `tools/i18n/`：

```bash
# 1) 抽出所有「顯示用」字串
node tools/i18n/extract.mjs

# 2) 看還有哪些沒翻（輸出 TSV 清單）
node tools/i18n/apply.mjs --list --out tools/i18n/pending.tsv

# 3) 把翻譯補進 tools/i18n/dict/*.json（英文 → 繁體中文）

# 4) 套用到前端原始碼（含測試斷言）
node tools/i18n/apply.mjs
node tools/i18n/apply.mjs --tests

# 5) 驗證：語法 / 型別 / 測試 / 建置
node tools/i18n/check-syntax.mjs
node tools/i18n/check-mixed.mjs      # 找出中英混雜的句子
cd frontend && pnpm exec tsc -b && pnpm run test && pnpm run build
```
