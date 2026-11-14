#!/usr/bin/env node
/**
 * 品質檢查：找出「中英混雜」的 JSX 文字節點。
 * 成因：句子被 {expression} 切成多段，只有前段翻到，後段仍是英文。
 *
 * 用法: node tools/i18n/check-mixed.mjs
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
            if (['node_modules', 'graphql', 'test-utils'].includes(entry.name)) {
                continue;
            }
            walk(full, out);
        } else if (entry.name.endsWith('.tsx') && !entry.name.includes('.test.')) {
            out.push(full);
        }
    }
    return out;
}

const CJK = /[\u3400-\u9fff]/;
const EN_WORD = /[A-Za-z]{3,}/;

let hits = 0;
for (const file of walk(path.join(ROOT, 'frontend/src'))) {
    const source = fs.readFileSync(file, 'utf8');
    const sf = ts.createSourceFile(file, source, ts.ScriptTarget.Latest, true, ts.ScriptKind.TSX);
    const visit = (node) => {
        if (ts.isJsxText(node)) {
            const text = node.getText(sf);
            if (CJK.test(text) && EN_WORD.test(text)) {
                const line = sf.getLineAndCharacterOfPosition(node.getStart(sf)).line + 1;
                hits += 1;
                console.log(`${path.relative(ROOT, file).replaceAll('\\', '/')}:${line}  ${text.trim().replaceAll(/\s+/g, ' ').slice(0, 90)}`);
            }
        }
        ts.forEachChild(node, visit);
    };
    visit(sf);
}
console.log(`\n中英混雜的 JSX 文字: ${hits} 處`);
