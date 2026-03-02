import { useState, useEffect } from 'react';
import { cn } from '@/lib/utils';

interface CyclingTextProps {
  /**
   * Array of text options to cycle through
   */
  options: string[];
  /**
   * Duration in milliseconds for each text display
   * Default: 3000ms (3 seconds)
   */
  interval?: number;
  /**
   * Transition duration in milliseconds
   * Default: 500ms
   */
  transitionDuration?: number;
  /**
   * Additional CSS classes
   */
  className?: string;
}

/**
 * CyclingText Component
 * 
 * Displays text that cycles through an array of options with smooth transitions.
 * 
 * @example
 * <CyclingText 
 *   options={['Football', 'Futsal']} 
 *   interval={3000}
 * />
 */
export const CyclingText: React.FC<CyclingTextProps> = ({
  options,
  interval = 3000,
  transitionDuration = 500,
  className,
}) => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [isVisible, setIsVisible] = useState(true);

  useEffect(() => {
    if (options.length <= 1) return;

    const cycleInterval = setInterval(() => {
      // Fade out
      setIsVisible(false);
      
      // After fade out, change text and fade in
      setTimeout(() => {
        setCurrentIndex((prevIndex) => (prevIndex + 1) % options.length);
        setIsVisible(true);
      }, transitionDuration);
    }, interval);

    return () => clearInterval(cycleInterval);
  }, [options.length, interval, transitionDuration]);

  if (options.length === 0) return null;
  if (options.length === 1) {
    return <span className={className}>{options[0]}</span>;
  }

  // Find the longest option to set fixed width so adjacent text doesn't shift
  // Reduce width slightly to bring adjacent text closer
  const longestOption = options.reduce((a, b) => (a.length > b.length ? a : b));

  return (
    <span
      className={cn(
        'inline-block relative text-left',
        className
      )}
      style={{ 
        minWidth: `${longestOption.length - 0.5}ch`,
        marginRight: '0',
        paddingRight: '0'
      }}
    >
      <span
        className={cn(
          'inline-block transition-opacity duration-design',
          isVisible ? 'opacity-100' : 'opacity-0'
        )}
        style={{ transitionDuration: `${transitionDuration}ms` }}
      >
        {options[currentIndex]}
      </span>
    </span>
  );
};

