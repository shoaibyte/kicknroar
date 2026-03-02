import { useState, useRef } from 'react';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useAuthStore } from '@/store/authStore';
import { uploadApi } from '@/api/endpoints/upload.api';
import { userApi } from '@/api/endpoints/user.api';
import { LoadingSpinner } from '@/components/common/LoadingSpinner';
import { extractErrorMessage } from '@/api/types';

const MAX_FILE_SIZE = 5 * 1024 * 1024; // 5MB
const ALLOWED_FILE_TYPES = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp'];

export const Profile: React.FC = () => {
  const { user, getCurrentUser } = useAuthStore();
  const [isUploading, setIsUploading] = useState(false);
  const [uploadError, setUploadError] = useState<string | null>(null);
  const [previewUrl, setPreviewUrl] = useState<string | null>(null);
  const [imageError, setImageError] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  // Helper function to get the full image URL
  const getImageUrl = (url: string | undefined | null): string | null => {
    if (!url) return null;
    // If URL is already absolute (starts with http:// or https://), return as is
    if (url.startsWith('http://') || url.startsWith('https://')) {
      return url;
    }
    // If URL is relative, construct full URL
    const baseURL = import.meta.env.DEV
      ? window.location.origin
      : (import.meta.env.VITE_API_URL || window.location.origin);
    // Remove /api/v1 if present in baseURL for image URLs
    const cleanBaseURL = baseURL.replace(/\/api\/v1$/, '');
    return url.startsWith('/') ? `${cleanBaseURL}${url}` : `${cleanBaseURL}/${url}`;
  };

  const validateFile = (file: File): string | null => {
    // Check file type
    if (!ALLOWED_FILE_TYPES.includes(file.type)) {
      return 'Invalid file type. Please upload jpg, jpeg, png, gif, or webp images only.';
    }

    // Check file size
    if (file.size > MAX_FILE_SIZE) {
      return 'File size exceeds 5MB limit. Please choose a smaller image.';
    }

    return null;
  };

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (!file) return;

    const validationError = validateFile(file);
    if (validationError) {
      setUploadError(validationError);
      setPreviewUrl(null);
      return;
    }

    // Create preview
    const reader = new FileReader();
    reader.onloadend = () => {
      setPreviewUrl(reader.result as string);
    };
    reader.readAsDataURL(file);

    setUploadError(null);
  };

  const handleUpload = async () => {
    const file = fileInputRef.current?.files?.[0];
    if (!file) {
      setUploadError('Please select a file to upload.');
      return;
    }

    const validationError = validateFile(file);
    if (validationError) {
      setUploadError(validationError);
      return;
    }

    setIsUploading(true);
    setUploadError(null);

    try {
      // Upload the file
      await uploadApi.uploadAvatar(file);
      
      // Fetch updated user profile data from /api/v1/users/me
      // The backend automatically updates profile_image_url after upload
      await getCurrentUser();

      // Reset file input and preview
      if (fileInputRef.current) {
        fileInputRef.current.value = '';
      }
      setPreviewUrl(null);
      setImageError(false);
    } catch (error: any) {
      setUploadError(extractErrorMessage(error));
    } finally {
      setIsUploading(false);
    }
  };

  const handleRemoveAvatar = async () => {
    setIsUploading(true);
    setUploadError(null);

    try {
      // Remove avatar by setting profile_image_url to empty
      await userApi.updateCurrentUser({
        profile_image_url: '',
      });

      // Fetch updated user profile data from /api/v1/users/me
      await getCurrentUser();

      if (fileInputRef.current) {
        fileInputRef.current.value = '';
      }
      setPreviewUrl(null);
      setImageError(false);
    } catch (error: any) {
      setUploadError(extractErrorMessage(error));
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <div className="container py-10">
      <Card>
        <CardHeader>
          <CardTitle>Profile</CardTitle>
          <CardDescription>Manage your account settings</CardDescription>
        </CardHeader>
        <CardContent>
          {user ? (
            <div className="space-y-6">
              {/* Avatar Section */}
              <div className="space-y-4">
                <Label>Profile Picture</Label>
                <div className="flex items-center gap-4">
                  <div className="relative">
                    {(previewUrl || (user.profile_image_url && !imageError)) ? (
                      <img
                        src={previewUrl || getImageUrl(user.profile_image_url) || ''}
                        alt="Profile"
                        className="h-24 w-24 rounded-full object-cover border-2 border-border"
                        onError={() => {
                          setImageError(true);
                        }}
                        onLoad={() => {
                          setImageError(false);
                        }}
                      />
                    ) : (
                      <div className="h-24 w-24 rounded-full bg-muted flex items-center justify-center border-2 border-border">
                        <span className="text-2xl font-semibold text-muted-foreground">
                          {user.full_name.charAt(0).toUpperCase()}
                        </span>
                      </div>
                    )}
                  </div>
                  <div className="flex flex-col gap-2">
                    <div className="flex gap-2">
                      <Input
                        ref={fileInputRef}
                        type="file"
                        accept="image/jpeg,image/jpg,image/png,image/gif,image/webp"
                        onChange={handleFileSelect}
                        className="hidden"
                        id="avatar-upload"
                        disabled={isUploading}
                      />
                      <Label htmlFor="avatar-upload" className="cursor-pointer">
                        <Button
                          type="button"
                          variant="outline"
                          size="sm"
                          disabled={isUploading}
                          onClick={() => fileInputRef.current?.click()}
                        >
                          Choose File
                        </Button>
                      </Label>
                      {(previewUrl || fileInputRef.current?.files?.[0]) && (
                        <Button
                          type="button"
                          size="sm"
                          onClick={handleUpload}
                          disabled={isUploading}
                        >
                          {isUploading ? (
                            <>
                              <LoadingSpinner size="sm" className="mr-2" />
                              Uploading...
                            </>
                          ) : (
                            'Upload'
                          )}
                        </Button>
                      )}
                      {user.profile_image_url && !previewUrl && (
                        <Button
                          type="button"
                          variant="outline"
                          size="sm"
                          onClick={handleRemoveAvatar}
                          disabled={isUploading}
                        >
                          Remove
                        </Button>
                      )}
                    </div>
                    <p className="text-xs text-muted-foreground">
                      Max 5MB. Supports: JPG, JPEG, PNG, GIF, WEBP
                    </p>
                    {uploadError && (
                      <p className="text-sm text-destructive">{uploadError}</p>
                    )}
                  </div>
                </div>
              </div>

              {/* User Info Section */}
              <div className="space-y-4 border-t pt-6">
                <div>
                  <Label className="text-sm font-medium">Full Name</Label>
                  <p className="text-muted-foreground mt-1">{user.full_name}</p>
                </div>
                <div>
                  <Label className="text-sm font-medium">Email</Label>
                  <p className="text-muted-foreground mt-1">{user.email}</p>
                </div>
                {user.phone && (
                  <div>
                    <Label className="text-sm font-medium">Phone</Label>
                    <p className="text-muted-foreground mt-1">{user.phone}</p>
                  </div>
                )}
                {user.skill_level && (
                  <div>
                    <Label className="text-sm font-medium">Skill Level</Label>
                    <p className="text-muted-foreground mt-1 capitalize">{user.skill_level}</p>
                  </div>
                )}
              </div>
            </div>
          ) : (
            <p className="text-muted-foreground">No user data available</p>
          )}
        </CardContent>
      </Card>
    </div>
  );
};
