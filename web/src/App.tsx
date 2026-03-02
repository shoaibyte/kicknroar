import { Suspense } from 'react';
import { ErrorBoundary } from './components/common/ErrorBoundary';
import { AppRouter } from './router';
import { LoadingSpinner } from './components/common/LoadingSpinner';

function App() {
  return (
    <ErrorBoundary>
      <Suspense fallback={<LoadingSpinner />}>
        <AppRouter />
      </Suspense>
    </ErrorBoundary>
  );
}

export default App;

