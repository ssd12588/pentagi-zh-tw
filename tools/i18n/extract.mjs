#!/usr/bin/env node
/**
 * PentAGI 漢化工具鏈 —— 字串抽取器
 *
 * 以 TypeScript AST 掃描前端原始碼，抽出「顯示給人看」的字串字面量。
 * 只抽取，不改檔；輸出 tools/i18n/strings.json 供翻譯。
 *
 * 用法:
 *   node tools/i18n/extract.mjs              # 抽取並寫入 strings.json
 *   node tools/i18n/extract.mjs --stdout     # 只印出統計
 */
import fs from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '../..');

// 從 frontend 的依賴取用 TypeScript（本工具不額外引入相依套件）
const requireFromFrontend = createRequire(path.join(ROOT, 'frontend/package.json'));
const ts = requireFromFrontend('typescript');

export { ts };
const SRC = path.join(ROOT, 'frontend/src');
const OUT = path.join(__dirname, 'strings.json');

/** JSX 屬性：值是「顯示文字」的屬性名 */
const DISPLAY_ATTRS = new Set([
    'alt',
    'aria-description',
    'aria-label',
    'aria-placeholder',
    'cancelText',
    'caption',
    'confirmText',
    'description',
    'emptyDescription',
    'emptyMessage',
    'emptyTitle',
    'heading',
    'hint',
    'label',
    'message',
    'okText',
    'placeholder',
    'subtitle',
    'text',
    'title',
    'tooltip',
]);

/** 物件屬性：值是「顯示文字」的屬性名 */
const DISPLAY_PROPS = new Set([
    'cancelText',
    'caption',
    'confirmLabel',
    'confirmText',
    'description',
    'emptyDescription',
    'emptyMessage',
    'emptyTitle',
    'errorMessage',
    'headline',
    'heading',
    'helpText',
    'hint',
    'label',
    'loadingText',
    'message',
    'noResultsText',
    'okText',
    'placeholder',
    'subtext',
    'subtitle',
    'successMessage',
    'text',
    'title',
    'tooltip',
]);

/** toast 家族的呼叫 */
const TOAST_RE = /^toast(\.(success|error|info|warning|loading|message|promise|dismiss))?$/;

/** 測試檔案中比對畫面文字的呼叫 */
const MATCHER_RE =
    /^(screen\.)?(get|find|query|getAll|findAll|queryAll)By(Text|Role|PlaceholderText|LabelText|AltText|Title|DisplayValue)$/;

/** 顯示用字串的最小長度；太短（單一符號）不處理 */
const MIN_LEN = 2;

/** 明顯不該翻的字串（純技術值） */
const SKIP_PATTERNS = [
    /^[a-z0-9_-]+$/, // 小寫識別字 / css class / 狀態值
    /^\d+(\.\d+)?$/, // 純數字
    /^https?:\/\//,
    /^\/[^\s]*$/, // 路徑
    /^\{\{.*\}\}$/,
    /^[A-Z0-9_]+$/, // 常數
];

export function isSkippable(text) {
    const t = text.trim();
    if (t.length < MIN_LEN) {
        return true;
    }
    return SKIP_PATTERNS.some((re) => re.test(t));
}

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

/**
 * 抽出一份檔案的所有候選字串。
 * @returns {{file:string,start:number,end:number,text:string,kind:string,quote:string,ctx:string}[]}
 */
