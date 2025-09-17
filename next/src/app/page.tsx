import type { CSSProperties } from 'react';
import Link from 'next/link';

export default function Home() {
  const pageStyle: CSSProperties = {
    minHeight: '100svh',
    display: 'grid',
    placeItems: 'center',
    padding: 24,
  };

  const containerStyle: CSSProperties = {
    textAlign: 'center',
    display: 'grid',
    gap: 16,
  };

  const navStyle: CSSProperties = {
    display: 'grid',
    gap: 12,
    gridTemplateColumns: '1fr 1fr',
    maxWidth: 360,
    width: '100%',
    margin: '0 auto',
  };

  const buttonBaseStyle: CSSProperties = {
    display: 'inline-block',
    padding: '12px 16px',
    borderRadius: 8,
    textAlign: 'center',
    fontWeight: 600,
    textDecoration: 'none',
  };

  const loginButtonStyle: CSSProperties = {
    ...buttonBaseStyle,
    backgroundColor: '#0ea5e9',
    color: '#ffffff',
  };

  const signupButtonStyle: CSSProperties = {
    ...buttonBaseStyle,
    backgroundColor: '#10b981',
    color: '#ffffff',
  };

  return (
    <main style={pageStyle}>
      <section style={containerStyle}>
        <h1>TODO App</h1>
        <p>学習用 TODO アプリへようこそ。</p>
        <nav aria-label="メインナビゲーション" style={navStyle}>
          <Link href="/login" style={loginButtonStyle}>
            ログイン
          </Link>
          <Link href="/signup" style={signupButtonStyle}>
            サインアップ
          </Link>
        </nav>
      </section>
    </main>
  );
}
