import { NavLink } from 'react-router-dom';
import { Home, Calendar, MapPin, User, Plus } from 'lucide-react';
import { cn } from '@/lib/utils';
import { useState } from 'react';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '../ui/dialog';

const navItems = [
  { path: '/', icon: Home, label: 'Home' },
  { path: '/matches', icon: Calendar, label: 'Matches' },
  { path: '/venues', icon: MapPin, label: 'Venues' },
  { path: '/profile', icon: User, label: 'Profile' },
];

export const MobileNav: React.FC = () => {
  const [isCreateOpen, setIsCreateOpen] = useState(false);
  return (
    <>
      <nav className="fixed bottom-0 left-0 right-0 z-50 bg-background border-t md:hidden">
        <div className="grid grid-cols-5 h-16">
          {navItems.slice(0, 2).map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              className={({ isActive }) =>
                cn(
                  'flex flex-col items-center justify-center space-y-1',
                  'text-muted-foreground hover:text-primary',
                  isActive && 'text-primary'
                )
              }
            >
              <item.icon className="h-5 w-5" />
              <span className="text-xs">{item.label}</span>
            </NavLink>
          ))}

          {/* Center Create Button */}
          <button
            onClick={() => setIsCreateOpen(true)}
            className="flex flex-col items-center justify-center"
          >
            <div className="bg-primary text-primary-foreground rounded-full p-3">
              <Plus className="h-6 w-6" />
            </div>
          </button>

          {navItems.slice(2).map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              className={({ isActive }) =>
                cn(
                  'flex flex-col items-center justify-center space-y-1',
                  'text-muted-foreground hover:text-primary',
                  isActive && 'text-primary'
                )
              }
            >
              <item.icon className="h-5 w-5" />
              <span className="text-xs">{item.label}</span>
            </NavLink>
          ))}
        </div>
      </nav>

      <Dialog open={isCreateOpen} onOpenChange={setIsCreateOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Create Match</DialogTitle>
          </DialogHeader>
          <p className="text-sm text-muted-foreground">
            Match creation form will be implemented here.
          </p>
        </DialogContent>
      </Dialog>
    </>
  );
};

