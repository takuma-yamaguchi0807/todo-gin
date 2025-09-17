'use client';

import type { CSSProperties } from 'react';
import { useState } from 'react';
import Link from 'next/link';
import { Modal } from '@/components/ui/Modal';
import { signupApi } from '@/lib/apiClient';

export default function SignupPage() {
  const [errorOpen, setErrorOpen] = useState(false);
  const [errorTitle, setErrorTitle] = useState<string>('エラー');
  const [errorMsg, setErrorMsg] = useState<string>('');
  const pageStyle: CSSProperties = {
    minHeight: '100svh',
    display: 'grid',
    placeItems: 'center',
    padding: 24,
  };

  const cardStyle: CSSProperties = {
    width: '100%',
    maxWidth: 360,
    display: 'grid',
    gap: 16,
  };

  const labelStyle: CSSProperties = {
    display: 'grid',
    gap: 6,
    textAlign: 'left',
  };

  const inputStyle: CSSProperties = {
    padding: '10px 12px',
    border: '1px solid #ccc',
    borderRadius: 8,
    fontSize: 16,
  };

  const submitStyle: CSSProperties = {
    padding: '12px 16px',
    borderRadius: 8,
    backgroundColor: '#10b981',
    color: '#fff',
    fontWeight: 600,
    border: 'none',
    cursor: 'pointer',
    display: 'block',
    margin: '16px auto 0',
  };

  const helperStyle: CSSProperties = { textAlign: 'center' };

  return (
    <main style={pageStyle}>
      <section style={cardStyle}>
        <h1 style={{ textAlign: 'center' }}>サインアップ</h1>
        <p style={{ textAlign: 'center' }}>メールアドレスとパスワードを入力してください。</p>

        {/* 組み込みバリデーション + onSubmit */}
        <form
          onSubmit={async (e) => {
            e.preventDefault();
            const form = e.currentTarget as HTMLFormElement;
            const formData = new FormData(form);
            const email = String(formData.get('email') || '');
            const password = String(formData.get('password') || '');
            try {
              await signupApi({ email, password });
              // 201 Created（空ボディ想定）: ログイン画面へ遷移
              window.location.href = '/login';
            } catch (err) {
              const eobj = err as Error & { status?: number; kind?: string; field?: string };
              const kind = eobj.kind ?? 'error';
              setErrorTitle('サインアップに失敗しました');
              if (kind === 'invalid') {
                setErrorMsg(eobj.message || '入力内容を確認してください。');
              } else if (kind === 'conflict') {
                setErrorMsg('このメールアドレスは既に登録されています。');
              } else {
                setErrorMsg(eobj.message || 'サーバーでエラーが発生しました。');
              }
              setErrorOpen(true);
            }
          }}
        >
          <label style={labelStyle}>
            <span>メールアドレス</span>
            <input
              name="email"
              type="email"
              placeholder="you@example.com"
              required
              style={inputStyle}
              autoComplete="email"
              inputMode="email"
            />
          </label>

          <label style={labelStyle}>
            <span>パスワード</span>
            <input
              name="password"
              type="password"
              placeholder="8文字以上を推奨"
              required
              minLength={8}
              style={inputStyle}
              autoComplete="new-password"
            />
          </label>

          <button type="submit" style={submitStyle} aria-label="サインアップ">
            サインアップ
          </button>
        </form>

        <p style={helperStyle}>
          すでにアカウントをお持ちの方は <Link href="/login">ログイン</Link>
        </p>

        <p style={helperStyle}>
          <Link href="/">← ホームへ戻る</Link>
        </p>
      </section>
      <Modal
        open={errorOpen}
        title={errorTitle}
        message={errorMsg}
        onClose={() => setErrorOpen(false)}
      />
    </main>
  );
}
