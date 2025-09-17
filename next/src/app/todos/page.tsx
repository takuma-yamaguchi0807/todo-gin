'use client';

import type { CSSProperties } from 'react';
import { useEffect, useMemo, useState } from 'react';
import Link from 'next/link';
import { Modal } from '@/components/ui/Modal';
import { deleteTodosApi, getTodosApi } from '@/lib/apiClient';
import type { Todo } from '@/types/todo';

export default function TodoListPage() {
  const pageStyle: CSSProperties = { padding: 24 };
  const toolbarStyle: CSSProperties = {
    display: 'flex',
    gap: 12,
    alignItems: 'center',
    marginBottom: 16,
  };
  const gridStyle: CSSProperties = {
    width: '100%',
    display: 'grid',
    gridTemplateColumns: 'auto 1fr auto auto',
    gap: 8,
    alignItems: 'center',
  };
  const headerStyle: CSSProperties = { fontWeight: 700 };
  const buttonStyle: CSSProperties = {
    padding: '8px 12px',
    borderRadius: 8,
    backgroundColor: '#ef4444',
    color: '#fff',
    border: 'none',
    cursor: 'pointer',
  };
  const linkButtonStyle: CSSProperties = {
    padding: '8px 12px',
    borderRadius: 8,
    backgroundColor: '#0ea5e9',
    color: '#fff',
    textDecoration: 'none',
  };

  const [items, setItems] = useState<Todo[]>([]);
  const [checked, setChecked] = useState<Record<string, boolean>>({});
  const [loading, setLoading] = useState(true);
  const [errorOpen, setErrorOpen] = useState(false);
  const [errorMsg, setErrorMsg] = useState('');

  useEffect(() => {
    (async () => {
      try {
        const data = await getTodosApi();
        setItems(data);
      } catch (err) {
        const eobj = err as Error & { kind?: string };
        setErrorMsg(eobj.message || '一覧取得に失敗しました');
        setErrorOpen(true);
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  const selectedIds = useMemo(
    () =>
      Object.entries(checked)
        .filter(([, v]) => v)
        .map(([k]) => k),
    [checked],
  );
  const toggle = (id: string, value: boolean) => setChecked((p) => ({ ...p, [id]: value }));
  const toggleAll = (value: boolean) =>
    setChecked(Object.fromEntries(items.map((t) => [t.id, value])));

  return (
    <main style={pageStyle}>
      <h1>TODO 一覧</h1>
      <div style={toolbarStyle}>
        <label style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
          <input
            type="checkbox"
            onChange={(e) => toggleAll(e.currentTarget.checked)}
            aria-label="全選択"
          />{' '}
          全選択
        </label>
        <button
          type="button"
          style={buttonStyle}
          disabled={selectedIds.length === 0}
          onClick={async () => {
            if (selectedIds.length === 0) return;
            if (!confirm(`${selectedIds.length} 件を削除します。よろしいですか？`)) return;
            try {
              await deleteTodosApi(selectedIds);
              setItems((prev) => prev.filter((t) => !selectedIds.includes(t.id)));
              setChecked({});
            } catch (err) {
              const eobj = err as Error & { kind?: string };
              setErrorMsg(eobj.message || '削除に失敗しました');
              setErrorOpen(true);
            }
          }}
        >
          選択削除
        </button>
        <Link href="/todos/new" style={linkButtonStyle}>
          新規作成
        </Link>
      </div>

      {loading ? (
        <p>読み込み中...</p>
      ) : items.length === 0 ? (
        <p>TODOはありません。</p>
      ) : (
        <div role="table" style={{ display: 'grid', gap: 8 }}>
          <div role="row" style={gridStyle}>
            <div role="columnheader" style={headerStyle}>
              選択
            </div>
            <div role="columnheader" style={headerStyle}>
              タイトル
            </div>
            <div role="columnheader" style={headerStyle}>
              ステータス
            </div>
            <div role="columnheader" style={headerStyle}>
              詳細
            </div>
          </div>
          {items.map((t) => (
            <div role="row" style={gridStyle} key={t.id}>
              <div>
                <input
                  type="checkbox"
                  checked={!!checked[t.id]}
                  onChange={(e) => toggle(t.id, e.currentTarget.checked)}
                  aria-label={`${t.title} を選択`}
                />
              </div>
              <div>{t.title}</div>
              <div>{t.status}</div>
              <div>
                <Link href={`/todos/${t.id}`}>詳細</Link>
              </div>
            </div>
          ))}
        </div>
      )}

      <Modal
        open={errorOpen}
        title="エラー"
        message={errorMsg}
        onClose={() => setErrorOpen(false)}
      />
    </main>
  );
}
