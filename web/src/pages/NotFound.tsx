import { EmptyState } from '@/components/common/EmptyState';
import { AlertCircle } from 'lucide-react';

export const NotFound: React.FC = () => {
  return (
    <div className="container py-10">
      <EmptyState
        icon={AlertCircle}
        title="404 - Page Not Found"
        description="The page you're looking for doesn't exist or has been moved."
        action={{
          label: 'Go Home',
          onClick: () => (window.location.href = '/'),
        }}
      />
    </div>
  );
};

