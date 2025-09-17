import { API_BASE_URL } from '@/config/env';
import type { ApiErrorPayload } from '@/types/api';
import type { LoginRequest, LoginResponse, SignupRequest } from '@/types/auth';
import type { Todo, TodoCreateRequest } from '@/types/todo';

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${API_BASE_URL}${path}`, {
    ...init,
    headers: { 'Content-Type': 'application/json', ...(init?.headers ?? {}) },
  });
  if (!res.ok) {
    // 可能ならバックエンドの標準エラー形式を解釈する
    let parsed: ApiErrorPayload | undefined;
    try {
      parsed = (await res.json()) as ApiErrorPayload;
    } catch (_) {
      // JSONでない場合は素のテキスト
    }
    const message = parsed?.msg || res.statusText;
    const kind = parsed?.error || 'unknown';
    const field = parsed?.field;
    const error = new Error(message) as Error & {
      status?: number;
      kind?: string;
      field?: string;
    };
    error.status = res.status;
    error.kind = kind;
    if (field) error.field = field;
    throw error;
  }
  // 空ボディ/非JSONの成功レスポンス（201, 204など）を許容
  if (res.status === 204) return undefined as unknown as T;
  const contentType = res.headers.get('content-type')?.toLowerCase() ?? '';
  const isJson = contentType.includes('application/json');
  if (!isJson) return undefined as unknown as T;
  const text = await res.text();
  if (!text) return undefined as unknown as T;
  return JSON.parse(text) as T;
}

export async function loginApi(body: LoginRequest): Promise<LoginResponse> {
  return request<LoginResponse>('/auth/login', {
    method: 'POST',
    body: JSON.stringify(body),
  });
}

export async function signupApi(body: SignupRequest): Promise<void> {
  await request<void>('/auth/signup', {
    method: 'POST',
    body: JSON.stringify(body),
  });
}

function getCookie(name: string): string | undefined {
  if (typeof document === 'undefined') return undefined;
  const m = document.cookie.match(
    new RegExp('(?:^|; )' + name.replace(/([.$?*|{}()\[\]\\\/\+^])/g, '\\$1') + '=([^;]*)'),
  );
  return m ? decodeURIComponent(m[1]) : undefined;
}

function authHeaders(): HeadersInit {
  const token = getCookie('access_token');
  return token ? { Authorization: `Bearer ${token}` } : {};
}

export async function getTodosApi(): Promise<Todo[]> {
  return request<Todo[]>('/todos', {
    method: 'GET',
    headers: { ...authHeaders() },
  });
}

export async function getTodoDetailApi(id: string): Promise<Todo> {
  return request<Todo>(`/todos/${encodeURIComponent(id)}`, {
    method: 'GET',
    headers: { ...authHeaders() },
  });
}

export async function createTodoApi(body: TodoCreateRequest): Promise<Todo> {
  return request<Todo>('/todos', {
    method: 'POST',
    headers: { ...authHeaders() },
    body: JSON.stringify(body),
  });
}

export async function updateTodoApi(id: string, body: TodoCreateRequest): Promise<void> {
  await request<void>(`/todos/${encodeURIComponent(id)}`, {
    method: 'PUT',
    headers: { ...authHeaders() },
    body: JSON.stringify(body),
  });
}

export async function deleteTodosApi(ids: string[]): Promise<void> {
  await request<void>('/todos', {
    method: 'DELETE',
    headers: { ...authHeaders() },
    body: JSON.stringify({ ids }),
  });
}
