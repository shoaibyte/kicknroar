import { EmptyState } from '@/components/common/EmptyState';
import { CalendarX } from 'lucide-react';
import { Link } from 'react-router-dom';
import { Button } from '@/components/ui/button';

export const Matches: React.FC = () => {
  return (
    <div className="container py-10">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-3xl font-bold">Matches</h1>
        <Link to="/matches/create">
          <Button>Create Match</Button>
        </Link>
      </div>

      <EmptyState
        icon={CalendarX}
        title="No matches found"
        description="Try adjusting your filters or create a new match"
        action={{
          label: 'Create Match',
          onClick: () => (window.location.href = '/matches/create'),
        }}
      />
    </div>
  );
};

