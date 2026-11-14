#!/usr/bin/env node
/**
 * 安裝程式（Go）字串套用器
 *
 * 依 tools/i18n/dict-go/*.json（英文 → 繁體中文）替換「顯示文字集中檔」的字串常量。
 * 只替換完全相等的雙引號字串內容，不動識別字與邏輯值。
 *
 * 用法:
 *   node tools/i18n/go-apply.mjs --dry-run
 *   node tools/i18n/go-apply.mjs
 *   node tools/i18n/go-apply.mjs --list --out tools/i18n/go-pending.tsv
 */
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

import { TARGET_FILES, extractGoStrings, isSkippable } from './go-extract.mjs';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '../..');
const DICT_DIR = path.join(__dirname, 'dict-go');

function loadDict() {
    const merged = {};
    if (!fs.existsSync(DICT_DIR)) {
        return merged;
    }
    for (const name of fs.readdirSync(DICT_DIR).filter((n) => n.endsWith('.json')).sort()) {
        Object.assign(merged, JSON.parse(fs.readFileSync(path.join(DICT_DIR, name), 'utf8')));
    }
    return merged;
}

/** Go 雙引號字串的跳脫 */
function encodeGo(text) {
    return text.replaceAll('\\', '\\\\').replaceAll('"', '\\"');
}

function main() {
    const argv = process.argv.slice(2);
    const dict = loadDict();
    const dryRun = argv.includes('--dry-run');

    const files = TARGET_FILES.map((rel) => ({ abs: path.join(ROOT, rel), rel })).filter((f) => fs.existsSync(f.abs));

    if (argv.includes('--list')) {
        const counts = new Map();
        for (const file of files) {
            for (const item of extractGoStrings(file.abs)) {
                if (dict[item.text] || isSkippable(item.text) || /[\u3400-\u9fff]/.test(item.text)) {
                    continue;
                }
                counts.set(item.text, (counts.get(item.text) ?? 0) + 1);
            }
        }
        const rows = [...counts.entries()].sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]));
        const outAt = argv.indexOf('--out');
        const lines = rows.map(([text, n]) => `${n}\t${text}`);
        if (outAt !== -1 && argv[outAt + 1]) {
            fs.writeFileSync(path.resolve(ROOT, argv[outAt + 1]), lines.join('\n') + '\n', 'utf8');
            console.log(`未翻譯: ${rows.length} → ${argv[outAt + 1]}`);
        } else {
            console.log(`未翻譯: ${rows.length}`);
            for (const line of lines) {
                console.log(line);
            }
        }
        return;
    }

    let total = 0;
    let changed = 0;
    for (const file of files) {
        const src = fs.readFileSync(file.abs, 'utf8');
        let out = src;
        let fileHits = 0;
        // 長字串優先，避免短字串先吃掉長句
        const keys = Object.keys(dict).sort((a, b) => b.length - a.length);
        for (const key of keys) {
            const needle = '"' + key + '"';
            if (!out.includes(needle)) {
                continue;
            }
            const before = out;
            out = out.split(needle).join('"' + encodeGo(dict[key]) + '"');
            if (out !== before) {
                fileHits += 1;
            }
        }
        total += fileHits;
        if (out !== src) {
            changed += 1;
            if (!dryRun) {
                fs.writeFileSync(file.abs, out, 'utf8');
            }
        }
    }
    console.log(`${dryRun ? '[dry-run] ' : ''}命中字串種類: ${total}；${dryRun ? '預計變更' : '已變更'}檔案: ${changed}`);
}

if (process.argv[1] && path.resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
    main();
}
