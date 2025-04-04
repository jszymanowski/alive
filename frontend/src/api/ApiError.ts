export class ApiError extends Error {
  constructor(
    message: string,
    public readonly statusCode?: number,
    public readonly type?: string
  ) {
    super(message);
    this.name = 'ApiError';
  }
}
