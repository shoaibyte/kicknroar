import { Link } from 'react-router-dom';
import { useAuthStore } from '@/store/authStore';
import { Button } from '../ui/button';
import { User, LogOut } from 'lucide-react';

export const Header: React.FC = () => {
  const { isAuthenticated, logout } = useAuthStore();

  return (
    <header className="sticky top-0 z-40 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-16 items-center justify-between">
        <Link 
          to="/" 
          className="flex items-center space-x-2 transition-colors hover:opacity-80 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 rounded-button"
          aria-label="Kick&Roar Home"
        >
          <div className="flex items-center justify-center w-8 h-8 rounded-full bg-primary text-primary-foreground">
            <span className="text-sm font-bold">f</span>
          </div>
          <span className="text-xl font-bold text-primary">Kick&Roar</span>
        </Link>

        <nav className="hidden md:flex items-center space-x-6" aria-label="Main navigation">
          <Link 
            to="/matches" 
            className="text-sm font-medium text-foreground hover:text-primary transition-colors duration-design focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 rounded-button px-2 py-1"
          >
            Find Matches
          </Link>
          <Link 
            to="/venues" 
            className="text-sm font-medium text-foreground hover:text-primary transition-colors duration-design focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 rounded-button px-2 py-1"
          >
            Fields
          </Link>
        </nav>

        <div className="flex items-center space-x-4">
          {isAuthenticated ? (
            <>
              <Link to="/profile" aria-label="User profile">
                <Button 
                  variant="ghost" 
                  size="icon"
                  className="focus-visible:ring-2 focus-visible:ring-ring"
                >
                  <User className="h-5 w-5" aria-hidden="true" />
                </Button>
              </Link>
              <Button 
                variant="ghost" 
                size="icon" 
                onClick={logout}
                aria-label="Sign out"
                className="focus-visible:ring-2 focus-visible:ring-ring"
              >
                <LogOut className="h-5 w-5" aria-hidden="true" />
              </Button>
            </>
          ) : (
            <>
              <Link to="/login">
                <Button 
                  variant="ghost"
                  className="focus-visible:ring-2 focus-visible:ring-ring"
                >
                  Login
                </Button>
              </Link>
              <Link to="/signup">
                <Button 
                  className="bg-primary hover:bg-primary/90 focus-visible:ring-2 focus-visible:ring-ring"
                >
                  Sign In
                </Button>
              </Link>
            </>
          )}
        </div>
      </div>
    </header>
  );
};

