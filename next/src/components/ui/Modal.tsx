'use client';

import type { CSSProperties, ReactNode } from 'react';

type ModalProps = {
  open: boolean;
  title?: string;
  message?: string;
  onClose: () => void;
  children?: ReactNode;
};

export function Modal({ open, title, message, onClose, children }: ModalProps) {
  if (!open) return null;

  const overlayStyle: CSSProperties = {
    position: 'fixed',
    inset: 0,
    backgroundColor: 'rgba(0,0,0,0.5)',
    display: 'grid',
    placeItems: 'center',
    zIndex: 1000,
  };
  const dialogStyle: CSSProperties = {
    width: 'min(92vw, 420px)',
    background: '#fff',
    borderRadius: 12,
    padding: 20,
    display: 'grid',
    gap: 12,
  };
  const titleStyle: CSSProperties = { fontSize: 18, fontWeight: 700 };
  const closeStyle: CSSProperties = {
    marginTop: 8,
    display: 'inline-block',
    alignSelf: 'end',
    padding: '8px 12px',
    borderRadius: 8,
    backgroundColor: '#e5e7eb',
    border: 'none',
    cursor: 'pointer',
  };

  return (
    <div role="dialog" aria-modal="true" style={overlayStyle} onClick={onClose}>
      <section style={dialogStyle} onClick={(e) => e.stopPropagation()}>
        {title ? <h2 style={titleStyle}>{title}</h2> : null}
        {message ? <p>{message}</p> : null}
        {children}
        <button type="button" onClick={onClose} style={closeStyle} aria-label="閉じる">
          閉じる
        </button>
      </section>
    </div>
  );
}
