import { dirname } from 'path';
import { fileURLToPath } from 'url';
import { FlatCompat } from '@eslint/eslintrc';
import prettierPlugin from 'eslint-plugin-prettier';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const compat = new FlatCompat({
  baseDirectory: __dirname,
});

const eslintConfig = [
  // Ignore generated files and build outputs
  {
    ignores: ['node_modules/**', '.next/**', 'out/**', 'build/**', 'next-env.d.ts'],
  },
  // Next.js recommended + TypeScript rules
  ...compat.extends('next/core-web-vitals', 'next/typescript'),
  // Disable rules that conflict with Prettier
  ...compat.extends('prettier'),
  // Register plugins and project-level rules
  {
    plugins: {
      prettier: prettierPlugin,
    },
    rules: {
      // Surface Prettier issues as ESLint errors
      'prettier/prettier': ['error', { endOfLine: 'auto' }],
    },
  },
  // Fallback: explicitly disable triple-slash rule for the generated file
  {
    files: ['next-env.d.ts'],
    rules: {
      '@typescript-eslint/triple-slash-reference': 'off',
    },
  },
];

export default eslintConfig;