export function extractFromFile(file) {
    const source = fs.readFileSync(file, 'utf8');
    const isTest = /\.test\.tsx?$/.test(file);
    const emp = file.endsWith('.tsx') ? ts.ScriptKind.TSX : ts.ScriptKind.TS;
    const sf = ts.createSourceFile(file, source, ts.ScriptTarget.Latest, true, emp);
    const found = [];
    const rel = path.relative(ROOT, file).replaceAll('\\', '/');

    const push = (node, kind) => {
        const raw = source.slice(node.getStart(sf), node.getEnd());
        const quote = raw[0] === "'" || raw[0] === '"' || raw[0] === '`' ? raw[0] : '';
        const text = quote ? raw.slice(1, -1) : raw;
        if (isSkippable(text)) {
            return;
        }
        const line = sf.getLineAndCharacterOfPosition(node.getStart(sf)).line + 1;
        found.push({
            ctx: `${rel}:${line}`,
            end: node.getEnd(),
            file: rel,
            kind,
            quote,
            start: node.getStart(sf),
            text,
        });
    };

    const visit = (node) => {
        // 1) JSX 純文字
        if (ts.isJsxText(node)) {
            const raw = node.getText(sf);
            const lead = raw.length - raw.trimStart().length;
            const core = raw.trim();
            if (!isSkippable(core)) {
                const line = sf.getLineAndCharacterOfPosition(node.getStart(sf)).line + 1;
                found.push({
                    ctx: `${rel}:${line}`,
                    end: node.getStart(sf) + lead + core.length,
                    file: rel,
                    kind: 'jsx-text',
                    quote: '',
                    start: node.getStart(sf) + lead,
                    text: core,
                });
            }
        }

        // 2) JSX 屬性="..."
        if (ts.isJsxAttribute(node) && node.initializer && ts.isStringLiteral(node.initializer)) {
            const name = node.name.getText(sf);
            if (DISPLAY_ATTRS.has(name)) {
                push(node.initializer, 'jsx-attr');
            }
        }

        // 3.5) JSX 運算式內的文案（三元／?? 等）：<span>{isNew ? 'Create' : 'Save'}</span>
        if (ts.isJsxExpression(node) && node.expression) {
            const collect = (sub, depth = 0) => {
                if (depth > 12) {
                    return;
                }
                if (ts.isStringLiteral(sub)) {
                    push(sub, 'jsx-expr-literal');
                    return;
                }
                if (ts.isObjectLiteralExpression(sub) || ts.isTemplateExpression(sub) || ts.isCallExpression(sub)) {
                    return;
                }
                ts.forEachChild(sub, (child) => collect(child, depth + 1));
            };
            collect(node.expression);
        }

        if (
            ts.isPropertyAssignment(node) &&
            node.initializer &&
            (ts.isStringLiteral(node.initializer) || ts.isNoSubstitutionTemplateLiteral(node.initializer))
        ) {
            const name = node.name.getText(sf).replaceAll(/['"]/g, '');
            if (DISPLAY_PROPS.has(name)) {
                push(node.initializer, 'obj-prop');
            }
        }

        // 4) toast('...') 與測試的文字比對
        if (ts.isCallExpression(node) && node.arguments.length > 0) {
            const callee = node.expression.getText(sf);
            const first = node.arguments[0];
            const isLiteral = ts.isStringLiteral(first) || ts.isNoSubstitutionTemplateLiteral(first);
            if (isLiteral && (TOAST_RE.test(callee) || (isTest && MATCHER_RE.test(callee)))) {
                push(first, isTest ? 'test-matcher' : 'toast');
            }
        }

        ts.forEachChild(node, visit);
    };

    visit(sf);

    // 同一位置可能因 AST 走訪重複命中，這裡去重避免重複替換
    const seen = new Set();
    return found.filter((item) => {
        const key = `${item.start}:${item.end}`;
        if (seen.has(key)) {
            return false;
        }
        seen.add(key);
        return true;
    });
}

function main() {
    const files = walk(SRC);
    const all = [];
    for (const file of files) {
        all.push(...extractFromFile(file));
    }

    const byText = new Map();
    for (const item of all) {
        if (!byText.has(item.text)) {
            byText.set(item.text, { count: 0, kinds: new Set(), refs: [] });
        }
        const bucket = byText.get(item.text);
        bucket.count += 1;
        bucket.kinds.add(item.kind);
        if (bucket.refs.length < 5) {
            bucket.refs.push(item.ctx);
        }
    }

    const sorted = [...byText.entries()]
        .sort((a, b) => b[1].count - a[1].count || a[0].localeCompare(b[0]))
        .map(([text, meta]) => ({
            count: meta.count,
            kinds: [...meta.kinds].sort(),
            refs: meta.refs,
            text,
        }));

    if (!process.argv.includes('--stdout')) {
        fs.writeFileSync(OUT, JSON.stringify(sorted, null, 2) + '\n', 'utf8');
    }

    const production = all.filter((i) => i.kind !== 'test-matcher').length;
    console.log(`掃描檔案: ${files.length}`);
    console.log(`候選字串出現次數: ${all.length}（UI ${production} / 測試 ${all.length - production}）`);
    console.log(`不重複字串: ${sorted.length}`);
    if (!process.argv.includes('--stdout')) {
        console.log(`已寫入: ${path.relative(ROOT, OUT)}`);
    }
}

if (process.argv[1] && path.resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
    main();
}
