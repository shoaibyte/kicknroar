# Lottie Animations Guide

This directory is for storing Lottie animation JSON files to be used in the hero section.

## How to Add Animations

### Option 1: Download from LottieFiles (Recommended)

1. Visit [LottieFiles.com](https://lottiefiles.com)
2. Search for animations using keywords:
   - "football"
   - "soccer"
   - "soccer player"
   - "football match"
   - "sports"
   - "game"
3. Filter by:
   - **Free** animations
   - **JSON** format
4. Download the JSON file
5. Save it as `football.json` in this directory (`public/animations/`)

### Option 2: Use LottieFiles URL Directly

You can also use LottieFiles URLs directly in the component:

```tsx
<LottieBackground 
  src="https://lottie.host/embed/[animation-id].json" 
  opacity={0.4}
/>
```

### Recommended Animations

### Direct Links to LottieFiles (Free Animations)

Search on [LottieFiles.com](https://lottiefiles.com) for these keywords:

1. **Football/Soccer Animations:**
   - Search: "football", "soccer", "football player"
   - Look for: Player running, kicking, or celebrating animations
   - Recommended: Simple, looping animations work best for backgrounds

2. **Sports Background Animations:**
   - Search: "sports background", "game background"
   - Look for: Abstract sports-themed animations
   - Recommended: Subtle, non-distracting animations

3. **Ball Animations:**
   - Search: "soccer ball", "football ball"
   - Look for: Spinning or bouncing ball animations
   - Recommended: Simple looping animations

### Popular Free Animation Categories:
- "Football Match"
- "Soccer Player"
- "Sports Game"
- "Team Sports"
- "Football Field"

### Tips for Choosing:
- ✅ **Good:** Simple, looping animations
- ✅ **Good:** Animations with green/blue colors (matches your brand)
- ✅ **Good:** File size under 500KB
- ❌ **Avoid:** Complex animations with many elements
- ❌ **Avoid:** Animations that are too bright or distracting

## File Structure

```
public/
  animations/
    football.json  ← Place your animation file here
```

## Usage in Components

The `LottieBackground` component automatically loads animations from this directory:

```tsx
<LottieBackground 
  src="/animations/football.json" 
  opacity={0.4}
  loop={true}
  autoplay={true}
/>
```

## Tips

- **File Size:** Keep animations under 500KB for best performance
- **Complexity:** Simpler animations work better as backgrounds
- **Colors:** Animations with colors that complement your brand (green/blue) work best
- **Loop:** Most background animations should loop continuously
- **Opacity:** Use 0.3-0.5 opacity for subtle background effects

## Alternative Sources

- [Icons8 Lottie Animations](https://icons8.com/lottie-animations)
- [IconScout Lottie](https://iconscout.com/lotties)
- [Lordicon](https://lordicon.com)

