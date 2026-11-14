#!/usr/bin/env node
/**
 * PentAGI 漢化工具鏈 —— 字典套用器
 *
 * 依 tools/i18n/zh-TW.json（英文 → 繁體中文）把前端「顯示用字面量」就地替換。
 * 只替換 extract.mjs 認定的顯示位置（JSX 文字、顯示屬性、toast、測試比對），
 * 不動識別字、GraphQL、狀態值比較、className、路徑。
 *
 * 用法:
 *   node tools/i18n/apply.mjs              # 套用並寫回檔案
 *   node tools/i18n/apply.mjs --dry-run    # 只統計，不寫檔
 *   node tools/i18n/apply.mjs --list [N]   # 列出尚未翻譯的字串（前 N 筆，預設全部）
 */
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

import { extractFromFile, isSkippable, ts } from './extract.mjs';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '../..');
const DICT_FILE = path.join(__dirname, 'zh-TW.json');

function walk(dir, out = []) {
    for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
        const full = path.join(dir, entry.name);
        if (entry.isDirectory()) {
            if (entry.name === 'node_modules' || entry.name === 'graphql' || entry.name === 'test-utils') {
                continue;
            }
            walk(full, out);
        } else if (/\.(tsx?|jsx?)$/.test(entry.name)) {
            out.push(full);
        }
    }
    return out;
}

export function loadDict() {
    const dir = path.join(__dirname, 'dict');
    const merged = {};
    if (!fs.existsSync(dir)) {
        return merged;
    }
    for (const name of fs.readdirSync(dir).filter((n) => n.endsWith('.json')).sort()) {
        Object.assign(merged, JSON.parse(fs.readFileSync(path.join(dir, name), 'utf8')));
    }
    return merged;
}

/** 依引號種類安全跳脫 */
function encode(text, quote) {
    if (!quote) {
        // JSX 文字：跳脫大括號與角括號，避免變成 expression
        return text.replaceAll('{', '&#123;').replaceAll('}', '&#125;');
    }
    if (quote === '`') {
        return text.replaceAll('\\', '\\\\').replaceAll('`', '\\`').replaceAll('${', '\\${');
    }
    return text.replaceAll('\\', '\\\\').replaceAll(quote, `\\${quote}`);
}

