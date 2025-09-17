export type TodoStatus = 'todo' | 'doing' | 'done';

export type Todo = {
  id: string;
  user_id: string;
  title: string;
  description?: string | null;
  status: TodoStatus;
  due_date?: string | null;
};

export type TodoCreateRequest = {
  title: string;
  description?: string;
  status?: TodoStatus;
  due_date?: string;
};

export type TodoDetailResponse = Todo;
