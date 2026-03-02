import { lazy } from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { Layout } from './components/layout/Layout';
import { ProtectedRoute } from './components/common/ProtectedRoute';

// Lazy load pages for code splitting
const Home = lazy(() => import('./pages/Home').then((m) => ({ default: m.Home })));
const Login = lazy(() => import('./pages/Login').then((m) => ({ default: m.Login })));
const Signup = lazy(() => import('./pages/Signup').then((m) => ({ default: m.Signup })));
const Matches = lazy(() => import('./pages/Matches').then((m) => ({ default: m.Matches })));
const MatchDetails = lazy(() =>
  import('./pages/MatchDetails').then((m) => ({ default: m.MatchDetails }))
);
const CreateMatch = lazy(() =>
  import('./pages/CreateMatch').then((m) => ({ default: m.CreateMatch }))
);
const Venues = lazy(() => import('./pages/Venues').then((m) => ({ default: m.Venues })));
const Profile = lazy(() => import('./pages/Profile').then((m) => ({ default: m.Profile })));
const Dashboard = lazy(() => import('./pages/Dashboard').then((m) => ({ default: m.Dashboard })));
const NotFound = lazy(() => import('./pages/NotFound').then((m) => ({ default: m.NotFound })));

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path: 'login',
        element: <Login />,
      },
      {
        path: 'signup',
        element: <Signup />,
      },
      {
        path: 'matches',
        element: (
          <ProtectedRoute>
            <Matches />
          </ProtectedRoute>
        ),
        children: [
          {
            path: ':id',
            element: <MatchDetails />,
          },
        ],
      },
      {
        path: 'matches/create',
        element: (
          <ProtectedRoute>
            <CreateMatch />
          </ProtectedRoute>
        ),
      },
      {
        path: 'venues',
        element: <Venues />,
      },
      {
        path: 'dashboard',
        element: (
          <ProtectedRoute>
            <Dashboard />
          </ProtectedRoute>
        ),
      },
      {
        path: 'profile',
        element: (
          <ProtectedRoute>
            <Profile />
          </ProtectedRoute>
        ),
      },
      {
        path: '*',
        element: <NotFound />,
      },
    ],
  },
]);

export const AppRouter = () => {
  return <RouterProvider router={router} />;
};

