import { readFile } from 'node:fs/promises';
import { spawnSync } from 'node:child_process';
import { fileURLToPath } from 'node:url';
import { dirname, resolve } from 'node:path';

const root = resolve(dirname(fileURLToPath(import.meta.url)), '..');
const generated = ['src/grammar.json', 'src/node-types.json', 'src/parser.c'];
const before = await Promise.all(generated.map(path => readFile(resolve(root, path))));

function run(command, args) {
  const printable = [command, ...args].join(' ');
  console.log(`\n$ ${printable}`);
  const executable = process.platform === 'win32' ? `${command}.exe` : command;
  const result = spawnSync(executable, args, {
    cwd: root,
    stdio: 'inherit',
  });
  if (result.error) throw result.error;
  if (result.status !== 0) throw new Error(`${printable} failed with exit ${result.status}`);
}

try {
  run('tree-sitter', ['generate', '--abi', '15']);
  let changed = false;
  for (const [index, path] of generated.entries()) {
    const after = await readFile(resolve(root, path));
    if (!before[index].equals(after)) {
      console.error(`Generated artifact differs from committed source: ${path}`);
      changed = true;
    }
  }
  if (changed) throw new Error('Regenerate and review generated parser artifacts');
  run('tree-sitter', ['test']);
  run('go', ['test', '-count=1', './...']);
  run('cargo', ['test']);
  console.log('\nVerification passed: reproducible parser, corpus, reviewed sources, fixtures, queries, Go editor adapter, Rust consumer binding.');
} catch (error) {
  console.error(error.message);
  process.exitCode = 1;
}
