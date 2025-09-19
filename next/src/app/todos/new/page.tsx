'use client';

import type { CSSProperties } from 'react';
import { useState } from 'react';
import Link from 'next/link';
import { Modal } from '@/components/ui/Modal';
import { createTodoApi } from '@/lib/apiClient';
import type { TodoCreateRequest, TodoStatus } from '@/types/todo';

export default function TodoNewPage() {
  const pageStyle: CSSProperties = {
    minHeight: '100svh',
    display: 'grid',
    placeItems: 'center',
    padding: 24,
  };
  const cardStyle: CSSProperties = { width: '100%', maxWidth: 480, display: 'grid', gap: 16 };
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
    backgroundColor: '#0ea5e9',
    color: '#fff',
    fontWeight: 600,
    border: 'none',
    cursor: 'pointer',
    display: 'block',
    margin: '16px auto 0',
  };
  const helperStyle: CSSProperties = { textAlign: 'center' };

  const [errorOpen, setErrorOpen] = useState(false);
  const [errorTitle, setErrorTitle] = useState('作成に失敗しました');
  const [errorMsg, setErrorMsg] = useState('');

  return (
    <main style={pageStyle}>
      <section style={cardStyle}>
        <h1 style={{ textAlign: 'center' }}>TODO 作成</h1>
        <p style={{ textAlign: 'center' }}>タイトルは必須。その他は任意です。</p>

        <form
          onSubmit={async (e) => {
            e.preventDefault();
            const form = e.currentTarget as HTMLFormElement;
            const fd = new FormData(form);
            const title = String(fd.get('title') || '');
            const description = String(fd.get('description') || '');
            const status = String(fd.get('status') || '');
            const due = String(fd.get('due_date') || '');
            try {
              const payload: Partial<TodoCreateRequest> & { title: string } = { title };
              if (description) payload.description = description;
              const isTodoStatus = (v: string): v is TodoStatus =>
                v === 'todo' || v === 'doing' || v === 'done';
              if (isTodoStatus(status)) payload.status = status;
              if (due) payload.due_date = due; // yyyy-mm-dd を送信
              const created = await createTodoApi(payload);
              // 作成後は一覧へ遷移
              window.location.href = '/todos';
            } catch (err) {
              const eobj = err as Error & { status?: number; kind?: string; field?: string };
              const kind = eobj.kind ?? 'error';
              if (kind === 'invalid') setErrorMsg(eobj.message || '入力内容を確認してください。');
              else if (kind === 'unauthorized') setErrorMsg('ログインが必要です。');
              else setErrorMsg(eobj.message || 'サーバーでエラーが発生しました。');
              setErrorOpen(true);
            }
          }}
        >
          <label style={labelStyle}>
            <span>タイトル（必須）</span>
            <input name="title" type="text" required maxLength={140} style={inputStyle} />
          </label>

          <label style={labelStyle}>
            <span>説明（任意）</span>
            <textarea name="description" rows={4} style={{ ...inputStyle, resize: 'vertical' }} />
          </label>

          <label style={labelStyle}>
            <span>ステータス（任意）</span>
            <select name="status" defaultValue="" style={selectStyle}>
              <option value="">未指定（todo）</option>
              <option value="todo">todo</option>
              <option value="doing">doing</option>
              <option value="done">done</option>
            </select>
          </label>

          <label style={labelStyle}>
            <span>期限（任意）</span>
            <input name="due_date" type="date" style={inputStyle} />
          </label>

          <button type="submit" style={submitStyle}>
            作成する
          </button>
        </form>

        <p style={helperStyle}>
          <Link href="/todos">← 一覧へ戻る</Link>
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
