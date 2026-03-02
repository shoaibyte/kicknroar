import { EmptyState } from '@/components/common/EmptyState';
import { MapPin } from 'lucide-react';

export const Venues: React.FC = () => {
  return (
    <div className="container py-10">
      <h1 className="text-3xl font-bold mb-6">Venues</h1>
      <EmptyState
        icon={MapPin}
        title="No venues found"
        description="Venue map and listing will be displayed here"
      />
    </div>
  );
};

