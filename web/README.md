# Kick&Roar Frontend

A Progressive Web Application (PWA) for discovering, creating, and joining football matches in Dhaka. Built with React 18, TypeScript, Vite, TailwindCSS, and ShadCN/UI.

## Tech Stack

- **React 18** - UI library
- **TypeScript** - Type safety
- **Vite** - Build tool and dev server
- **TailwindCSS** - Styling
- **ShadCN/UI** - Component library
- **Zustand** - State management
- **React Query** - Server state management
- **React Router** - Routing
- **Axios** - HTTP client
- **React Hook Form + Zod** - Form handling and validation

## Prerequisites

- Node.js 18+ and npm/yarn
- Git

## Development Setup

### 1. Clone Repository

```bash
git clone https://github.com/shoaibhassan/kicknroar-frontend.git
cd kicknroar-frontend
```

### 2. Install Dependencies

```bash
yarn install
# or
npm install
```

### 3. Setup Environment Variables

```bash
cp .env.example .env.local
```

Edit `.env.local` with your values:

```env
VITE_API_URL=http://localhost:8080/api/v1
VITE_GOOGLE_MAPS_API_KEY=your_google_maps_api_key_here
VITE_S3_BUCKET_URL=https://kicknroar-production.s3.ap-south-1.amazonaws.com
VITE_WS_URL=ws://localhost:8080
```

### 4. Start Development Server

```bash
yarn dev
# or
npm run dev
```

The app will be available at `http://localhost:3000`

## Available Scripts

- `yarn dev` - Start development server
- `yarn build` - Build for production
- `yarn build:staging` - Build for staging
- `yarn build:production` - Build for production
- `yarn preview` - Preview production build
- `yarn test` - Run tests
- `yarn test:ui` - Run tests with UI
- `yarn test:coverage` - Run tests with coverage
- `yarn lint` - Run ESLint
- `yarn format` - Format code with Prettier
- `yarn type-check` - Type check without emitting files

## Project Structure

```
src/
├── api/              # API client and endpoints
├── components/       # Reusable components
│   ├── ui/          # ShadCN UI components
│   ├── layout/      # Layout components
│   └── common/      # Common components
├── features/         # Feature modules
│   ├── auth/
│   ├── match/
│   ├── venue/
│   ├── user/
│   └── notification/
├── pages/           # Page components
├── hooks/           # Custom hooks
├── store/           # Zustand stores
├── lib/             # Utilities and constants
├── types/           # TypeScript types
└── styles/          # Global styles
```

## Features

- 🔐 User authentication (Login, Signup)
- ⚽ Match discovery and management
- 📍 Venue exploration with maps
- 👤 User profiles
- 🔔 Real-time notifications
- 📱 Mobile-first responsive design
- 🌐 PWA support (offline capability)

## Building for Production

```bash
yarn build
```

The production build will be in the `dist/` directory.

## Deployment

The project is configured for deployment on Render. See `render.yaml` for configuration details.

## Contributing

1. Create a feature branch: `git checkout -b feature/your-feature-name`
2. Make your changes
3. Commit: `git commit -m "feat: your feature description"`
4. Push: `git push origin feature/your-feature-name`
5. Create a Pull Request

## License

Copyright © 2025 Kick&Roar. All rights reserved.
