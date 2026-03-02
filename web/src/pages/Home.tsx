import { Link } from 'react-router-dom';
import { Button } from '@/components/ui/button';
import { Card, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { LottieBackground } from '@/components/common/LottieBackground';
import { CyclingText } from '@/components/common/CyclingText';
import { VenueMap } from '@/components/common/VenueMap';

export const Home: React.FC = () => {
  return (
    <div className="flex flex-col min-h-screen">
      {/* Hero Section with Lottie Background */}
      <section className="relative bg-gradient-to-br from-[#1b5e20] via-primary to-[#4caf50] text-white overflow-hidden">
        {/* Pattern overlay */}
        <div 
          className="absolute inset-0 opacity-30 z-[1]"
          style={{
            backgroundImage: `
              repeating-linear-gradient(90deg, transparent 0px, transparent 48px, rgba(255,255,255,0.03) 50px, rgba(255,255,255,0.03) 52px),
              repeating-linear-gradient(0deg, transparent 0px, transparent 48px, rgba(255,255,255,0.03) 50px, rgba(255,255,255,0.03) 52px)
            `
          }}
        />
        
        {/* Lottie Animation - Responsive positioning */}
        <div className="absolute top-0 bottom-0 w-full md:w-2/3 lg:w-1/2 left-1/2 md:left-[32%] lg:left-[32%] -translate-x-1/2 z-[1] overflow-hidden">
          <LottieBackground 
            src="/animations/Soccer_player_kicking_ball.json" 
            opacity={0.3}
            loop={true}
            autoplay={true}
            speed={1}
          />
        </div>
        
        <div className="container relative z-10 py-20 md:py-32">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 lg:gap-16 items-center">
            {/* Hero Text - Left Column */}
            <div className="text-left relative z-10">
              <h1 className="text-4xl sm:text-5xl md:text-6xl lg:text-7xl font-extrabold tracking-tight leading-[1.1] mb-4 md:mb-6">
                Find Your Perfect{' '}
                <span style={{ color: '#FFEB3B' }}>
                  <CyclingText 
                    options={['Football', 'Futsal']} 
                    interval={3000}
                    transitionDuration={500}
                    className="inline-block"
                  /><span className="ml-1">Match</span>
                </span>
              </h1>
              <p className="text-lg md:text-xl lg:text-2xl text-white/90 mb-8 md:mb-10 leading-relaxed font-light">
                Connect with players in Dhaka. Share costs, find fields, and never play alone again.
              </p>
              <div className="flex flex-col sm:flex-row gap-4">
                <Link to="/matches" className="w-full sm:w-auto">
                  <Button 
                    size="lg"
                    className="w-full sm:w-auto bg-white text-primary hover:bg-white/90 font-bold text-base md:text-lg px-8 md:px-10 py-6 md:py-7 shadow-elevation-3 hover:shadow-elevation-2 hover:-translate-y-1 transition-all duration-design"
                  >
                    🗺️ Explore Matches
                  </Button>
                </Link>
                <Link to="/matches/create" className="w-full sm:w-auto">
                  <Button 
                    size="lg" 
                    variant="outline"
                    className="w-full sm:w-auto border-2 border-white/80 bg-white/10 backdrop-blur-sm text-white hover:bg-white hover:text-primary hover:border-white font-bold text-base md:text-lg px-8 md:px-10 py-6 md:py-7 hover:-translate-y-1 transition-all duration-design"
                  >
                    ➕ Create Event
                  </Button>
                </Link>
              </div>
            </div>

            {/* Live Matches Card - Right Column */}
            <div className="hidden lg:block">
              <Card className="bg-white/15 backdrop-blur-md border-white/20 rounded-container p-6 animate-float">
                <div className="flex justify-between items-center mb-4">
                  <h3 className="text-xl font-semibold text-white">Live Matches Near You</h3>
                  <span className="bg-error text-error-foreground px-3 py-1 rounded-full text-xs font-semibold animate-pulse">
                    LIVE
                  </span>
                </div>
                <div className="space-y-3">
                  <div className="bg-white/20 rounded-card p-4 hover:bg-white/30 transition-all duration-design hover:scale-[1.02]">
                    <div className="flex justify-between items-start mb-2">
                      <div>
                        <div className="font-semibold text-white mb-1">Morning Scrimmage</div>
                        <div className="text-sm text-white/80">Greenwood Park • 9:00 AM</div>
                      </div>
                    </div>
                    <div className="flex justify-between items-center text-sm">
                      <span className="text-white/90">7/10 joined</span>
                      <span className="text-accent font-semibold">150 BDT</span>
                    </div>
                  </div>
                  <div className="bg-white/20 rounded-card p-4 hover:bg-white/30 transition-all duration-design hover:scale-[1.02]">
                    <div className="flex justify-between items-start mb-2">
                      <div>
                        <div className="font-semibold text-white mb-1">Evening Match</div>
                        <div className="text-sm text-white/80">Central Sports • 6:00 PM</div>
                      </div>
                    </div>
                    <div className="flex justify-between items-center text-sm">
                      <span className="text-white/90">3/10 joined</span>
                      <span className="text-accent font-semibold">200 BDT</span>
                    </div>
                  </div>
                </div>
              </Card>
            </div>
          </div>
        </div>
      </section>

      {/* Map Section */}
      <section className="bg-white py-16 md:py-24">
        <div className="container">
          <div className="text-center mb-8 md:mb-12">
            <h2 className="text-3xl md:text-4xl font-bold mb-4 text-foreground">
              Discover Fields Across Dhaka
            </h2>
            <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
              Interactive map showing available futsal courts and ongoing matches
            </p>
          </div>
          <VenueMap />
        </div>
      </section>

      {/* Features Section */}
      <section className="bg-muted/30 py-16 md:py-24">
        <div className="container">
          <div className="text-center mb-12 md:mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4 text-foreground">
              Why Choose Kick&Roar?
            </h2>
            <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
              Everything you need to organize and join football matches
            </p>
          </div>

          <div className="grid gap-8 md:grid-cols-3 max-w-6xl mx-auto">
            <Card className="border-2 hover:border-primary/50 transition-all duration-design hover:-translate-y-2 hover:shadow-elevation-3">
              <CardHeader>
                <div className="w-12 h-12 rounded-xl bg-primary/10 flex items-center justify-center text-2xl mb-4">
                  🗺️
                </div>
                <CardTitle className="text-xl">Find Nearby Fields</CardTitle>
                <CardDescription>Discover futsal courts and football fields across Dhaka with real-time availability</CardDescription>
              </CardHeader>
            </Card>

            <Card className="border-2 hover:border-secondary/50 transition-all duration-design hover:-translate-y-2 hover:shadow-elevation-3">
              <CardHeader>
                <div className="w-12 h-12 rounded-xl bg-secondary/10 flex items-center justify-center text-2xl mb-4">
                  👥
                </div>
                <CardTitle className="text-xl">Connect with Players</CardTitle>
                <CardDescription>Join existing matches or create your own. Share costs and build your football community</CardDescription>
              </CardHeader>
            </Card>

            <Card className="border-2 hover:border-accent/50 transition-all duration-design hover:-translate-y-2 hover:shadow-elevation-3">
              <CardHeader>
                <div className="w-12 h-12 rounded-xl bg-accent/10 flex items-center justify-center text-2xl mb-4">
                  💰
                </div>
                <CardTitle className="text-xl">Split Costs Easily</CardTitle>
                <CardDescription>Affordable football for everyone. Share field rental costs among all players</CardDescription>
              </CardHeader>
            </Card>
          </div>
        </div>
      </section>
    </div>
  );
};

