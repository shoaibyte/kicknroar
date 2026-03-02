import { useParams } from 'react-router-dom';
import { LoadingSpinner } from '@/components/common/LoadingSpinner';

export const MatchDetails: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  return (
    <div className="container py-10">
      <h1 className="text-3xl font-bold mb-6">Match Details</h1>
      <div className="flex items-center justify-center min-h-[400px]">
        <LoadingSpinner />
        <p className="ml-4 text-muted-foreground">Loading match details for ID: {id}</p>
      </div>
    </div>
  );
};

