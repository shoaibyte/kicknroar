import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';

export const CreateMatch: React.FC = () => {
  return (
    <div className="container py-10">
      <Card>
        <CardHeader>
          <CardTitle>Create Match</CardTitle>
          <CardDescription>Organize a new football match</CardDescription>
        </CardHeader>
        <CardContent>
          <p className="text-muted-foreground">Match creation form will be implemented here.</p>
        </CardContent>
      </Card>
    </div>
  );
};

