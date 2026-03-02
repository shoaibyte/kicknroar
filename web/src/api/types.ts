// API Response Types
// Swagger responses may return data directly or wrapped
// This type handles both cases
export type ApiResponse<T> = T | {
  data: T;
  message?: string;
  success?: boolean;
};

// Paginated Response - Swagger uses offset/limit pagination
export interface PaginatedResponse<T> {
  data: T[];
  pagination?: {
    offset: number;
    limit: number;
    total: number;
    total_pages?: number;
  };
}

// Swagger Error Response Format
export interface ErrorDetail {
  code: string;
  message: string;
  details?: string[];
}

export interface ErrorResponse {
  error: ErrorDetail;
}

export interface ApiError {
  code: string;
  message: string;
  details?: string[];
}

// Request Configuration
export interface RequestConfig {
  params?: Record<string, any>;
  headers?: Record<string, string>;
}

// Helper function to extract error message from Swagger error format
export function extractErrorMessage(error: any): string {
  if (error?.response?.data?.error) {
    const errorData = error.response.data.error;
    return errorData.message || errorData.code || 'An error occurred';
  }
  if (error?.response?.data?.message) {
    return error.response.data.message;
  }
  if (error?.message) {
    return error.message;
  }
  return 'An unexpected error occurred';
}

// Helper function to extract error details from Swagger error format
export function extractErrorDetails(error: any): string[] {
  if (error?.response?.data?.error?.details) {
    return error.response.data.error.details;
  }
  return [];
}

