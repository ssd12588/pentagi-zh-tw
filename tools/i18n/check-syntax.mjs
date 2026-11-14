#!/usr/bin/env node
/**
 * 語法檢查：找出被翻譯工具改壞（無法解析）的檔案。
 * 用法: node tools/i18n/check-syntax.mjs
 */
import fs from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '../..');
const requireFromFrontend = createRequire(path.join(ROOT, 'frontend/package.json'));
const ts = requireFromFrontend('typescript');

function walk(dir, out = []) {
    for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
        const full = path.join(dir, entry.name);
        if (entry.isDirectory()) {
            if (['node_modules', 'dist'].includes(entry.name)) {
                continue;
            }
            walk(full, out);
        } else if (/\.tsx?$/.test(entry.name)) {
            out.push(full);
        }
    }
    return out;
}

let bad = 0;
for (const file of walk(path.join(ROOT, 'frontend/src'))) {
    const source = fs.readFileSync(file, 'utf8');
    const kind = file.endsWith('.tsx') ? ts.ScriptKind.TSX : ts.ScriptKind.TS;
    const sf = ts.createSourceFile(file, source, ts.ScriptTarget.Latest, true, kind);
    const diags = sf.parseDiagnostics ?? [];
    if (diags.length > 0) {
        bad += 1;
        const pos = sf.getLineAndCharacterOfPosition(diags[0].start);
        const line = source.split(/\r?\n/)[pos.line] ?? '';
        console.log(`${path.relative(ROOT, file).replaceAll('\\', '/')}:${pos.line + 1}  ${line.trim().slice(0, 100)}`);
    }
}
console.log(`\n無法解析的檔案: ${bad}`);
process.exitCode = bad > 0 ? 1 : 0;
