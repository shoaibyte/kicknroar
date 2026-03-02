import apiClient from '../client';

// Swagger response format: object with string properties
export interface UploadAvatarResponse {
  [key: string]: string;
}

export const uploadApi = {
  async uploadAvatar(file: File): Promise<UploadAvatarResponse> {
    const formData = new FormData();
    formData.append('file', file);

    const response = await apiClient.post<UploadAvatarResponse>(
      '/upload/avatar',
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      }
    );
    return response.data;
  },
};

