#!/usr/bin/env node
/**
 * 安裝程式（Go）字串抽取器
 *
 * 只處理「純字串常數」檔（locale.go 等），不動任何邏輯比較值。
 * 用法:
 *   node tools/i18n/go-extract.mjs                 # 掃描預設檔案並輸出清單
 *   node tools/i18n/go-extract.mjs --out x.tsv     # 輸出 TSV（次數\t字串）
 */
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '../..');

/** 只處理這些「顯示文字集中檔」 */
export const TARGET_FILES = [
    'backend/cmd/installer/wizard/locale/locale.go',
    'backend/cmd/installer/processor/locale.go',
];

const SKIP = [
    /^[a-z0-9_.\-\/]+$/, // 識別字、路徑、套件名
    /^[A-Z0-9_]+$/, // 全大寫常數／環境變數
    /^https?:\/\//,
    /^%[sdvq]$/,
    /^\s*$/,
];

export function isSkippable(text) {
    const t = text.trim();
    if (t.length < 3) {
        return true;
    }
    return SKIP.some((re) => re.test(t));
}

/** 抽出檔案內所有雙引號字串常量的內容 */
export function extractGoStrings(file) {
    const src = fs.readFileSync(file, 'utf8');
    const out = [];
    const lines = src.split(/\r?\n/);
    lines.forEach((line, i) => {
        // 略過 import 區塊與註解行
        const trimmed = line.trim();
        if (trimmed.startsWith('//') || trimmed.startsWith('import') || trimmed.startsWith('"')) {
            return;
        }
        const re = /"((?:[^"\\]|\\.)*)"/g;
        let m;
        while ((m = re.exec(line)) !== null) {
            const text = m[1];
            if (isSkippable(text)) {
                continue;
            }
            out.push({ file, line: i + 1, text });
        }
    });
    return out;
}

function main() {
    const argv = process.argv.slice(2);
    const all = [];
    for (const rel of TARGET_FILES) {
        const abs = path.join(ROOT, rel);
        if (!fs.existsSync(abs)) {
            continue;
        }
        all.push(...extractGoStrings(abs));
    }

    const counts = new Map();
    for (const item of all) {
        counts.set(item.text, (counts.get(item.text) ?? 0) + 1);
    }
    const rows = [...counts.entries()].sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]));

    const outAt = argv.indexOf('--out');
    if (outAt !== -1 && argv[outAt + 1]) {
        const lines = rows.map(([text, n]) => `${n}\t${text}`);
        fs.writeFileSync(path.resolve(ROOT, argv[outAt + 1]), lines.join('\n') + '\n', 'utf8');
        console.log(`字串出現次數: ${all.length}；不重複: ${rows.length} → ${argv[outAt + 1]}`);
        return;
    }
    console.log(`字串出現次數: ${all.length}；不重複: ${rows.length}`);
}

if (process.argv[1] && path.resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
    main();
}
