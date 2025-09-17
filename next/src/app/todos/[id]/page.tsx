'use client';

import type { CSSProperties } from 'react';
import { useEffect, useState } from 'react';
import Link from 'next/link';
import { useParams } from 'next/navigation';
import { Modal } from '@/components/ui/Modal';
import { getTodoDetailApi, updateTodoApi } from '@/lib/apiClient';
import type { Todo } from '@/types/todo';

export default function TodoDetailPage() {
  const pageStyle: CSSProperties = {
    minHeight: '100svh',
    display: 'grid',
    placeItems: 'center',
    padding: 24,
  };
  const cardStyle: CSSProperties = { width: '100%', maxWidth: 520, display: 'grid', gap: 16 };
  const labelStyle: CSSProperties = { display: 'grid', gap: 6, textAlign: 'left' };
  const inputStyle: CSSProperties = {
    padding: '10px 12px',
    border: '1px solid #ccc',
    borderRadius: 8,
    fontSize: 16,
  };
  const selectStyle: CSSProperties = inputStyle;
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

  const params = useParams<{ id: string }>();
  const id = params?.id as string;
  const [todo, setTodo] = useState<Todo | null>(null);
  const [loading, setLoading] = useState(true);
  const [errorOpen, setErrorOpen] = useState(false);
  const [errorMsg, setErrorMsg] = useState('');

  useEffect(() => {
    let mounted = true;
    (async () => {
      try {
        const data = await getTodoDetailApi(id);
        if (mounted) setTodo(data);
      } catch (err) {
        const eobj = err as Error & { kind?: string };
        setErrorMsg(eobj.message || '取得に失敗しました');
        setErrorOpen(true);
      } finally {
        if (mounted) setLoading(false);
      }
    })();
    return () => {
      mounted = false;
    };
  }, [id]);

  if (loading)
    return (
      <main style={pageStyle}>
        <p>読み込み中...</p>
      </main>
    );
  if (!todo)
    return (
      <main style={pageStyle}>
        <p>データが見つかりません。</p>
      </main>
    );

  return (
    <main style={pageStyle}>
      <section style={cardStyle}>
        <h1 style={{ textAlign: 'center' }}>TODO 詳細</h1>
        <form
          onSubmit={async (e) => {
            e.preventDefault();
            const fd = new FormData(e.currentTarget as HTMLFormElement);
            const title = String(fd.get('title') || '');
            const description = String(fd.get('description') || '');
            const status = String(fd.get('status') || '');
            const due = String(fd.get('due_date') || '');
            try {
              const payload: any = { title };
              if (description) payload.description = description;
              if (status) payload.status = status;
              if (due) payload.due_date = due;
              await updateTodoApi(id, payload);
              alert('更新しました');
            } catch (err) {
              const eobj = err as Error & { kind?: string };
              setErrorMsg(eobj.message || '更新に失敗しました');
              setErrorOpen(true);
            }
          }}
        >
          <label style={labelStyle}>
            <span>タイトル</span>
            <input
              name="title"
              type="text"
              defaultValue={todo.title}
              required
              maxLength={140}
              style={inputStyle}
            />
          </label>

          <label style={labelStyle}>
            <span>説明</span>
            <textarea
              name="description"
              rows={4}
              defaultValue={todo.description ?? ''}
              style={{ ...inputStyle, resize: 'vertical' }}
            />
          </label>

          <label style={labelStyle}>
            <span>ステータス</span>
            <select name="status" defaultValue={todo.status} style={selectStyle}>
              <option value="todo">todo</option>
              <option value="doing">doing</option>
              <option value="done">done</option>
            </select>
          </label>

          <label style={labelStyle}>
            <span>期限</span>
            <input
              name="due_date"
              type="datetime-local"
              defaultValue={todo.due_date ?? ''}
              style={inputStyle}
            />
          </label>

          <button type="submit" style={submitStyle}>
            更新する
          </button>
        </form>

        <p style={helperStyle}>
          <Link href="/">← ホームへ戻る</Link>
        </p>
      </section>
      <Modal
        open={errorOpen}
        title="エラー"
        message={errorMsg}
        onClose={() => setErrorOpen(false)}
      />
    </main>
  );
}
