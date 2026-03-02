/**
 * Design System Tokens
 * Centralized design tokens following the Kick&Roar design system guidelines
 */

// Brand Colors
export const colors = {
  primary: '#2D7D32', // Forest Green - grass/field theme
  secondary: '#1565C0', // Sky Blue - trust/reliability
  accent: '#FF6F00', // Energy Orange - notifications/highlights
  success: '#4CAF50', // Bright Green
  warning: '#FF9800', // Amber
  error: '#F44336', // Red
} as const;

// Spacing Scale (8px grid system)
export const spacing = {
  0: '0',
  1: '0.125rem', // 2px
  2: '0.25rem', // 4px
  3: '0.375rem', // 6px
  4: '0.5rem', // 8px (base grid unit)
  5: '0.625rem', // 10px
  6: '0.75rem', // 12px
  8: '1rem', // 16px
  10: '1.25rem', // 20px
  12: '1.5rem', // 24px
  16: '2rem', // 32px
  20: '2.5rem', // 40px
  24: '3rem', // 48px
  32: '4rem', // 64px
  40: '5rem', // 80px
  48: '6rem', // 96px
  64: '8rem', // 128px
} as const;

// Border Radius
export const borderRadius = {
  button: '0.5rem', // 8px
  card: '0.75rem', // 12px
  container: '1rem', // 16px
} as const;

// Typography Scale
export const typography = {
  fontFamily: {
    primary: "'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', sans-serif",
  },
  fontWeight: {
    light: 300,
    regular: 400,
    medium: 500,
    semibold: 600,
    bold: 700,
    extrabold: 800,
  },
  fontSize: {
    xs: '0.75rem', // 12px
    sm: '0.875rem', // 14px
    base: '1rem', // 16px
    lg: '1.125rem', // 18px
    xl: '1.25rem', // 20px
    '2xl': '1.5rem', // 24px
    '3xl': '1.875rem', // 30px
    '4xl': '2.25rem', // 36px
    '5xl': '3rem', // 48px
    '6xl': '3.75rem', // 60px
  },
  lineHeight: {
    tight: 1.25,
    snug: 1.375,
    normal: 1.5,
    relaxed: 1.625,
    loose: 2,
  },
} as const;

// Shadows (Elevation)
export const shadows = {
  elevation1: '0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)',
  elevation2: '0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)',
  elevation3: '0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)',
} as const;

// Transitions
export const transitions = {
  duration: '300ms',
  timingFunction: 'ease',
  default: '300ms ease',
} as const;

// Touch Targets (Accessibility)
export const touchTargets = {
  minimum: {
    ios: '44px',
    android: '48px',
    default: '48px', // Use Android standard as default
  },
} as const;

// Component-specific tokens
export const components = {
  button: {
    minHeight: touchTargets.minimum.default,
    borderRadius: borderRadius.button,
    transition: transitions.default,
  },
  card: {
    borderRadius: borderRadius.card,
    shadow: shadows.elevation1,
    transition: transitions.default,
  },
  container: {
    borderRadius: borderRadius.container,
  },
} as const;

