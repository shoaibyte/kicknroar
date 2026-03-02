import { NavLink } from 'react-router-dom';
import { cn } from '@/lib/utils';

const navItems = [
  { path: '/', label: 'Home' },
  { path: '/matches', label: 'Matches' },
  { path: '/venues', label: 'Venues' },
  { path: '/dashboard', label: 'Dashboard' },
];

export const DesktopNav: React.FC = () => {
  return (
    <nav className="hidden md:flex items-center space-x-6">
      {navItems.map((item) => (
        <NavLink
          key={item.path}
          to={item.path}
          className={({ isActive }) =>
            cn(
              'text-sm font-medium transition-colors hover:text-primary',
              isActive ? 'text-primary' : 'text-muted-foreground'
            )
          }
        >
          {item.label}
        </NavLink>
      ))}
    </nav>
  );
};

