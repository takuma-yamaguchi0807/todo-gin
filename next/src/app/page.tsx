import Link from 'next/link';

export default function Home() {
  return (
    <main style={{ padding: 24 }}>
      <h1>TODO App</h1>
      <p>学習用 TODO アプリへようこそ。</p>
      <nav aria-label="メインナビゲーション">
        <ul style={{ display: 'grid', gap: 8, padding: 0, listStyle: 'none', marginTop: 16 }}>
          <li>
            <Link href="/login">ログイン</Link>
          </li>
          <li>
            <Link href="/signup">サインアップ</Link>
          </li>
        </ul>
      </nav>
    </main>
  );
}
