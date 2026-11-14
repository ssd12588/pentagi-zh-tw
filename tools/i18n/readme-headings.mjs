#!/usr/bin/env node
/**
 * README 標題漢化器
 *
 * 依 tools/i18n/readme-headings.json（英文標題 → 中文標題）改寫 `## / ### ` 標題，
 * 並在標題前插入 <a id="原英文錨點"></a>，讓既有連結與目錄持續有效。
 *
 * 用法: node tools/i18n/readme-headings.mjs [--dry-run]
 */
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '../..');
const README = path.join(ROOT, 'README.md');
const MAP = path.join(__dirname, 'readme-headings.json');

/** GitHub 風格的標題錨點 */
export function slugify(text) {
    return text
        .trim()
        .toLowerCase()
        .replaceAll(/[`*_[\]()]/g, '')
        .replaceAll(/[^\p{L}\p{N}\s-]/gu, '')
        .replaceAll(/\s+/g, '-');
}

function main() {
    const dryRun = process.argv.includes('--dry-run');
    const map = JSON.parse(fs.readFileSync(MAP, 'utf8'));
    const lines = fs.readFileSync(README, 'utf8').split(/\r?\n/);

    let changed = 0;
    const out = [];
    for (let i = 0; i < lines.length; i += 1) {
        const line = lines[i];
        const m = /^(#{2,3})\s+(.+?)\s*$/.exec(line);
        if (m && Object.hasOwn(map, m[2])) {
            const [, hashes, title] = m;
            const slug = slugify(title);
            // 若上一行已經是這個錨點就不重複插入
            if (!(out.length > 0 && out[out.length - 1].includes(`id="${slug}"`))) {
                out.push(`<a id="${slug}"></a>`);
            }
            out.push(`${hashes} ${map[title]}`);
            changed += 1;
            continue;
        }
        out.push(line);
    }

    if (dryRun) {
        console.log(`[dry-run] 將改寫標題: ${changed}`);
        return;
    }
    fs.writeFileSync(README, out.join('\n'), 'utf8');
    console.log(`已改寫標題: ${changed}`);
}

if (process.argv[1] && path.resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
    main();
}
