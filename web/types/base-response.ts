export interface ApiError {
  code: number;
  message: string;
}

export interface BaseResponse<T> {
  data: T | null;
  error: ApiError | null;
  meta?: {
    limit: number,
    page: number,
    total: number,
    totalPages: number,
  };
}