function escapeRe(text) {
    return text.replaceAll(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

/** 測試中「查詢／斷言畫面文字」的呼叫名 */
const MATCHER_NAME_RE =
    /^(get|find|query|getAll|findAll|queryAll)By(Text|Role|PlaceholderText|LabelText|AltText|Title|DisplayValue)$/;
const ASSERT_NAME_RE =
    /^(toHaveTextContent|toHaveAccessibleName|toHaveAccessibleDescription|toHaveValue|toBe|toEqual|toStrictEqual|toMatch|toContain|toHaveBeenCalledWith|toHaveBeenLastCalledWith|toHaveBeenNthCalledWith|toHaveBeenCalledTimes)$/;

function calleeName(node) {
    const expr = node.expression;
    if (ts.isIdentifier(expr)) {
        return expr.text;
    }
    if (ts.isPropertyAccessExpression(expr)) {
        return expr.name.text;
    }
    return '';
}

/**
 * 測試檔模式：只改「查詢／斷言呼叫的參數」，不碰 mock 資料與識別值。
 * - 字串參數：整串等於字典鍵，或以詞邊界包含字典鍵
 * - 正則參數：在 pattern 內以詞邊界替換（長鍵優先）
 * - getByRole 第二參數的 name 選項
 */
export function applyTests(dict, { dryRun = false } = {}) {
    const keys = Object.keys(dict).sort((a, b) => b.length - a.length);
    const files = walkTests(path.join(ROOT, 'frontend'));
    const stats = { changedFiles: 0, matches: 0 };

    const replaceInText = (text) => {
        return Object.hasOwn(dict, text) ? dict[text] : text;
    };

    // 正則參數：整段以字典替換後，若幾乎沒有殘留英文（≤3 個字母，允許 v2.4.0 這類）才採用
    const replaceInPattern = (pattern, flags = '') => {
        const leading = pattern.startsWith('^') ? '^' : '';
        const trailing = pattern.endsWith('$') ? '$' : '';
        const body = pattern.slice(leading.length, pattern.length - trailing.length);
        const unescaped = body.replaceAll(/\\(.)/g, '$1');

        if (Object.hasOwn(dict, unescaped)) {
            return leading + dict[unescaped] + trailing;
        }

        let candidate = body;
        for (const key of keys) {
            const re = new RegExp(`(?<![A-Za-z0-9_])${escapeRe(key)}(?![A-Za-z0-9_])`, flags.includes('i') ? 'gi' : 'g');
            candidate = candidate.replaceAll(re, dict[key]);
        }
        const leftover = (candidate.match(/[A-Za-z]/g) ?? []).length;
        if (candidate !== body && leftover <= 3) {
            return leading + candidate + trailing;
        }
        return pattern;
    };

    for (const file of files) {
        const source = fs.readFileSync(file, 'utf8');
        const sf = ts.createSourceFile(file, source, ts.ScriptTarget.Latest, true, ts.ScriptKind.TSX);
        const edits = [];

        const handleArg = (arg) => {
            if (ts.isCallExpression(arg)) {
                // expect.stringContaining('...') 這類包裝，往內一層處理
                arg.arguments.forEach(handleArg);
                return;
            }
            if (ts.isStringLiteral(arg) || ts.isNoSubstitutionTemplateLiteral(arg)) {
                const raw = source.slice(arg.getStart(sf), arg.getEnd());
                const quote = raw[0];
                const text = raw.slice(1, -1);
                const next = replaceInText(text);
                if (next !== text) {
                    edits.push({ end: arg.getEnd(), start: arg.getStart(sf), text: quote + encode(next, quote) + quote });
                }
                return;
            }
            if (ts.isRegularExpressionLiteral(arg)) {
                const raw = source.slice(arg.getStart(sf), arg.getEnd());
                const lastSlash = raw.lastIndexOf('/');
                const pattern = raw.slice(1, lastSlash);
                const flags = raw.slice(lastSlash + 1);
                const next = replaceInPattern(pattern, flags);
                if (next !== pattern) {
                    edits.push({ end: arg.getEnd(), start: arg.getStart(sf), text: `/${next}/${flags}` });
                }
                return;
            }
            if (ts.isArrayLiteralExpression(arg)) {
                for (const el of arg.elements) {
                    handleArg(el);
                }
                return;
            }
            if (ts.isObjectLiteralExpression(arg)) {
                for (const prop of arg.properties) {
                    if (!ts.isPropertyAssignment(prop) || !prop.name) {
                        continue;
                    }
                    const name = prop.name.getText(sf).replaceAll(/['"]/g, '');
                    if (name !== 'name' && name !== 'text' && name !== 'description' && name !== 'message' && name !== 'title') {
                        continue;
                    }
                    const init = prop.initializer;
                    if (ts.isStringLiteral(init)) {
                        const text = init.text;
                        const next = replaceInText(text);
                        if (next !== text) {
                            edits.push({
                                end: init.getEnd(),
                                start: init.getStart(sf),
                                text: `'${encode(next, "'")}'`,
                            });
                        }
                    } else if (ts.isRegularExpressionLiteral(init)) {
                        const raw = init.getText(sf);
                        const lastSlash = raw.lastIndexOf('/');
                        const pattern = raw.slice(1, lastSlash);
                        const flags = raw.slice(lastSlash + 1);
                        const next = replaceInPattern(pattern, flags);
                        if (next !== pattern) {
                            edits.push({ end: init.getEnd(), start: init.getStart(sf), text: `/${next}/${flags}` });
                        }
                    }
                }
            }
        };

        const visit = (node) => {
            if (ts.isCallExpression(node)) {
                const name = calleeName(node);
                if (MATCHER_NAME_RE.test(name) || ASSERT_NAME_RE.test(name)) {
                    node.arguments.forEach(handleArg);
                } else if (name === 'each') {
                    // it.each([['Bold', ...], ['Italic', ...]](...)) — 遞迴翻譯所有字串元素
                    const walkArray = (arr) => {
                        for (const el of arr.elements) {
                            if (ts.isStringLiteral(el)) {
                                const text = el.text;
                                const next = replaceInText(text);
                                if (next !== text) {
                                    edits.push({
                                        end: el.getEnd(),
                                        start: el.getStart(sf),
                                        text: `'${encode(next, "'")}'`,
                                    });
                                }
                            } else if (ts.isArrayLiteralExpression(el)) {
                                walkArray(el);
                            }
                        }
                    };
                    for (const arg of node.arguments) {
                        if (ts.isArrayLiteralExpression(arg)) {
                            walkArray(arg);
                        }
                    }
                }
            }
            ts.forEachChild(node, visit);
        };

        visit(sf);
        if (edits.length === 0) {
            continue;
        }
        stats.matches += edits.length;
        stats.changedFiles += 1;
        if (dryRun) {
            continue;
        }
        let text = source;
        const seen = new Set();
        let floor = Number.POSITIVE_INFINITY;
        for (const edit of edits.sort((a, b) => b.start - a.start)) {
            const key = `${edit.start}:${edit.end}`;
            if (seen.has(key) || edit.end > floor) {
                continue;
            }
            seen.add(key);
            text = text.slice(0, edit.start) + edit.text + text.slice(edit.end);
            floor = edit.start;
        }
        fs.writeFileSync(file, text, 'utf8');
    }
    return stats;
}

function walkTests(dir, out = []) {
    for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
        const full = path.join(dir, entry.name);
        if (entry.isDirectory()) {
            if (entry.name === 'node_modules' || entry.name === 'dist' || entry.name === 'playwright-report') {
                continue;
            }
            walkTests(full, out);
        } else if (/\.test\.tsx?$/.test(entry.name)) {
            out.push(full);
        }
    }
    return out;
}

function main() {
    const argv = process.argv.slice(2);
    const dict = loadDict();
    const dryRun = argv.includes('--dry-run');

    if (argv.includes('--tests')) {
        const stats = applyTests(dict, { dryRun });
        console.log(`${dryRun ? '[dry-run] ' : ''}測試檔字面量替換: ${stats.matches} 處 / ${stats.changedFiles} 檔`);
        return;
    }

    const files = walk(path.join(ROOT, 'frontend/src'));
    const allItems = [];
    for (const file of files) {
        allItems.push(...extractFromFile(file));
    }

    if (argv.includes('--list')) {
        const pending = new Map();
        for (const item of allItems) {
            if (dict[item.text] || isSkippable(item.text)) {
                continue;
            }
            if (!pending.has(item.text)) {
                pending.set(item.text, { count: 0, ui: false, contexts: [] });
            }
            const bucket = pending.get(item.text);
            bucket.count += 1;
            const isTest = /\.test\.tsx?$/.test(item.file);
            if (!isTest) {
                bucket.ui = true;
            }
            if (bucket.contexts.length < 2) {
                bucket.contexts.push(item.ctx);
            }
        }
        const rows = [...pending.entries()]
            .filter(([text, meta]) => meta.ui && !/[\r\n]/.test(text) && !/[\u3400-\u9fff]/.test(text) && /[A-Z]/.test(text))
            .sort((a, b) => b[1].count - a[1].count || a[0].localeCompare(b[0]));

        const lines = rows.map(([text, meta]) => `${meta.count}\t${text}`);
        const outIndex = argv.indexOf('--out');
        if (outIndex !== -1 && argv[outIndex + 1]) {
            fs.writeFileSync(path.resolve(ROOT, argv[outIndex + 1]), lines.join('\n') + '\n', 'utf8');
            console.log(`尚未翻譯（UI 專用）: ${rows.length} → ${argv[outIndex + 1]}`);
            return;
        }
        console.log(`尚未翻譯（UI 專用）: ${rows.length}`);
        for (const line of lines) {
            console.log(line);
        }
        return;
    }

    // 依檔案分組，從後往前替換，避免位移失效
    const byFile = new Map();
    let matched = 0;
    let missing = 0;
    for (const item of allItems) {
        const zh = dict[item.text];
        if (!zh) {
            missing += 1;
            continue;
        }
        if (!byFile.has(item.file)) {
            byFile.set(item.file, []);
        }
        byFile.get(item.file).push({ ...item, zh });
        matched += 1;
    }

    let changedFiles = 0;
    for (const [rel, items] of byFile) {
        const abs = path.join(ROOT, rel);
        const original = fs.readFileSync(abs, 'utf8');
        let text = original;

        // 去重（同一位置只保留一筆）並由後往前套用，重疊者直接跳過
        const seen = new Set();
        const unique = items.filter((item) => {
            const key = `${item.start}:${item.end}`;
            if (seen.has(key)) {
                return false;
            }
            seen.add(key);
            return true;
        });
        let floor = Number.POSITIVE_INFINITY;
        for (const item of unique.sort((a, b) => b.start - a.start)) {
            if (item.end > floor) {
                continue;
            }
            const replacement = item.quote ? item.quote + encode(item.zh, item.quote) + item.quote : encode(item.zh, '');
            text = text.slice(0, item.start) + replacement + text.slice(item.end);
            floor = item.start;
        }
        if (text !== original) {
            changedFiles += 1;
            if (!dryRun) {
                fs.writeFileSync(abs, text, 'utf8');
            }
        }
    }

    console.log(`${dryRun ? '[dry-run] ' : ''}命中字典: ${matched} 處 / ${byFile.size} 檔`);
    console.log(`未命中（保持原文）: ${missing} 處`);
    console.log(`${dryRun ? '預計變更' : '已變更'}檔案: ${changedFiles}`);
}

if (process.argv[1] && path.resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
    main();
}

