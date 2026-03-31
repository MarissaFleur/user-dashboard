// types.ts
export type AuthState = {
  isAuthenticated: boolean;
  token: string | null;
  user: User | null;
};

export type User = {
  id: string;
  name: string;
  email: string;
  role: string;
};

export type DashboardState = {
  user: User;
  isLoading: boolean;
  error: string | null;
};

export type DashboardAction = 
  | { type: 'SET_LOADING' }
  | { type: 'SET_ERROR', error: string }
  | { type: 'SET_USER', user: User }
  | { type: 'SET_AUTH_STATE', state: AuthState };

export type AuthAction = 
  | { type: 'LOGIN', user: User }
  | { type: 'LOGOUT' }
  | { type: 'TOKEN_VALIDATED', token: string }
  | { type: 'TOKEN_INVALIDATED' };