export type LoginRequest = {
  email: string;
  password: string;
};

export type LoginResponse = {
  access_token: string;
};

export type SignupRequest = {
  email: string;
  password: string;
};
