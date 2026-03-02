import { lazy, Suspense, useState, useEffect } from 'react';
import { cn } from '@/lib/utils';

// Lazy load Lottie component for performance
const Lottie = lazy(() => import('lottie-react'));

interface LottieBackgroundProps {
  /**
   * Path to the .json animation file or URL
   * Can be placed in public/animations/ directory or use LottieFiles URL
   */
  src?: string;
  /**
   * Opacity of the animation (0-1)
   * Default: 0.4 for subtle background effect
   */
  opacity?: number;
  /**
   * Additional CSS classes
   */
  className?: string;
  /**
   * Whether the animation should loop
   * Default: true
   */
  loop?: boolean;
  /**
   * Whether the animation should autoplay
   * Default: true
   */
  autoplay?: boolean;
  /**
   * Animation speed (1 = normal, 2 = double speed, etc.)
   * Default: 1
   */
  speed?: number;
}

/**
 * LottieBackground Component
 * 
 * A reusable component for displaying Lottie animations as backgrounds.
 * Supports lazy loading and configurable opacity for subtle effects.
 * 
 * @example
 * <LottieBackground 
 *   src="/animations/football.json" 
 *   opacity={0.4}
 * />
 * 
 * Or use LottieFiles URL:
 * <LottieBackground 
 *   src="https://lottie.host/embed/..." 
 *   opacity={0.4}
 * />
 */
export const LottieBackground: React.FC<LottieBackgroundProps> = ({
  src,
  opacity = 0.4,
  className,
  loop = true,
  autoplay = true,
  speed: _speed = 1, // accepted for API; lottie-react type may not expose speed
}) => {
  const [animationData, setAnimationData] = useState<any>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [hasError, setHasError] = useState(false);

  useEffect(() => {
    if (!src) {
      setIsLoading(false);
      return;
    }

    // Check if it's a URL or local path
    const isUrl = src.startsWith('http://') || src.startsWith('https://');
    
    if (isUrl) {
      // Fetch from URL
      fetch(src)
        .then((response) => {
          if (!response.ok) throw new Error('Failed to fetch animation');
          return response.json();
        })
        .then((data) => {
          setAnimationData(data);
          setIsLoading(false);
        })
        .catch(() => {
          setHasError(true);
          setIsLoading(false);
        });
    } else {
      // Fetch local file
      fetch(src)
        .then((response) => {
          if (!response.ok) throw new Error('Failed to fetch animation');
          return response.json();
        })
        .then((data) => {
          setAnimationData(data);
          setIsLoading(false);
        })
        .catch(() => {
          setHasError(true);
          setIsLoading(false);
        });
    }
  }, [src]);

  return (
    <div
      className={cn(
        'absolute inset-0 overflow-hidden pointer-events-none z-0',
        className
      )}
      aria-hidden="true"
    >
      {/* Always show gradient as base layer */}
      <div 
        className="absolute inset-0 bg-gradient-to-br from-primary/15 via-primary/8 to-secondary/15"
        style={{ opacity }}
      />
      
      {/* Overlay Lottie animation if loaded successfully */}
      {!isLoading && !hasError && animationData && (
        <div 
          className="absolute inset-0 flex items-center justify-center"
          style={{ opacity }}
        >
          <Suspense fallback={null}>
            <Lottie
              animationData={animationData}
              loop={loop}
              autoplay={autoplay}
              style={{
                width: '100%',
                height: '100%',
                maxWidth: '100%',
                maxHeight: '100%',
              }}
            />
          </Suspense>
        </div>
      )}
    </div>
  );
};

